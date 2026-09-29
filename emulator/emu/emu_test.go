package emu

import (
	"encoding/binary"
	"math"
	"net"
	"testing"
	"time"

	"x32emu/osc"
)

// newTestEmulator starts an emulator on a free port, with no resource file.
func newTestEmulator(t *testing.T) *Emulator {
	t.Helper()
	e, err := New(Config{
		BindIP:   "127.0.0.1",
		Port:     "0",
		Name:     "Test Desk",
		Resource: t.TempDir() + "\\.X32res.rc",
	})
	if err != nil {
		t.Fatal(err)
	}
	go func() { _ = e.Run() }()
	t.Cleanup(func() {
		e.Stop()
		_ = e.Close()
	})
	return e
}

// The public wrapper must answer like the emulator itself: /xinfo carries the
// name given at startup and the address it advertises.
func TestEmulatorXinfo(t *testing.T) {
	e := newTestEmulator(t)
	conn, err := net.Dial("udp4", e.Address())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	if _, err := conn.Write(osc.Start("/xinfo", "")); err != nil {
		t.Fatal(err)
	}
	_ = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	m := osc.Decode(buf[:n])
	if m.Addr != "/xinfo" || len(m.Args) != 4 {
		t.Fatalf("answer = %+v", m)
	}
	if m.Args[0].S != "127.0.0.1" || m.Args[1].S != "Test Desk" || m.Args[3].S != "4.06" {
		t.Fatalf("answer args = %+v", m.Args)
	}
	if e.Port() == "0" || e.Port() == "" {
		t.Fatalf("port = %q, want the port in use", e.Port())
	}
}

// A client that subscribes to /meters/1 must receive frames carrying the levels
// fed through SetMeterLevel(): that is how the bridge shows REAPER's volume
// meters on the desk (and in X32-Edit or Mixing Station).
func TestEmulatorMeterFeed(t *testing.T) {
	e := newTestEmulator(t)
	conn, err := net.Dial("udp4", e.Address())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	req := osc.Start("/meters", "siii")
	req = osc.AppendString(req, "/meters/1")
	req = osc.AppendInt32(req, 0)
	req = osc.AppendInt32(req, 0)
	req = osc.AppendInt32(req, 0)
	if _, err := conn.Write(req); err != nil {
		t.Fatal(err)
	}

	// feed channel 3 and wait for a frame that carries it
	const want = 0.625
	e.SetMeterLevel(MeterCh, 2, want)
	deadline := time.Now().Add(3 * time.Second)
	buf := make([]byte, 1024)
	for time.Now().Before(deadline) {
		_ = conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		n, err := conn.Read(buf)
		if err != nil {
			continue
		}
		frame := buf[:n]
		m := osc.Decode(frame)
		if m.Addr != "/meters/1" {
			t.Fatalf("unexpected frame %q", m.Addr)
		}
		if len(m.Args) != 1 || len(m.Args[0].B) != 4*(96+1) {
			t.Fatalf("blob length = %d", len(m.Args[0].B))
		}
		// the frame sent right at subscription still has the level of the moment
		// the feed started
		if got := math.Float32frombits(binary.LittleEndian.Uint32(frame[24+4*2:])); got != want {
			continue
		}
		if e.MeterLevel(MeterCh, 2) != want {
			t.Fatalf("level read back = %v", e.MeterLevel(MeterCh, 2))
		}
		// a channel nobody feeds stays at zero
		if v := math.Float32frombits(binary.LittleEndian.Uint32(frame[24:28])); v != 0 {
			t.Fatalf("channel 1 level = %v, want 0", v)
		}
		return
	}
	t.Fatal("no meter frame with the fed level")
}
