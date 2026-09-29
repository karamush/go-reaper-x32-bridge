package bridge

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"x32emu/osc"
)

// TestLiveReaper talks to a real REAPER. It is skipped unless X32BRIDGE_LIVE=1
// is set: it needs REAPER running with an OSC control surface listening on
// 127.0.0.1:8000 and sending to 127.0.0.1:9000 (REAPER: Preferences,
// Control/OSC/web; the defaults of the bridge).
//
// It never modifies the project: the only message it sends back to REAPER
// carries the value REAPER itself reported for that track.
//
//	$env:X32BRIDGE_LIVE=1; go test ./internal/bridge -run Live -v
func TestLiveReaper(t *testing.T) {
	if os.Getenv("X32BRIDGE_LIVE") != "1" {
		t.Skip("set X32BRIDGE_LIVE=1 to talk to a running REAPER")
	}
	console, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer console.Close()

	cfg := DefaultConfig()
	cfg.Verbose = false
	cfg.X32Addr = console.LocalAddr().String()
	cfg.Listen = "127.0.0.1:9000"
	cfg.Host, cfg.Port = "127.0.0.1", "8000"
	b, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	var mu sync.Mutex
	var lines []string
	b.log = func(format string, args ...any) {
		mu.Lock()
		lines = append(lines, fmt.Sprintf(format, args...))
		mu.Unlock()
	}
	b.verbose = true
	go func() { _ = b.Run() }()
	defer b.Stop()

	// the bridge creates its console socket in Run(): the desk messages of the
	// test are sent to that very port (the socket listens on the wildcard
	// address, which cannot be used as a destination)
	var bridgeAddr *net.UDPAddr
	for i := 0; i < 100 && bridgeAddr == nil; i++ {
		time.Sleep(20 * time.Millisecond)
		if b.x32Conn != nil {
			la := b.x32Conn.LocalAddr().(*net.UDPAddr)
			bridgeAddr = &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: la.Port}
		}
	}
	if bridgeAddr == nil {
		t.Fatal("the bridge did not start")
	}

	// collect what the bridge forwards to the desk for a while
	seen := map[string]*osc.Message{}
	count := map[string]int{}
	deadline := time.Now().Add(3 * time.Second)
	buf := make([]byte, 4096)
	for time.Now().Before(deadline) {
		_ = console.SetReadDeadline(time.Now().Add(50 * time.Millisecond))
		n, _, err := console.ReadFromUDP(buf)
		if err != nil {
			continue
		}
		m := osc.Decode(append([]byte(nil), buf[:n]...))
		if m.Addr != "" {
			seen[m.Addr] = m
			count[m.Addr]++
		}
	}
	classes := map[string]int{}
	for addr := range seen {
		switch {
		case strings.HasSuffix(addr, "/config/name"):
			classes["names"]++
		case strings.HasSuffix(addr, "/mix/fader"):
			classes["faders"]++
		case strings.HasSuffix(addr, "/mix/pan"):
			classes["pans"]++
		case strings.HasSuffix(addr, "/mix/on"):
			classes["mutes"]++
		case strings.HasSuffix(addr, "/level"):
			classes["sends"]++
		case strings.HasPrefix(addr, "/-stat/solosw"):
			classes["solo"]++
		case strings.HasPrefix(addr, "/-stat/userpar"), strings.HasPrefix(addr, "/config/userctrl"):
			classes["bank C setup"]++
		}
	}
	t.Logf("REAPER -> desk: %d distinct addresses", len(seen))
	for k, v := range classes {
		t.Logf("  %-14s %d", k, v)
	}
	for _, addr := range []string{"/ch/01/mix/fader", "/ch/01/config/name"} {
		if m := seen[addr]; m != nil {
			t.Logf("  %-20s = %s", addr, valueOf(m))
		}
	}
	if seen["/ch/01/mix/fader"] == nil {
		t.Fatalf("REAPER reported no track 1: is the OSC surface sending to %s?", cfg.Listen)
	}

	// desk -> REAPER: send back the value REAPER reported (no audible change) and
	// check that the bridge forwarded it
	v, ok := argNum(seen["/ch/01/mix/fader"])
	if !ok {
		t.Fatal("cannot read the volume REAPER reported")
	}
	mu.Lock()
	lines = nil
	mu.Unlock()
	if _, err := console.WriteToUDP(osc.Build("/ch/01/mix/fader", "f", osc.Arg{Tag: 'f', F: v}), bridgeAddr); err != nil {
		t.Fatal(err)
	}
	time.Sleep(300 * time.Millisecond)
	mu.Lock()
	defer mu.Unlock()
	found := false
	for _, l := range lines {
		if strings.Contains(l, "->R") && strings.Contains(l, "/track/1/volume") {
			found = true
			t.Logf("desk -> REAPER: %s", strings.TrimSpace(l))
		}
	}
	if !found {
		t.Fatalf("the desk message was not forwarded to REAPER (log: %v)", lines)
	}
}

// valueOf renders the first argument of a message for the log.
func valueOf(m *osc.Message) string {
	if len(m.Args) == 0 {
		return "(no argument)"
	}
	switch m.Args[0].Tag {
	case 'f':
		return fmt.Sprintf("%.4f", m.Args[0].F)
	case 'i':
		return fmt.Sprintf("%d", m.Args[0].I)
	case 's':
		return m.Args[0].S
	}
	return "?"
}
