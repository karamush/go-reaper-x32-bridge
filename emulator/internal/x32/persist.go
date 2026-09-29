package x32

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ResourceFileName is the file used to keep the console state between two runs
// (".X32res.rc" in X32.c).
const ResourceFileName = ".X32res.rc"

// persistTables returns the tables saved by X32Shutdown(), in the very same
// order as X32.c (this order is part of the file format).
func persistTables() []Table {
	t := []Table{Xconfig, Xmain, Xprefs, Xstat}
	t = append(t, Xchannelset...)
	t = append(t, Xauxinset...)
	t = append(t, Xfxrtnset...)
	t = append(t, Xbusset...)
	t = append(t, Xmtxset...)
	t = append(t, Xdca)
	t = append(t, Xfxset...)
	t = append(t, Xoutput)
	t = append(t, Xheadmpset...)
	t = append(t, Xmisc, Xurec, Xlibsc, Xlibsr, Xlibsf)
	return t
}

// X32Shutdown saves the console state and asks the main loop to stop
// (X32Shutdown/funct_shutdown in X32.c).
func (s *State) X32Shutdown() int {
	if err := s.SaveResource(ResourceFileName); err != nil {
		s.logf("cannot save %s: %v\n", ResourceFileName, err)
		return 0
	}
	s.logf("saving init file... Done\n")
	s.Shutdown = true
	s.Stop = true
	return 0
}

// SaveResource writes every table to a file. Integers and floats are stored as
// decimal integers (floats keep their bit pattern, exactly like X32.c which
// writes "%d" of value.ii), strings are stored as "<len> <text>".
func (s *State) SaveResource(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, tab := range persistTables() {
		for i := range tab {
			c := &tab[i]
			if c.Type == TS32 {
				v := c.Str()
				fmt.Fprintf(w, "%d %d %s\n", c.Type, len(v), v)
				continue
			}
			bits := c.ValI
			if c.Type == TF32 {
				bits = int32(math.Float32bits(c.ValF))
			}
			fmt.Fprintf(w, "%d %d\n", c.Type, bits)
		}
	}
	return nil
}

// ResetValues clears every value of every persistent table (used by tests and
// by the initial state of a fresh process).
func ResetValues() {
	for _, tab := range persistTables() {
		for i := range tab {
			tab[i].ValI = 0
			tab[i].ValF = 0
			tab[i].ValS = nil
		}
	}
}

// X32Init loads the console state from a file. It returns true when the file
// does not exist (all values keep their zero defaults), like X32Init in X32.c.
func (s *State) X32Init(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return true
	}
	s.logf("Reading init file...")
	rd := bufio.NewReader(strings.NewReader(string(data)))
	for _, tab := range persistTables() {
		for i := range tab {
			c := &tab[i]
			typ, err := readInt(rd)
			if err != nil {
				s.logf(" Done\n")
				return false
			}
			if Type(typ) == TS32 {
				n, err := readInt(rd)
				if err != nil {
					s.logf(" Done\n")
					return false
				}
				if n > 0 {
					buf := make([]byte, n)
					if _, err := rd.Read(buf); err != nil {
						s.logf(" Done\n")
						return false
					}
					_, _ = rd.ReadByte() // the trailing newline
					c.SetStr(string(buf))
				} else {
					c.ValS = nil
				}
				continue
			}
			v, err := readInt(rd)
			if err != nil {
				s.logf(" Done\n")
				return false
			}
			if Type(typ) == TF32 {
				c.ValF = math.Float32frombits(uint32(v))
			} else {
				c.ValI = v
			}
		}
	}
	s.logf(" Done\n")
	return false
}

// readInt reads one decimal integer, skipping white space.
func readInt(rd *bufio.Reader) (int32, error) {
	var sb strings.Builder
	for {
		b, err := rd.ReadByte()
		if err != nil {
			if sb.Len() > 0 {
				break
			}
			return 0, err
		}
		if b == ' ' || b == '\n' || b == '\r' || b == '\t' {
			if sb.Len() > 0 {
				break
			}
			continue
		}
		sb.WriteByte(b)
	}
	v, err := strconv.ParseInt(sb.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}
