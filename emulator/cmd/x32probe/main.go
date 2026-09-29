// Command x32probe runs a scripted OSC session against an X32 emulator and
// prints every datagram it receives in a stable, diffable form. It is used to
// compare the Go port with the reference C emulator (X32-emulator.exe of the
// upstream X32-Behringer project; tools/difftest.ps1 drives both):
//
//	go run ./cmd/x32probe -addr 127.0.0.1:10023 > c.txt
//	go run ./cmd/x32probe -addr 127.0.0.1:10023 > go.txt
//
// Socket "A" registers itself with /xremote, socket "B" is the plain requester,
// so notifications can be observed separately from the answers.
package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"net"
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

type step struct {
	desc string
	pkt  []byte
	from byte // 'A' = xremote client, 'B' = requester
	wait time.Duration
}

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	addrStr := flag.String("addr", "127.0.0.1:10023", "emulator address")
	waitMS := flag.Int("wait", 100, "reply collection window per step (ms)")
	userPar := flag.Bool("userpar", false, "append the /-stat/userpar scenario (a real console feature the C emulator does not have)")
	flag.Parse()
	if *showVersion {
		fmt.Printf("x32probe %s\n", versionString())
		return
	}

	raddr, err := net.ResolveUDPAddr("udp4", *addrStr)
	if err != nil {
		fmt.Println("bad address:", err)
		return
	}
	a, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		fmt.Println("socket A:", err)
		return
	}
	defer a.Close()
	b, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		fmt.Println("socket B:", err)
		return
	}
	defer b.Close()

	// the notifying client registers first, like X32-Edit does
	_, _ = a.WriteToUDP(osc.Start("/xremote", ""), raddr)
	time.Sleep(30 * time.Millisecond)

	steps := script()
	if *userPar {
		steps = append(steps, scriptUserPar()...)
	}
	for i, st := range steps {
		// Keep the /xremote subscription of A alive: a console drops a client that
		// has been silent for 11 seconds, and real clients (X32-Edit, Mixing
		// Station, X32ReaperW) therefore re-register every 9 seconds or so. Doing
		// it well inside that window keeps the whole session deterministic (a
		// silent client would stop receiving notifications mid-run, which used to
		// make the number of differing lines of the differential test fluctuate).
		if i > 0 && i%20 == 0 {
			_, _ = a.WriteToUDP(osc.Start("/xremote", ""), raddr)
		}
		sock, who := b, byte('B')
		if st.from == 'A' {
			sock, who = a, 'A'
		}
		_, _ = sock.WriteToUDP(st.pkt, raddr)
		w := st.wait
		if w == 0 {
			w = time.Duration(*waitMS) * time.Millisecond
		}
		fromA, fromB := collectBoth(a, b, w)
		fmt.Printf("=== step %02d %s\n", i+1, st.desc)
		fmt.Printf("  ->%c %s\n", who, render(st.pkt))
		for _, p := range fromB {
			fmt.Printf("  <-B %s\n", render(p))
		}
		for _, p := range fromA {
			fmt.Printf("  <-A %s\n", render(p))
		}
	}
}

// collectBoth reads both sockets with a shared deadline so that the answers on
// B and the notifications on A cannot be attributed to different steps.
func collectBoth(a, b *net.UDPConn, w time.Duration) (fromA, fromB [][]byte) {
	end := time.Now().Add(w)
	buf := make([]byte, 2048)
	push := func(dst *[][]byte, n int) {
		p := make([]byte, n)
		copy(p, buf[:n])
		if len(*dst) > 0 && hex.EncodeToString((*dst)[len(*dst)-1]) == hex.EncodeToString(p) {
			return // identical meter frame
		}
		*dst = append(*dst, p)
	}
	for time.Now().Before(end) {
		_ = a.SetReadDeadline(time.Now().Add(2 * time.Millisecond))
		if n, _, err := a.ReadFromUDP(buf); err == nil {
			push(&fromA, n)
		}
		_ = b.SetReadDeadline(time.Now().Add(2 * time.Millisecond))
		if n, _, err := b.ReadFromUDP(buf); err == nil {
			push(&fromB, n)
		}
	}
	return fromA, fromB
}

