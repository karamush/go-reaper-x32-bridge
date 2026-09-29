// Command x32listen dumps every OSC datagram received on a UDP port, splitting
// bundles and decoding each message. It is the REAPER side companion of the
// emulator's x32probe:
//
//	x32listen -listen 0.0.0.0:9000                 # print what REAPER sends
//	x32listen -listen 0.0.0.0:9000 -send 127.0.0.1:8000 -msg "/action/41743"
//	x32listen -listen 0.0.0.0:9100 -send 127.0.0.1:10023 -meters 1
//
// The -msg flag is repeated for every probe message; its argument can be given
// after a space ("/track/1/volume 0.5", floats and integers are detected).
// The -meters flag subscribes to a meter group of a console and prints the
// non-zero values of every frame, which is how the meter feed of the emulator is
// checked (see docs/X32_REAPER_BRIDGE.md).
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"x32emu/osc"
)

// version, commit and date are injected by the linker in release builds (see
// .goreleaser.yml); a plain "go build" leaves them at these defaults.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func versionString() string {
	return fmt.Sprintf("%s (commit %s, built %s)", version, commit, date)
}

type repeated []string

func (r *repeated) String() string { return strings.Join(*r, ", ") }
func (r *repeated) Set(v string) error {
	*r = append(*r, v)
	return nil
}

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	listen := flag.String("listen", "0.0.0.0:9000", "local address to listen on")
	send := flag.String("send", "", "optional host:port to send the probe messages to")
	dur := flag.Duration("t", 5*time.Second, "how long to listen")
	raw := flag.Bool("raw", false, "also print the hexadecimal dump")
	var probes repeated
	flag.Var(&probes, "msg", "probe message, e.g. \"/track/1/volume 0.5\" (repeatable)")
	metersGroup := flag.Int("meters", -1, "subscribe to /meters/<n> and decode the frames (-1 = off)")
	flag.Parse()
	if *showVersion {
		fmt.Printf("x32listen %s\n", versionString())
		return
	}

	conn, err := net.ListenUDP("udp4", mustAddr(*listen))
	if err != nil {
		fmt.Fprintln(os.Stderr, "listen:", err)
		os.Exit(1)
	}
	defer conn.Close()

	if *send != "" && (*metersGroup >= 0 || len(probes) > 0) {
		dst, err := net.ResolveUDPAddr("udp4", *send)
		if err != nil {
			fmt.Fprintln(os.Stderr, "send:", err)
			os.Exit(1)
		}
		if *metersGroup >= 0 {
			sub := osc.Start("/meters", "siii")
			sub = osc.AppendString(sub, "/meters/"+strconv.Itoa(*metersGroup))
			sub = osc.AppendInt32(sub, 0)
			sub = osc.AppendInt32(sub, 0)
			sub = osc.AppendInt32(sub, 0)
			fmt.Printf("-> %s\n", osc.Dump(">", sub, false))
			if _, err := conn.WriteToUDP(sub, dst); err != nil {
				fmt.Fprintln(os.Stderr, "send:", err)
			}
		}
		for _, p := range probes {
			msg, err := buildProbe(p)
			if err != nil {
				fmt.Fprintln(os.Stderr, "msg:", err)
				continue
			}
			fmt.Printf("-> %s\n", osc.Dump(">", msg, false))
			if _, err := conn.WriteToUDP(msg, dst); err != nil {
				fmt.Fprintln(os.Stderr, "send:", err)
			}
			time.Sleep(20 * time.Millisecond)
		}
	}

	fmt.Printf("listening on %s for %s\n", conn.LocalAddr(), *dur)
	buf := make([]byte, 8192)
	end := time.Now().Add(*dur)
	seen := map[string]int{}
	total := 0
	for time.Now().Before(end) {
		_ = conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
		n, from, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}
		total++
		_ = from
		for _, sub := range splitBundle(append([]byte(nil), buf[:n]...)) {
			m := osc.Decode(sub)
			if m.Addr == "" {
				continue
			}
			seen[m.Addr]++
			if strings.HasPrefix(m.Addr, "/meters/") && len(m.Args) == 1 && m.Args[0].Tag == 'b' {
				printMeters(m.Addr, m.Args[0].B)
				if !*raw {
					continue
				}
			}
			fmt.Printf("<- %s\n", osc.Dump("<", sub, *raw))
		}
	}
	fmt.Printf("\n%d datagrams, %d distinct addresses\n", total, len(seen))
	for addr, n := range seen {
		fmt.Printf("  %-34s x%d\n", addr, n)
	}
}

