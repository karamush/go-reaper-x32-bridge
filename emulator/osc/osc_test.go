package osc

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"
)

func TestPaddingAndHeader(t *testing.T) {
	// "/info" + NUL + 3 padding, then ",s" + NUL + 1 padding
	b := Start("/info", "s")
	if !bytes.Equal(b, []byte{'/', 'i', 'n', 'f', 'o', 0, 0, 0, ',', 's', 0, 0}) {
		t.Fatalf("unexpected header: %v", b)
	}
	// addresses are padded so that the tag block starts on a 4 byte boundary
	b = Start("/ch/01/mix/fader", "f")
	if len(b) != 24 {
		t.Fatalf("header length = %d, want 24", len(b))
	}
	if b[20] != ',' || b[21] != 'f' {
		t.Fatalf("tag block at wrong offset: %v", b)
	}
}

func TestDecodeOldNotation(t *testing.T) {
	m := Decode(Start("/ch/01/mix/fader", ""))
	if m.Addr != "/ch/01/mix/fader" {
		t.Fatalf("addr = %q", m.Addr)
	}
	if m.Tags != "" || len(m.Args) != 0 {
		t.Fatalf("old notation should have no tags, got %q %v", m.Tags, m.Args)
	}
}

func TestDecodeArguments(t *testing.T) {
	b := Start("/ch/01/mix/fader", "fs")
	b = AppendFloat32(b, 0.5)
	b = AppendString(b, "Test")
	m := Decode(b)
	if m.Addr != "/ch/01/mix/fader" || m.Tags != "fs" || len(m.Args) != 2 {
		t.Fatalf("decoded %+v", m)
	}
	if m.Args[0].F != 0.5 || m.Args[1].S != "Test" {
		t.Fatalf("args = %+v", m.Args)
	}
	// old style reply of a float GET
	b = Start("/ch/01/mix/fader", "f")
	b = AppendFloat32(b, 1.0)
	if len(b) != 28 { // 20 bytes header + 4 + 4 (see the X32_Command trace)
		t.Fatalf("float reply length = %d, want 28", len(b))
	}
	if got := math.Float32frombits(binary.BigEndian.Uint32(b[24:])); got != 1.0 {
		t.Fatalf("float = %v", got)
	}
}

func TestMeterFrameLayout(t *testing.T) {
	// what prepMeter() builds: header + BE blob size + LE count + zero data
	const l = 96
	head := Start("/meters/1", "b")
	frame := make([]byte, len(head)+4+4+4*l)
	copy(frame, head)
	binary.BigEndian.PutUint32(frame[len(head):], uint32((l+1)*4))
	binary.LittleEndian.PutUint32(frame[len(head)+4:], uint32(l))
	if len(frame) != 408 {
		t.Fatalf("frame length = %d, want 408", len(frame))
	}
	m := Decode(frame)
	if m.Addr != "/meters/1" || m.Tags != "b" || len(m.Args) != 1 {
		t.Fatalf("decoded %+v", m)
	}
	if len(m.Args[0].B) != (l+1)*4 {
		t.Fatalf("blob size = %d, want %d", len(m.Args[0].B), (l+1)*4)
	}
	if got := binary.LittleEndian.Uint32(m.Args[0].B); got != l {
		t.Fatalf("value count = %d, want %d", got, l)
	}
}

func TestDump(t *testing.T) {
	b := Start("/ch/01/mix/fader", "f")
	b = AppendFloat32(b, 1.0)
	got := Dump("X->", b, false)
	want := "X->,   28 B: /ch/01/mix/fader~~~~,f~~[1.0000]"
	if got != want {
		t.Fatalf("dump = %q, want %q", got, want)
	}
}
