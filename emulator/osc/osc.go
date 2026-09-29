// Package osc implements the OSC 1.0 subset used by Behringer X32 consoles,
// exactly as the reference emulator handles it (X32.c v0.88 of the upstream
// X32-Behringer project, https://github.com/pmaillot/X32-Behringer):
//
//   - address + NUL, padded to a 4 byte boundary
//   - optional type tag string ",ifsb"; when it is missing the message is
//     interpreted as a GET request ("old" notation). An empty tag string (",")
//     is also accepted and means "no arguments"
//   - i / f are 32 bit big-endian, s is NUL terminated and padded, b is a
//     big-endian 32 bit size followed by padded data
//   - no bundles, no 64 bit types, no true/false/nil tags
package osc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

// Arg is a single decoded OSC argument.
type Arg struct {
	Tag byte
	I   int32
	F   float32
	S   string
	B   []byte
}

// Message is a decoded OSC message.
type Message struct {
	Addr string // address, without padding (as received, NUL terminated)
	Tags string // tag characters without the leading comma ("" for old notation)
	Args []Arg  // decoded arguments, one per tag
}

func pad4(n int) int { return (n + 3) &^ 3 }

// Decode decodes a datagram. It never panics and never fails: truncated or
// unknown constructs simply stop the decoding (like the C emulator, which
// ignores anything it cannot parse).
func Decode(b []byte) *Message {
	m := &Message{}
	if len(b) == 0 {
		return m
	}
	z := bytes.IndexByte(b, 0)
	if z < 0 {
		m.Addr = string(b)
		return m
	}
	m.Addr = string(b[:z])
	p := pad4(z + 1)
	if p >= len(b) || b[p] != ',' {
		return m // old notation: no type tag string at all
	}
	q := bytes.IndexByte(b[p:], 0)
	if q < 0 {
		q = len(b)
	} else {
		q += p
	}
	m.Tags = string(b[p+1 : q])
	d := pad4(q + 1)
	for i := 0; i < len(m.Tags); i++ {
		switch m.Tags[i] {
		case 'i':
			if d+4 > len(b) {
				return m
			}
			m.Args = append(m.Args, Arg{Tag: 'i', I: int32(binary.BigEndian.Uint32(b[d:]))})
			d += 4
		case 'f':
			if d+4 > len(b) {
				return m
			}
			m.Args = append(m.Args, Arg{Tag: 'f', F: math.Float32frombits(binary.BigEndian.Uint32(b[d:]))})
			d += 4
		case 's':
			e := bytes.IndexByte(b[d:], 0)
			if e < 0 {
				e = len(b)
			} else {
				e += d
			}
			m.Args = append(m.Args, Arg{Tag: 's', S: string(b[d:e])})
			d = pad4(e + 1)
		case 'b':
			if d+4 > len(b) {
				return m
			}
			n := int(int32(binary.BigEndian.Uint32(b[d:])))
			d += 4
			if n < 0 || d+n > len(b) {
				return m
			}
			blob := make([]byte, n)
			copy(blob, b[d:d+n])
			m.Args = append(m.Args, Arg{Tag: 'b', B: blob})
			d = pad4(d + n)
		default:
			return m // unknown tag: stop parsing (C behaviour)
		}
	}
	return m
}

// AppendString appends a NUL terminated string padded to 4 bytes (Xsprint 's').
func AppendString(b []byte, s string) []byte {
	b = append(b, s...)
	b = append(b, 0)
	for len(b)&3 != 0 {
		b = append(b, 0)
	}
	return b
}

// AppendInt32 appends a 32 bit big-endian integer (Xsprint 'i').
func AppendInt32(b []byte, v int32) []byte {
	var t [4]byte
	binary.BigEndian.PutUint32(t[:], uint32(v))
	return append(b, t[:]...)
}

// AppendFloat32 appends a 32 bit big-endian float (Xsprint 'f').
func AppendFloat32(b []byte, v float32) []byte {
	return AppendInt32(b, int32(math.Float32bits(v)))
}

