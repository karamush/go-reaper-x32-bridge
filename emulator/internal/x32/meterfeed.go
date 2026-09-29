package x32

import (
	"encoding/binary"
	"math"
)

// ---------------------------------------------------------------- meter levels
//
// A real console fills its meter frames with the levels of its own signals. The
// emulator has no audio, so X32.c reports zeros and so does this port - unless an
// outside source feeds values through SetMeterLevel(). The REAPER bridge of this
// workspace does exactly that (see docs/X32_REAPER_BRIDGE.md), which is why the
// level table is protected by a mutex: the bridge runs in another goroutine than
// the server loop.

// MeterKind selects a family of console meters. MeterNone is the zero value:
// an unset slot of the layout table carries no meter at all and leaves the value
// at zero, which is what the C emulator always reports.
type MeterKind int

const (
	MeterNone  MeterKind = iota // no meter at all
	MeterCh                     // input channels 1..32
	MeterAuxin                  // aux inputs 1..8
	MeterFxrtn                  // effect returns 1..8
	MeterBus                    // mix buses 1..16
	MeterMtx                    // matrices 1..6
	MeterMain                   // main bus: 0 = L, 1 = R, 2 = mono
)

// MeterLevels holds the meters of one console, indexed per family. Values use
// the console meter scale: 0 is -oo, 1 is 0 dBFS and the scale reaches 8.0
// (+18 dB), exactly what a real console sends. The zero value means "no signal",
// which is what the C emulator reports all the time.
type MeterLevels struct {
	Ch   [32]float32
	Aux  [8]float32
	Fxr  [8]float32
	Bus  [16]float32
	Mtx  [6]float32
	Main [3]float32
}

// get returns one level, 0 for an unknown kind or index.
func (l *MeterLevels) get(k MeterKind, i int) float32 {
	switch k {
	case MeterCh:
		if i >= 0 && i < len(l.Ch) {
			return l.Ch[i]
		}
	case MeterAuxin:
		if i >= 0 && i < len(l.Aux) {
			return l.Aux[i]
		}
	case MeterFxrtn:
		if i >= 0 && i < len(l.Fxr) {
			return l.Fxr[i]
		}
	case MeterBus:
		if i >= 0 && i < len(l.Bus) {
			return l.Bus[i]
		}
	case MeterMtx:
		if i >= 0 && i < len(l.Mtx) {
			return l.Mtx[i]
		}
	case MeterMain:
		if i >= 0 && i < len(l.Main) {
			return l.Main[i]
		}
	}
	return 0
}

// set stores one level, ignoring unknown kinds and out of range indices.
func (l *MeterLevels) set(k MeterKind, i int, v float32) {
	switch k {
	case MeterCh:
		if i >= 0 && i < len(l.Ch) {
			l.Ch[i] = v
		}
	case MeterAuxin:
		if i >= 0 && i < len(l.Aux) {
			l.Aux[i] = v
		}
	case MeterFxrtn:
		if i >= 0 && i < len(l.Fxr) {
			l.Fxr[i] = v
		}
	case MeterBus:
		if i >= 0 && i < len(l.Bus) {
			l.Bus[i] = v
		}
	case MeterMtx:
		if i >= 0 && i < len(l.Mtx) {
			l.Mtx[i] = v
		}
	case MeterMain:
		if i >= 0 && i < len(l.Main) {
			l.Main[i] = v
		}
	}
}

// SetMeterLevel stores the level of one console meter. It may be called from
// another goroutine than the one running the server loop (that is what the
// bridge does).
func (s *State) SetMeterLevel(k MeterKind, i int, v float32) {
	s.meterMu.Lock()
	s.levels.set(k, i, v)
	s.meterMu.Unlock()
}

// SetMeterLevels replaces every level at once.
func (s *State) SetMeterLevels(l MeterLevels) {
	s.meterMu.Lock()
	s.levels = l
	s.meterMu.Unlock()
}

// MeterLevel reads one level back.
func (s *State) MeterLevel(k MeterKind, i int) float32 {
	s.meterMu.Lock()
	defer s.meterMu.Unlock()
	return s.levels.get(k, i)
}

// MeterLevels returns a copy of the levels in use.
func (s *State) MeterLevels() MeterLevels {
	s.meterMu.Lock()
	defer s.meterMu.Unlock()
	return s.levels
}

// ---------------------------------------------------------------- frame layout

// meterSlot is the source of one value of a meter frame. MeterNone means "no
// source": the value stays at zero, which is what the C emulator always reports.
type meterSlot struct {
	kind MeterKind
	idx  int
}