// printMeters decodes one meter frame. The blob starts with the number of values
// (little-endian, the count a console sends) followed by the values themselves:
// 32 bit little-endian floats for the level groups, 16 bit samples for the RTA
// (/meters/15) and the gain reduction group (/meters/16). See docs/
// X32_OSC_PROTOCOL.md §7.6.
func printMeters(addr string, b []byte) {
	if len(b) < 4 {
		fmt.Printf("  %s: short frame (%d bytes)\n", addr, len(b))
		return
	}
	n := int(binary.LittleEndian.Uint32(b[:4]))
	data := b[4:]
	var parts []string
	if strings.HasSuffix(addr, "/15") || strings.HasSuffix(addr, "/16") {
		for i := 0; i+1 < len(data); i += 2 {
			if v := float32(int16(binary.LittleEndian.Uint16(data[i:]))) / 256; v != 0 {
				parts = append(parts, fmt.Sprintf("%d=%07.2f", i/2, v))
			}
		}
	} else {
		for i := 0; i < n && 4*i+4 <= len(data); i++ {
			if v := math.Float32frombits(binary.LittleEndian.Uint32(data[4*i:])); v != 0 {
				parts = append(parts, fmt.Sprintf("%d=%.4f", i, v))
			}
		}
	}
	fmt.Printf("  %s: %d values, %d non-zero: %s\n", addr, n, len(parts), strings.Join(parts, " "))
}

// splitBundle splits an OSC bundle into its messages (same helper as the bridge).
func splitBundle(b []byte) [][]byte {
	if len(b) < 16 || string(b[:7]) != "#bundle" || b[7] != 0 {
		return [][]byte{b}
	}
	out := [][]byte{}
	p := 16
	for p+4 <= len(b) {
		n := int(binary.BigEndian.Uint32(b[p:]))
		p += 4
		if n <= 0 || p+n > len(b) {
			break
		}
		out = append(out, b[p:p+n])
		p += n
	}
	return out
}

// buildProbe turns "\"/track/1/volume 0.5\"" into an OSC datagram.
func buildProbe(s string) ([]byte, error) {
	s = strings.TrimSpace(strings.Trim(s, "\""))
	if s == "" {
		return nil, fmt.Errorf("empty message")
	}
	addr := s
	var arg string
	if i := strings.IndexByte(s, ' '); i > 0 {
		addr, arg = s[:i], strings.TrimSpace(s[i+1:])
	}
	if arg == "" {
		return osc.Start(addr, ""), nil
	}
	if f, err := strconv.ParseFloat(arg, 32); err == nil {
		if strings.ContainsAny(arg, ".eE") {
			return osc.Build(addr, "f", osc.Arg{Tag: 'f', F: float32(f)}), nil
		}
		return osc.Build(addr, "i", osc.Arg{Tag: 'i', I: int32(f)}), nil
	}
	return osc.Build(addr, "s", osc.Arg{Tag: 's', S: arg}), nil
}

func mustAddr(s string) *net.UDPAddr {
	a, err := net.ResolveUDPAddr("udp4", s)
	if err != nil {
		fmt.Fprintln(os.Stderr, "address:", err)
		os.Exit(1)
	}
	return a
}