// collect reads datagrams until the deadline, dropping consecutive duplicates
// (meter frames repeat with identical content).
func collect(c *net.UDPConn, w time.Duration) [][]byte {
	var out [][]byte
	deadline := time.Now().Add(w)
	buf := make([]byte, 2048)
	for {
		_ = c.SetReadDeadline(deadline)
		n, _, err := c.ReadFromUDP(buf)
		if err != nil {
			return out
		}
		p := make([]byte, n)
		copy(p, buf[:n])
		if len(out) > 0 && hex.EncodeToString(out[len(out)-1]) == hex.EncodeToString(p) {
			continue // duplicate meter frame
		}
		out = append(out, p)
	}
}

func render(p []byte) string {
	return hex.EncodeToString(p) + " | " + osc.Dump("<", p, false)
}

// scriptUserPar covers "/-stat/userpar/NN/value", the user assign controls of a
// real console (the C reference emulator has no such address, so these steps are
// opt-in: they are only used when checking the Go port on its own).
func scriptUserPar() []step {
	s := func(desc, path, tags string, args ...osc.Arg) step {
		return step{desc: desc, pkt: osc.Build(path, tags, args...), from: 'B'}
	}
	old := func(desc, path string) step {
		return step{desc: desc, pkt: osc.AppendString(nil, path), from: 'B'}
	}
	sa := func(desc string) step {
		return step{desc: desc, pkt: osc.Start("/xremote", ""), from: 'A'}
	}
	// A is registered with /xremote (see main), so its answers are notifications.
	// The subscription expires after 11 s (like on a real console), so it is
	// refreshed here - exactly what a bridge does every 9 s.
	return []step{
		sa("renew the /xremote subscription of A (no answer)"),
		old("GET userpar button 17 (Bank C)", "/-stat/userpar/17/value"),
		old("GET userpar encoder 33 (Bank C, centered)", "/-stat/userpar/33/value"),
		s("SET userpar 18 = 127 (button down, X32ReaperW ignores it)", "/-stat/userpar/18/value", "i", osc.Arg{Tag: 'i', I: 127}),
		s("SET userpar 18 = 0 (button up, X32ReaperW acts on it)", "/-stat/userpar/18/value", "i", osc.Arg{Tag: 'i', I: 0}),
		s("SET userpar 18 = 0 again (silence)", "/-stat/userpar/18/value", "i", osc.Arg{Tag: 'i', I: 0}),
		s("SET userpar encoder 33 = 100", "/-stat/userpar/33/value", "i", osc.Arg{Tag: 'i', I: 100}),
		s("SET userpar 01 (not a Bank C control, silence)", "/-stat/userpar/01/value", "i", osc.Arg{Tag: 'i', I: 55}),
		s("SET userpar 40 (out of range, silence)", "/-stat/userpar/40/value", "i", osc.Arg{Tag: 'i', I: 55}),
	}
}

