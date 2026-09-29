package x32

import (
	"strings"

	"x32emu/osc"
)

// Copy masks of /copy (C_HA ... C_SEND in X32.c).
const (
	cHA     = 0x0001
	cConfig = 0x0002
	cGate   = 0x0004
	cDyn    = 0x0008
	cEQ     = 0x0010
	cSend   = 0x0020
)

// functionCopy handles "/copy ,siii libchan <src> <dst> <mask>" (only the
// channel library is implemented, like in X32.c). It answers
// "/copy ,si libchan 1" on success and "... 0" on error.
func (s *State) functionCopy() int {
	if s.argS(0) != "libchan" {
		return 0
	}
	src, dst, mask := int(s.argI(1)), int(s.argI(2)), s.argI(3)
	ok := false
	if src >= 0 && src < 32 && dst >= 0 && dst < 32 {
		ok = true
		from, to := Xchannelset[src], Xchannelset[dst]
		copies := []struct {
			bit   int32
			first int
			last  int
		}{
			{cConfig, 0, 8}, {cHA, 8, 20}, {cGate, 20, 33},
			{cDyn, 33, 53}, {cEQ, 53, 75}, {cSend, 75, 146},
		}
		// NOTE: X32.c copies value.ii only, which corrupts string entries;
		// here every entry is copied according to its type.
		for _, cpy := range copies {
			if mask&cpy.bit == 0 {
				continue
			}
			for i := cpy.first; i < cpy.last && i < len(from) && i < len(to); i++ {
				to[i].ValI, to[i].ValF = from[i].ValI, from[i].ValF
				to[i].ValS = nil
				if from[i].ValS != nil {
					v := *from[i].ValS
					to[i].ValS = &v
				}
			}
		}
	}
	res := int32(0)
	if ok {
		res = 1
	}
	b := osc.Start("/copy", "si")
	b = osc.AppendString(b, "libchan")
	b = osc.AppendInt32(b, res)
	s.Out = b
	return SSnd
}

// statusReply builds a "/<cmd> ,si <type> <ok>" status answer.
func statusReply(cmd, typ string, ok int32) []byte {
	b := osc.Start("/"+cmd, "si")
	b = osc.AppendString(b, typ)
	b = osc.AppendInt32(b, ok)
	return b
}

// functionAdd handles /add (cue only, like X32.c: nothing is really added).
func (s *State) functionAdd() int {
	if s.argS(0) != "cue" {
		return 0
	}
	s.Out = statusReply("add", "cue", 1)
	return SSnd
}

// functionLoad handles /load (nothing is really loaded, like X32.c).
func (s *State) functionLoad() int {
	s.Out = statusReply("load", "libchan", 1)
	return SSnd
}

// functionSave handles /save for scene, snippet and libchan.
func (s *State) functionSave() int {
	typ := s.argS(0)
	num := s.argI(1)
	name, notes := s.argS(2), s.argS(3)
	switch typ {
	case "scene":
		tab := Xscene
		label := "/-show/showfile/scene"
		idx := findEntry(tab, label, num, "name")
		if idx < 0 {
			s.Out = statusReply("save", "scene", 0)
			return SSnd
		}
		s.storeString(&tab[idx], name)
		s.storeString(&tab[idx+1], notes)
		// The third entry after "name" is "safes" (X32.c writes Xscene[j+2] and
		// notifies it, although its comment calls it "hasdata" - that flag is at
		// j+3): the value and the notification are part of the protocol.
		s.setIntNotify(&tab[idx+2], 1)
		s.Out = statusReply("save", "scene", 1)
		return SSnd
	case "snippet":
		tab := Xsnippet
		label := "/-show/showfile/snippet"
		idx := findEntry(tab, label, num, "name")
		if idx < 0 {
			s.Out = statusReply("save", "snippet", 0)
			return SSnd
		}
		s.storeString(&tab[idx], name)
		// Xsnippet[j+1] is "eventtyp", not "hasdata" (X32.c comment is wrong)
		s.setIntNotify(&tab[idx+1], 1)
		// the two last parameters (name, value) are not used, like X32.c
		s.Out = statusReply("save", "snippet", 1)
		return SSnd
	case "libchan":
		s.Out = statusReply("save", "libchan", 1)
		return SSnd
	}
	return 0
}

// functionDelete handles /delete for scene, snippet and libchan.
func (s *State) functionDelete() int {
	typ := s.argS(0)
	num := s.argI(1)
	switch typ {
	case "scene":
		idx := findEntry(Xscene, "/-show/showfile/scene", num, "name")
		if idx < 0 {
			s.Out = statusReply("delete", "scene", 0)
			return SSnd
		}
		s.clearString(&Xscene[idx])
		s.clearString(&Xscene[idx+1])
		s.setIntNotify(&Xscene[idx+2], 0)
		s.Out = statusReply("delete", "scene", 1)
		return SSnd
	case "snippet":
		idx := findEntry(Xsnippet, "/-show/showfile/snippet", num, "name")
		if idx < 0 {
			s.Out = statusReply("delete", "snippet", 0)
			return SSnd
		}
		s.clearString(&Xsnippet[idx])
		s.setIntNotify(&Xsnippet[idx+1], 0)
		s.Out = statusReply("delete", "snippet", 1)
		return SSnd
	case "libchan":
		s.Out = statusReply("delete", "libchan", 1)
		return SSnd
	}
	return 0
}

// findEntry locates "<prefix>/<nnn>/<leaf>" in a table.
func findEntry(tab Table, prefix string, num int32, leaf string) int {
	if num < 0 || num > 999 {
		return -1
	}
	addr := prefix + "/" + pad3(int(num)) + "/" + leaf
	for i := range tab {
		if tab[i].Addr == addr {
			return i
		}
	}
	return -1
}

// storeString sets a string value and notifies the other xremote clients
// (Xfprint + Xsend(S_REM) in X32.c).
func (s *State) storeString(c *Command, v string) {
	v = strings.TrimPrefix(strings.TrimSuffix(v, "\""), "\"")
	if c.Str() == v {
		return
	}
	c.SetStr(v)
	s.NotifyValue(c)
}

// clearString resets a string value and notifies the other xremote clients.
func (s *State) clearString(c *Command) {
	if c.ValS == nil {
		return
	}
	c.SetStr("")
	s.NotifyValue(c)
}

// setIntNotify stores an integer value and notifies the other xremote clients.
// X32.c does this unconditionally for the "safes"/"eventtyp" flags of a saved or
// deleted scene/snippet (value assignment followed by Xfprint + Xsend(S_REM)),
// so there is no change detection here.
func (s *State) setIntNotify(c *Command, v int32) {
	c.ValI = v
	s.NotifyValue(c)
}