// AppendBlob appends a big-endian size and the padded blob data.
func AppendBlob(b []byte, data []byte) []byte {
	b = AppendInt32(b, int32(len(data)))
	b = append(b, data...)
	for len(b)&3 != 0 {
		b = append(b, 0)
	}
	return b
}

// Start builds a message header: address, padding, ",", tag characters, padding.
func Start(addr, tags string) []byte {
	b := make([]byte, 0, len(addr)+len(tags)+8)
	b = AppendString(b, addr)
	b = append(b, ',')
	b = append(b, tags...)
	b = append(b, 0)
	for len(b)&3 != 0 {
		b = append(b, 0)
	}
	return b
}

// Build builds a complete OSC message with the given arguments.
func Build(addr, tags string, args ...Arg) []byte {
	b := Start(addr, tags)
	for _, a := range args {
		switch a.Tag {
		case 'i':
			b = AppendInt32(b, a.I)
		case 'f':
			b = AppendFloat32(b, a.F)
		case 's':
			b = AppendString(b, a.S)
		case 'b':
			b = AppendBlob(b, a.B)
		}
	}
	return b
}

// Dump renders a datagram the way the reference tool Xfdump() does: printable
// characters are kept, unprintable ones become '~' and every argument is shown
// between brackets at the position of its data.
func Dump(header string, b []byte, debug bool) string {
	out := &strings.Builder{}
	if debug {
		for i := 0; i < len(b); i += 16 {
			for j := i; j < i+16 && j < len(b); j++ {
				fmt.Fprintf(out, "%02x ", b[j])
			}
			out.WriteByte('\n')
		}
		return out.String()
	}
	fmt.Fprintf(out, "%s, %4d B: ", header, len(b))
	z := bytes.IndexByte(b, 0)
	if z < 0 {
		return out.String() + "truncated"
	}
	put := func(c byte) {
		if c < 32 || c == 127 || c == 255 {
			c = '~'
		}
		out.WriteByte(c)
	}
	for i := 0; i < z; i++ {
		put(b[i])
	}
	p := pad4(z + 1)
	for i := z; i < p && i < len(b); i++ {
		put(b[i])
	}
	m := Decode(b)
	if p >= len(b) || b[p] != ',' {
		return out.String()
	}
	q := bytes.IndexByte(b[p:], 0)
	if q < 0 {
		q = len(b)
	} else {
		q += p
	}
	for i := p; i < q; i++ {
		out.WriteByte(b[i])
	}
	for i := q; i < pad4(q+1) && i < len(b); i++ {
		put(b[i])
	}
	for _, a := range m.Args {
		switch a.Tag {
		case 'i':
			fmt.Fprintf(out, "[%6d]", a.I)
		case 'f':
			switch {
			case a.F < 10.:
				fmt.Fprintf(out, "[%06.4f]", a.F)
			case a.F < 100.:
				fmt.Fprintf(out, "[%06.3f]", a.F)
			case a.F < 1000.:
				fmt.Fprintf(out, "[%06.2f]", a.F)
			case a.F < 10000.:
				fmt.Fprintf(out, "[%06.1f]", a.F)
			default:
				fmt.Fprintf(out, "[%06f]", a.F)
			}
		case 's':
			out.WriteString(a.S)
			for i := len(a.S) + 1; i < pad4(len(a.S)+1); i++ {
				out.WriteByte('~')
			}
		case 'b':
			if len(a.B) >= 4 {
				n := int(int32(binary.LittleEndian.Uint32(a.B)))
				be := int(int32(binary.BigEndian.Uint32(a.B)))
				if n == be {
					fmt.Fprintf(out, "%d chrs: ", n)
				} else {
					fmt.Fprintf(out, "%3d vals: ", n)
				}
			} else {
				out.WriteString("blob")
			}
		}
	}
	return out.String()
}