func script() []step {
	s := func(desc, path, tags string, args ...osc.Arg) step {
		return step{desc: desc, pkt: osc.Build(path, tags, args...), from: 'B'}
	}
	sa := func(desc, path, tags string, args ...osc.Arg) step {
		st := s(desc, path, tags, args...)
		st.from = 'A'
		return st
	}
	old := func(desc, path string) step {
		return step{desc: desc, pkt: osc.AppendString(nil, path), from: 'B'}
	}
	var out []step

	// ---- life cycle and discovery
	out = append(out,
		sa("xremote from A (no answer expected)", "/xremote", ""),
		s("info", "/info", ""),
		s("xinfo", "/xinfo", ""),
		s("status", "/status", ""),
		s("unknown command (silence)", "/foobar", "i", osc.Arg{Tag: 'i', I: 1}),
		s("renew (silence)", "/renew", ""),
		s("libs (silence)", "/-libs", ""),
	)

	// ---- GET of single parameters (old notation)
	for _, p := range []string{
		"/ch/01/mix/fader", "/ch/01/config/name", "/ch/01/config/icon", "/ch/01/config/color",
		"/ch/01/preamp/trim", "/ch/01/preamp/invert", "/ch/01/gate/thr", "/ch/01/dyn/ratio",
		"/ch/01/eq/1/f", "/ch/01/mix/09/level", "/ch/01/mix/09/type", "/ch/01/grp/mute",
		"/ch/01/automix/group", "/ch/32/mix/fader", "/bus/01/mix/fader", "/bus/16/mix/fader",
		"/auxin/01/mix/fader", "/auxin/08/preamp/trim", "/fxrtn/01/mix/fader",
		"/mtx/01/mix/fader", "/main/st/mix/fader", "/dca/1/fader", "/dca/8/config/name",
		"/fx/1/type", "/fx/1/par/01", "/fx/5/type", "/fx/5/par/03",
		"/headamp/000/gain", "/headamp/001/phantom", "/headamp/127/gain",
		"/outputs/main/01/src", "/outputs/main/01/delay/time",
		"/config/solo/level", "/config/mute/1", "/config/routing/IN/1-8",
		"/-prefs/name", "/-prefs/bright", "/-prefs/ip/addr/0", "/-stat/selidx", "/-stat/lock",
		"/-show/showfile/scene/001/name", "/-show/showfile/scene/001/hasdata",
		"/-libs/ch/001/name", "/-ha/00/index",
	} {
		out = append(out, old("GET "+p, p))
	}
	out = append(out, s("GET with empty tag string", "/ch/01/mix/fader", ""))

	// ---- SET (change detection and notification of the other client)
	out = append(out,
		s("SET fader (notification to A only)", "/ch/01/mix/fader", "f", osc.Arg{Tag: 'f', F: 0.5}),
		s("SET fader again (silence)", "/ch/01/mix/fader", "f", osc.Arg{Tag: 'f', F: 0.5}),
		s("SET name", "/ch/01/config/name", "s", osc.Arg{Tag: 's', S: "Guitar"}),
		s("SET multi parameter (config group)", "/ch/01/config", "siii",
			osc.Arg{Tag: 's', S: "Gtr"}, osc.Arg{Tag: 'i', I: 3}, osc.Arg{Tag: 'i', I: 4}, osc.Arg{Tag: 'i', I: 2}),
		s("SET enum", "/ch/01/preamp/invert", "i", osc.Arg{Tag: 'i', I: 1}),
		s("SET int", "/ch/01/grp/mute", "i", osc.Arg{Tag: 'i', I: 12}),
		s("SET bus level", "/bus/01/mix/fader", "f", osc.Arg{Tag: 'f', F: 0.75}),
		s("SET dca", "/dca/1/fader", "f", osc.Arg{Tag: 'f', F: 0.25}),
		s("SET fx type", "/fx/1/type", "i", osc.Arg{Tag: 'i', I: 10}),
		s("SET fx param", "/fx/1/par/05", "f", osc.Arg{Tag: 'f', F: 0.5}),
		s("SET scene name", "/-show/showfile/scene/001/name", "s", osc.Arg{Tag: 's', S: "Scene 1"}),
	)

	// ---- /node group queries
	for _, p := range []string{
		"ch/01/config", "ch/01/preamp", "ch/01/gate", "ch/01/gate/filter", "ch/01/dyn",
		"ch/01/eq/1", "ch/01/insert", "ch/01/delay", "ch/01/grp", "ch/01/mix/09",
		"ch/01/automix", "ch/32/config", "bus/01/config", "main/st/config", "main/st/mix/01",
		"dca/1", "fx/1", "fx/5", "fx/1/source", "headamp/001", "outputs/main/01",
		"auxin/01/preamp", "mtx/01/dyn", "config/solo", "config/mute", "config/routing/IN",
		"-prefs", "-prefs/rta", "-stat", "-stat/solosw", "-show/showfile/scene/001",
		"-libs/ch/001", "-urec", "-usb", "ch/01",
	} {
		out = append(out, s("/node "+p, "/node", "s", osc.Arg{Tag: 's', S: p}))
	}
	out = append(out,
		s("/node single -prefs/rta/visibility", "/node", "s", osc.Arg{Tag: 's', S: "-prefs/rta/visibility"}),
		s("/node single ch/01/mix/fader", "/node", "s", osc.Arg{Tag: 's', S: "ch/01/mix/fader"}),
	)

	// ---- "/" commands
	for _, p := range []string{
		"/ch/01/mix/fader 0.62",
		"/ch/01/config/name \"Lead\"",
		"/ch/01/preamp/trim 6",
		"/ch/01/eq/1/f 1000",
		"/ch/01/mix/09/level -10",
		"/ch/01/grp/mute %0110",
		"/ch/01/dyn/thr -20",
		"/bus/01/config/name \"Bass\"",
		"/dca/1/config/name \"Drums\"",
		"/main/st/config/name \"MAIN\"",
		"/headamp/001/gain 20",
		"/config/solo/level 3",
		"/-stat/selmode",
	} {
		out = append(out, s("/ "+p, "/", "s", osc.Arg{Tag: 's', S: p}))
	}

	// ---- library, scene and file commands
	out = append(out,
		s("copy libchan", "/copy", "siii", osc.Arg{Tag: 's', S: "libchan"}, osc.Arg{Tag: 'i', I: 1}, osc.Arg{Tag: 'i', I: 2}, osc.Arg{Tag: 'i', I: 3}),
		s("copy bad type", "/copy", "siii", osc.Arg{Tag: 's', S: "scene"}, osc.Arg{Tag: 'i', I: 1}, osc.Arg{Tag: 'i', I: 2}, osc.Arg{Tag: 'i', I: 3}),
		s("save scene", "/save", "siss", osc.Arg{Tag: 's', S: "scene"}, osc.Arg{Tag: 'i', I: 2}, osc.Arg{Tag: 's', S: "My Scene"}, osc.Arg{Tag: 's', S: "note"}),
		old("GET scene 002 name", "/-show/showfile/scene/002/name"),
		s("save snippet", "/save", "siss", osc.Arg{Tag: 's', S: "snippet"}, osc.Arg{Tag: 'i', I: 3}, osc.Arg{Tag: 's', S: "Snip"}, osc.Arg{Tag: 's', S: "n"}),
		s("save libchan", "/save", "siss", osc.Arg{Tag: 's', S: "libchan"}, osc.Arg{Tag: 'i', I: 1}, osc.Arg{Tag: 's', S: "x"}, osc.Arg{Tag: 's', S: "y"}),
		s("delete scene", "/delete", "si", osc.Arg{Tag: 's', S: "scene"}, osc.Arg{Tag: 'i', I: 2}),
		old("GET scene 002 name after delete", "/-show/showfile/scene/002/name"),
		s("add cue", "/add", "sis", osc.Arg{Tag: 's', S: "cue"}, osc.Arg{Tag: 'i', I: 0}, osc.Arg{Tag: 's', S: "x"}),
		s("load", "/load", "sis", osc.Arg{Tag: 's', S: "libchan"}, osc.Arg{Tag: 'i', I: 0}, osc.Arg{Tag: 's', S: "x"}),
		s("showdump", "/showdump", ""),
	)

	// ---- meters (frames arrive asynchronously)
	out = append(out,
		meterStep("meters subscription /meters/1 from A", "/meters/1", 'A'),
		meterStep("meters subscription /meters/15 from B", "/meters/15", 'B'),
	)

	// ---- lock (1 is not the shutdown value) and unsubscribe
	out = append(out,
		s("lock 1 (no shutdown)", "/-stat/lock", "i", osc.Arg{Tag: 'i', I: 1}),
		sa("unsubscribe from A", "/unsubscribe", ""),
		s("SET from B after A unsubscribed", "/ch/01/mix/fader", "f", osc.Arg{Tag: 'f', F: 0.4}),
		old("GET fader (final state)", "/ch/01/mix/fader"),
	)
	return out
}

// meterStep builds a "/meters ,siii <group> 0 0 0" request.
func meterStep(desc, group string, from byte) step {
	p := osc.Start("/meters", "siii")
	p = osc.AppendString(p, group)
	p = osc.AppendInt32(p, 0)
	p = osc.AppendInt32(p, 0)
	p = osc.AppendInt32(p, 0)
	return step{desc: desc, pkt: p, from: from, wait: 250 * time.Millisecond}
}
