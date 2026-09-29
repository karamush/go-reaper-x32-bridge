package x32

import (
	"encoding/binary"
	"math"
	"testing"

	"x32emu/osc"
)

// subscribeMeters sends the subscription a console client uses
// ("/meters ,siii \"/meters/N\" 0 0 0", see X32Automix.c/X32TapW.c) and returns
// the frame the emulator builds for it.
func subscribeMeters(t *testing.T, st *State, sends *[]sent, group int) []byte {
	t.Helper()
	before := len(*sends)
	req := osc.Start("/meters", "siii")
	req = osc.AppendString(req, "/meters/"+itoa(group))
	req = osc.AppendInt32(req, 0)
	req = osc.AppendInt32(req, 0)
	req = osc.AppendInt32(req, 0)
	st.HandlePacket(req, udpAddr(t, "192.168.0.20:1000"))
	if len(*sends) != before {
		t.Fatalf("the meter request itself must not be answered")
	}
	st.PumpMeters()
	if len(*sends) != before+1 {
		t.Fatalf("expected one meter frame, got %d", len(*sends)-before)
	}
	return (*sends)[before].b
}

// value reads the 32 bit little-endian float at the given value index of a frame.
func meterValue(t *testing.T, frame []byte, index int) float32 {
	t.Helper()
	off := 24 + 4*index
	if off+4 > len(frame) {
		t.Fatalf("value %d is out of the frame (%d bytes)", index, len(frame))
	}
	return math.Float32frombits(binary.LittleEndian.Uint32(frame[off:]))
}

// The layout table must match the value count of every group (Xprepmeter in
// X32.c): a group is either fully described or left untouched (zeros).
func TestMeterLayoutSizes(t *testing.T) {
	for i, size := range meterSizes {
		n := len(meterLayout[i])
		if n != 0 && n != size {
			t.Fatalf("group %d: layout has %d values, frame carries %d", i, n, size)
		}
	}
	// the group the reference client reads (X32Automix.c)
	if n := len(meterLayout[1]); n != 96 {
		t.Fatalf("group 1 has %d values, want 96", n)
	}
	for ch := 0; ch < 32; ch++ {
		sl := meterLayout[1][ch]
		if sl.kind != MeterCh || sl.idx != ch {
			t.Fatalf("group 1 slot %d = %+v, want channel %d", ch, sl, ch)
		}
	}
	// the RTA and the gain reduction groups carry 16 bit samples, not levels
	if len(meterLayout[15]) != 0 || len(meterLayout[16]) != 0 || len(meterLayout[4]) != 0 {
		t.Fatalf("groups 4, 15 and 16 must stay at zero")
	}
}

// Without an outside source the frames stay zero: this is the reference
// behaviour of X32.c (the emulator has no audio) and what the differential test
// against X32-emulator.exe compares.
func TestMetersAreZeroWithoutSource(t *testing.T) {
	st, sends := newTestState(t)
	frame := subscribeMeters(t, st, sends, 1)
	for _, b := range frame[24:] {
		if b != 0 {
			t.Fatalf("frame carries data without a source: % x", frame[24:])
		}
	}
}

// Levels fed through SetMeterLevel() must show up at the offsets clients read:
// /meters/1 carries the 32 channel levels first (X32Automix.c reads the level of
// channel ch at 24 + 4*ch), the other groups follow the layout table.
func TestMeterLevelsInFrames(t *testing.T) {
	st, sends := newTestState(t)
	st.SetMeterLevel(MeterCh, 0, 0.5)
	st.SetMeterLevel(MeterCh, 31, 0.25)
	st.SetMeterLevel(MeterBus, 0, 0.75)
	st.SetMeterLevel(MeterMain, 0, 0.125)

	frame := subscribeMeters(t, st, sends, 1)
	if got := meterValue(t, frame, 0); got != 0.5 {
		t.Fatalf("channel 1 level = %v, want 0.5", got)
	}
	if got := meterValue(t, frame, 31); got != 0.25 {
		t.Fatalf("channel 32 level = %v, want 0.25", got)
	}
	// the two blocks after the levels are gain reductions: no source, so zero
	if got := meterValue(t, frame, 32); got != 0 {
		t.Fatalf("gain reduction = %v, want 0", got)
	}
	if len(frame) != 24+4*meterSizes[1] {
		t.Fatalf("frame length = %d, want %d", len(frame), 24+4*meterSizes[1])
	}

	// group 7 is the 16 bus levels, one value each
	frame = subscribeMeters(t, st, sends, 7)
	if got := meterValue(t, frame, 0); got != 0.75 {
		t.Fatalf("bus 1 level = %v, want 0.75", got)
	}
	if got := meterValue(t, frame, 1); got != 0 {
		t.Fatalf("bus 2 level = %v, want 0", got)
	}

	// group 6 carries the main bus at the first three values
	frame = subscribeMeters(t, st, sends, 6)
	if got := meterValue(t, frame, 0); got != 0.125 {
		t.Fatalf("main L level = %v, want 0.125", got)
	}

	// group 9 and 10 are the channels again, one value each
	frame = subscribeMeters(t, st, sends, 9)
	if got := meterValue(t, frame, 31); got != 0.25 {
		t.Fatalf("group 9 channel 32 level = %v, want 0.25", got)
	}

	// group 14 lists every fader: 32 channels, then aux, fx, bus, matrix, main
	st.SetMeterLevel(MeterAuxin, 0, 0.375)
	frame = subscribeMeters(t, st, sends, 14)
	if got := meterValue(t, frame, 32); got != 0.375 {
		t.Fatalf("group 14 aux 1 level = %v, want 0.375", got)
	}
	if got := meterValue(t, frame, 32+8+8); got != 0.75 {
		t.Fatalf("group 14 bus 1 level = %v, want 0.75", got)
	}
	if got := meterValue(t, frame, 32+8+8+16+6); got != 0.125 {
		t.Fatalf("group 14 main L level = %v, want 0.125", got)
	}
	// the eight DCA slots have no meters and stay zero
	for i := 72; i < 80; i++ {
		if got := meterValue(t, frame, i); got != 0 {
			t.Fatalf("group 14 value %d = %v, want 0", i, got)
		}
	}
}

// Values are little-endian floats (Xdump.c) and out of range indices are simply
// ignored.
func TestMeterLevelsAreLittleEndian(t *testing.T) {
	st, sends := newTestState(t)
	st.SetMeterLevel(MeterCh, 1, 1)
	frame := subscribeMeters(t, st, sends, 1)
	off := 24 + 4
	if got := frame[off : off+4]; got[3] != 0x3f || got[2] != 0x80 || got[1] != 0 || got[0] != 0 {
		t.Fatalf("1.0 is not little-endian at offset %d: % x", off, got)
	}
	st.SetMeterLevel(MeterCh, 99, 1) // out of range: ignored
	st.SetMeterLevel(MeterCh, -1, 1) // out of range: ignored
	if got := st.MeterLevel(MeterCh, 99); got != 0 {
		t.Fatalf("out of range level = %v", got)
	}
	if got := st.MeterLevel(MeterCh, 1); got != 1 {
		t.Fatalf("level = %v, want 1", got)
	}
}
