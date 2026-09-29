package x32

import (
	"bytes"
	"encoding/binary"
	"math"

	"x32emu/osc"
)

// functParams is the port of funct_params() of X32.c: it implements both the GET
// and the SET form of any table entry, including change detection and the update
// of the other xremote clients.
//
// Return value: SSnd (answer the requesting client only), SRem (push to the other
// clients only) or 0 (answer nothing - this is what happens when a SET does not
// change anything, or when the command is unknown).
func (s *State) functParams(tab Table, i int) int {
	raw := s.Raw
	// save the command for a possible /node single argument reply
	if i >= 0 && i < len(tab) {
		s.nodeSingleCmd = &tab[i]
	}
	cLen := len(tab[i].Addr)
	fLen := ((cLen + 4) &^ 3) + 1 // first character of the type tag block

	if len(raw)-4 > cLen && fLen < len(raw) && raw[fLen] != 0 {
		// ------------------------------------------------------------ SET form
		if tab[i].Flags&FFND != 0 {
			i++
		}
		if i >= len(tab) || tab[i].Flags&FSET == 0 {
			return 0
		}
		c := fLen
		nTags := 0
		for c < len(raw) && raw[c] != 0 { // count the type tag characters
			nTags++
			c++
		}
		d := (c + 3) &^ 3 // first byte of the data block
		update := false
		for t := 0; t < nTags && i < len(tab); t++ {
			switch raw[fLen+t] {
			case 'i':
				if d+4 > len(raw) {
					return 0
				}
				v := int32(binary.BigEndian.Uint32(raw[d:]))
				d += 4
				if tab[i].ValI != v {
					update = true
					tab[i].ValI = v
				}
			case 'f':
				if d+4 > len(raw) {
					return 0
				}
				v := math.Float32frombits(binary.BigEndian.Uint32(raw[d:]))
				d += 4
				if tab[i].ValF != v {
					update = true
					tab[i].ValF = v
				}
			case 's':
				e := bytes.IndexByte(raw[d:], 0)
				var str string
				if e < 0 {
					str, d = string(raw[d:]), len(raw)
				} else {
					str, d = string(raw[d:d+e]), (d+e+4)&^3
				}
				if str != "" {
					if tab[i].Str() != str {
						update = true
					}
					tab[i].SetStr(str)
				} else if tab[i].ValS != nil {
					tab[i].ValS = nil
					update = true
				}
			case 'b':
				// blobs are accepted but ignored (X32.c only prints them)
			default:
				// unknown tag: stop parsing, like the C switch default
				t = nTags
			}
			i++
		}
		if update {
			s.Out = append(s.Out[:0], raw...)
			return SRem // echo the request to the other xremote clients only
		}
		return 0 // nothing changed: no answer at all
	}

	// ---------------------------------------------------------------- GET form
	// (no type tag block at all, or an empty tag string: "old" OSC notation)
	if tab[i].Flags&FFND != 0 {
		i++ // first of a series gives the first of the next data
	}
	if i >= len(tab) || tab[i].Flags&FGET == 0 {
		return 0
	}
	c := &tab[i]
	s.Out = osc.AppendString(s.Out[:0], c.Addr)
	switch c.Type {
	case TI32, TE32, TP32:
		s.Out = osc.AppendString(s.Out, ",i")
		s.Out = osc.AppendInt32(s.Out, c.ValI)
	case TF32:
		s.Out = osc.AppendString(s.Out, ",f")
		s.Out = osc.AppendFloat32(s.Out, c.ValF)
	case TS32:
		s.Out = osc.AppendString(s.Out, ",s")
		s.Out = osc.AppendString(s.Out, c.Str()) // "" is a single NUL byte, like &zero
	case TFX32:
		// type of an FX parameter is defined by the effect loaded in the slot
		switch s.fxArgType(tab, i) {
		case TI32, TE32:
			s.Out = osc.AppendString(s.Out, ",i")
			s.Out = osc.AppendInt32(s.Out, c.ValI)
		default:
			s.Out = osc.AppendString(s.Out, ",f")
			s.Out = osc.AppendFloat32(s.Out, c.ValF)
		}
	case TB32:
		// TODO: blob replies are not implemented in X32.c either
	default:
		// composed types have no direct value: the address alone is sent
		// (this is what X32.c does when value.str is NULL)
	}
	return SSnd
}