// meterLayout tells which level feeds which value of every /meters/NN frame. The
// number of values per group is fixed by Xprepmeter() in X32.c; the meaning of
// the values comes from the two reference sources that decode them:
//
//   - X32Automix.c reads the level of channel ch (0 based) as the float at byte
//     offset 24 + 4*ch of a /meters/1 frame: the group carries the 32 channel
//     levels first and then two blocks of gain reductions, which no bridge can
//     supply and which therefore stay at zero.
//   - Xdump.c fixes the byte order (the value count is little-endian, the blob
//     size is big-endian and the floats are little-endian) and shows that
//     /meters/15 is the RTA and /meters/16 the gain reductions, both sent as 16
//     bit samples rather than floats: those two groups are left at zero too.
//
// The remaining groups are only known by their size. The layouts below are the
// ones that fit the X32 structure exactly (16 values for the 16 buses, 6 for the
// 6 matrices, 8+8+6 = 22 for aux inputs, effect returns and matrices,
// 32+8+8+16+6 = 70 and 32+8+8+16+6+2+8 = 80 for "every fader"): they are marked
// as inferred in docs/X32_OSC_PROTOCOL.md. The group whose layout is unknown
// (/meters/4, 82 values) is left empty, i.e. at zero.
var meterLayout = buildMeterLayout()

// buildMeterLayout builds the table described above.
func buildMeterLayout() [MaxMeters][]meterSlot {
	var l [MaxMeters][]meterSlot

	// series builds n consecutive values of one family.
	series := func(k MeterKind, n int) []meterSlot {
		out := make([]meterSlot, n)
		for i := range out {
			out[i] = meterSlot{kind: k, idx: i}
		}
		return out
	}
	// join concatenates the blocks of a group.
	join := func(parts ...[]meterSlot) []meterSlot {
		var out []meterSlot
		for _, p := range parts {
			out = append(out, p...)
		}
		return out
	}
	// pad grows a group to n values with unused slots (the gain reductions of the
	// same elements, and the trailing spare values): these stay at zero because
	// no outside source can supply them.
	pad := func(s []meterSlot, n int) []meterSlot {
		for len(s) < n {
			s = append(s, meterSlot{kind: MeterNone})
		}
		return s
	}

	ch, aux := series(MeterCh, 32), series(MeterAuxin, 8)
	fxr, bus := series(MeterFxrtn, 8), series(MeterBus, 16)
	mtx, main3 := series(MeterMtx, 6), series(MeterMain, 3)

	// Every group lists its levels block-wise: the level of element n sits at
	// index n of its block. X32Automix.c confirms it for the channels, where the
	// level of channel ch is the float at byte 24 + 4*ch, i.e. value index ch.
	l[0] = join(ch, aux, fxr, bus, mtx)                                 // 70: every fader meter
	l[1] = pad(ch, 96)                                                  // 96: the 32 channel levels
	l[2] = pad(bus, 49)                                                 // 49: the 16 bus levels
	l[3] = join(aux, fxr, mtx)                                          // 22: 8 + 8 + 6
	l[5] = join(pad(aux, 24), main3)                                    // 27: 8 aux levels + main
	l[6] = pad(main3, 4)                                                // 4: main
	l[7] = bus                                                          // 16: buses
	l[8] = mtx                                                          // 6: matrices
	l[9] = ch                                                           // 32: channels
	l[10] = ch                                                          // 32: channels
	l[11] = pad(main3, 5)                                               // 5: main
	l[12] = pad(main3, 4)                                               // 4: main
	l[13] = pad(bus, 48)                                                // 48: the 16 bus levels
	l[14] = pad(join(ch, aux, fxr, bus, mtx, series(MeterMain, 2)), 80) // 80: every fader
	// l[4] (82 values), l[15] (RTA) and l[16] (gain reductions) stay empty.
	return l
}

// fillMeterFrame writes the levels in use into the value area of a meter frame.
// Every value is a 32 bit little-endian float (Xdump.c: "floats are
// little-endian format"), which is also what client parsers assume.
func (s *State) fillMeterFrame(i int, m *meter) {
	if i < 0 || i >= MaxMeters || len(m.buf) == 0 || len(meterLayout[i]) == 0 {
		return
	}
	levels := s.MeterLevels()
	for j, sl := range meterLayout[i] {
		if sl.kind == MeterNone {
			continue
		}
		off := m.data + 4*j
		if off+4 > len(m.buf) {
			return
		}
		binary.LittleEndian.PutUint32(m.buf[off:], math.Float32bits(levels.get(sl.kind, sl.idx)))
	}
}
