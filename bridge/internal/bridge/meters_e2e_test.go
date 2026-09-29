package bridge

import (
	"encoding/binary"
	"math"
	"net"
	"testing"
	"time"

	"x32emu/emu"
	"x32emu/osc"
)

// End to end: the levels REAPER reports for a track arrive in the /meters/1 frame
// of the emulator, on the value index of the console channel the track is shown
// on. That is what X32-Edit and Mixing Station display, and the reason the
// all-in-one x32reaper command builds both sides in one process.
func TestMetersEndToEnd(t *testing.T) {
	console, err := emu.New(emu.Config{
		BindIP:   "127.0.0.1",
		Port:     "0",
		Name:     "e2e",
		Resource: t.TempDir() + "\\.X32res.rc",
	})
	if err != nil {
		t.Fatal(err)
	}
	go func() { _ = console.Run() }()
	defer func() {
		console.Stop()
		_ = console.Close()
	}()

	cfg := DefaultConfig()
	cfg.Verbose = false
	cfg.X32Addr = console.Address()
	cfg.Listen = "127.0.0.1:0"
	b, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	b.log = func(string, ...any) {}
	b.SetMeters(MeterSink{
		Track: func(section string, number int, v float32) {
			if section == "ch" {
				console.SetMeterLevel(emu.MeterCh, number-1, v)
			}
		},
	})
	go func() { _ = b.Run() }()
	defer b.Stop()

	// where the fake REAPER sends its feedback
	var reaperAddr *net.UDPAddr
	for i := 0; i < 200 && reaperAddr == nil; i++ {
		time.Sleep(10 * time.Millisecond)
		if b.reapConn != nil {
			la := b.reapConn.LocalAddr().(*net.UDPAddr)
			reaperAddr = &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: la.Port}
		}
	}
	if reaperAddr == nil {
		t.Fatal("the bridge did not start")
	}
	reaper, err := net.DialUDP("udp4", nil, reaperAddr)
	if err != nil {
		t.Fatal(err)
	}
	defer reaper.Close()

	// the console client subscribes to the channel meters
	client, err := net.Dial("udp4", console.Address())
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	req := osc.Start("/meters", "siii")
	req = osc.AppendString(req, "/meters/1")
	req = osc.AppendInt32(req, 0)
	req = osc.AppendInt32(req, 0)
	req = osc.AppendInt32(req, 0)
	if _, err := client.Write(req); err != nil {
		t.Fatal(err)
	}

	const want = 0.875
	if _, err := reaper.Write(osc.Build("/track/5/vu", "f", osc.Arg{Tag: 'f', F: want})); err != nil {
		t.Fatal(err)
	}

	deadline := time.Now().Add(3 * time.Second)
	buf := make([]byte, 1024)
	for time.Now().Before(deadline) {
		_ = client.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		n, err := client.Read(buf)
		if err != nil {
			continue
		}
		frame := buf[:n]
		m := osc.Decode(frame)
		if m.Addr != "/meters/1" || len(m.Args) != 1 {
			t.Fatalf("unexpected frame %+v", m)
		}
		// the first frames are sent before REAPER's level reaches the bridge
		if got := math.Float32frombits(binary.LittleEndian.Uint32(frame[24+4*4:])); got != want {
			continue
		}
		// every other channel is silent
		for ch := 0; ch < 32; ch++ {
			if ch == 4 {
				continue
			}
			if v := math.Float32frombits(binary.LittleEndian.Uint32(frame[24+4*ch:])); v != 0 {
				t.Fatalf("channel %d level = %v, want 0", ch+1, v)
			}
		}
		return
	}
	t.Fatal("no meter frame carrying the REAPER level")
}
