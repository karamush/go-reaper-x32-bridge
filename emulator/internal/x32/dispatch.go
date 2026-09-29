package x32

import (
	"net"
	"strings"

	"x32emu/osc"
)

// header is one entry of the Xheader[] table of X32.c: the first four bytes of
// the received address select the handling function.
type header struct {
	key string
	fn  func(*State) int
}

// headers reproduces the Xheader[] table (order kept for readability only,
// the lookup is an exact four byte comparison). It is built in init() because
// the handlers themselves use it (no initialization cycle).
var headers []header

func init() {
	headers = []header{
		{"/shu", (*State).functionShutdown},
		{"/inf", (*State).functionInfo},
		{"/xin", (*State).functionXinfo},
		{"/sta", (*State).functionStatus},
		{"/xre", (*State).functionXremote},
		{"/nod", (*State).functionNode},
		{"/\x00\x00\x00", (*State).functionSlash},
		{"/con", (*State).functionConfig},
		{"/mai", (*State).functionMain},
		{"/-pr", (*State).functionPrefs},
		{"/-st", (*State).functionStat},
		{"/-ur", (*State).functionUrec},
		{"/ch/", (*State).functionChannel},
		{"/aux", (*State).functionAuxin},
		{"/fxr", (*State).functionFxrtn},
		{"/bus", (*State).functionBus},
		{"/mtx", (*State).functionMtx},
		{"/dca", (*State).functionDca},
		{"/fx/", (*State).functionFx},
		{"/out", (*State).functionOutput},
		{"/hea", (*State).functionHeadamp},
		{"/met", (*State).functionMeters},
		{"/-ha", (*State).functionMisc},
		{"/ins", (*State).functionMisc},
		{"/-sh", (*State).functionShow},
		{"/ren", (*State).functionRenew},
		{"/cop", (*State).functionCopy},
		{"/add", (*State).functionAdd},
		{"/loa", (*State).functionLoad},
		{"/sav", (*State).functionSave},
		{"/del", (*State).functionDelete},
		{"/uns", (*State).functionUnsubscribe},
		{"/-us", (*State).functionMisc},
		{"/und", (*State).functionDummy},
		{"/-ac", (*State).functionAction},
		{"/-li", (*State).functionLibs},
		{"/sho", (*State).functionShowdump},
	}
}

// HandlePacket is the main entry point: it decodes the datagram, dispatches it
// to the subsystem handler and sends whatever the handler prepared. This is the
// body of the main loop of X32.c.
func (s *State) HandlePacket(b []byte, from *net.UDPAddr) {
	s.Raw = b
	s.Msg = osc.Decode(b)
	s.From = from
	s.Out = s.Out[:0]
	if s.Verbose {
		switch {
		case strings.HasPrefix(s.Msg.Addr, "/xre"):
			s.dump("->X", b, s.VerbRemote)
		case strings.HasPrefix(s.Msg.Addr, "/bat"):
			s.dump("->X", b, s.VerbBatch)
		case strings.HasPrefix(s.Msg.Addr, "/for"):
			s.dump("->X", b, s.VerbFormat)
		case strings.HasPrefix(s.Msg.Addr, "/ren"):
			s.dump("->X", b, s.VerbRenew)
		case strings.HasPrefix(s.Msg.Addr, "/met"):
			s.dump("->X", b, s.VerbMeter)
		default:
			s.dump("->X", b, true)
		}
	}
	who := 0
	key := make([]byte, 4)
	copy(key, b)
	k := string(key)
	for i := range headers {
		if headers[i].key == k {
			who = headers[i].fn(s)
			break
		}
	}
	s.Xsend(who)
}

// lookup searches a table for an exact address match and parses it
// (the strcmp() loop of the subsystem handlers of X32.c).
func (s *State) lookup(tab Table) int {
	addr := s.Msg.Addr
	for i := range tab {
		if tab[i].Addr == addr {
			return s.functParams(tab, i)
		}
	}
	return 0
}

// lookupFirst searches the given tables in order and parses the entry that holds
// the address. It is used where one subsystem is served by more than one table
// (X32.c only ever has one): the address decides which table is used, so a
// "silent" answer coming from the first table can never leak into the next one.
func (s *State) lookupFirst(tabs ...Table) int {
	addr := s.Msg.Addr
	for _, tab := range tabs {
		for i := range tab {
			if tab[i].Addr == addr {
				return s.functParams(tab, i)
			}
		}
	}
	return 0
}

// digit returns the decimal digit of the current packet at offset off, -1 when
// the packet is too short (the C code would read whatever is in the buffer).
func (s *State) digit(off int) int {
	if off >= len(s.Raw) {
		return -1
	}
	c := s.Raw[off]
	if c < '0' || c > '9' {
		return -1
	}
	return int(c - '0')
}

func (s *State) numbered(offs []int) int {
	v := 0
	for _, o := range offs {
		d := s.digit(o)
		if d < 0 {
			return -1
		}
		v = v*10 + d
	}
	return v
}

// ---------------------------------------------------------------- subsystems
func (s *State) functionConfig() int  { return s.lookup(Xconfig) }
func (s *State) functionMain() int    { return s.lookup(Xmain) }
func (s *State) functionPrefs() int   { return s.lookup(Xprefs) }
func (s *State) functionUrec() int    { return s.lookup(Xurec) }
func (s *State) functionDca() int     { return s.lookup(Xdca) }
func (s *State) functionOutput() int  { return s.lookup(Xoutput) }
func (s *State) functionMisc() int    { return s.lookup(Xmisc) }
func (s *State) functionAction() int  { return s.lookup(Xaction) }
func (s *State) functionDummy() int   { return 0 } // /undo: prints and does nothing
func (s *State) functionLibs() int    { return 0 } // /-libs: not implemented
func (s *State) functionRenew() int   { return SSnd }
func (s *State) functionChannel() int { return s.indexed(s.numbered([]int{4, 5}), 32, Xchannelset) }
func (s *State) functionAuxin() int   { return s.indexed(s.numbered([]int{7, 8}), 8, Xauxinset) }
func (s *State) functionFxrtn() int   { return s.indexed(s.numbered([]int{7, 8}), 8, Xfxrtnset) }
func (s *State) functionBus() int     { return s.indexed(s.numbered([]int{5, 6}), 16, Xbusset) }
func (s *State) functionMtx() int     { return s.indexed(s.numbered([]int{5, 6}), 6, Xmtxset) }
func (s *State) functionHeadamp() int {
	return s.indexed0(s.numbered([]int{9, 10, 11}), 128, Xheadmpset)
}

// indexed resolves a numbered subsystem into its table (1 based number).
func (s *State) indexed(n, max int, set []Table) int {
	if n < 1 || n > max {
		return 0
	}
	return s.lookup(set[n-1])
}

// indexed0 is the 0 based variant used by /headamp/NNN (X32.c uses the number
// itself as the index of Xheadmpset[]).
func (s *State) indexed0(n, max int, set []Table) int {
	if n < 0 || n >= max {
		return 0
	}
	return s.lookup(set[n])
}

// functionFx handles /fx/<n>/... (1 to 8).
func (s *State) functionFx() int {
	n := s.digit(4)
	if n < 1 || n > 8 {
		return 0
	}
	return s.lookup(Xfxset[n-1])
}

// functionStat handles /-stat/... and the (non standard) shutdown request
// "/-stat/lock ,i 2" (byte 19, like /-stat/lock ,i 2 in X32.c).
//
// The hand written Xuserpar table (user assign controls of a real console) is
// not part of the C tables, so it is looked up separately - and after Xstat, so
// that the addresses shared with the reference emulator keep their C behaviour.
func (s *State) functionStat() int {
	if strings.Contains(s.Msg.Addr, "/lock") && len(s.Raw) > 19 && s.Raw[19] == 2 {
		return s.X32Shutdown()
	}
	return s.lookupFirst(Xstat, Xuserpar)
}

// functionShow handles /-show/showfile/scene/..., /-show/showfile/snippet/...
// and the other /-show/... commands. X32.c tests the address starting at
// offset 16 ("/-show/showfile/" is 16 characters long).
func (s *State) functionShow() int {
	addr := s.Msg.Addr
	if len(addr) >= 21 && strings.HasPrefix(addr[16:], "scene") {
		return s.lookup(Xscene)
	}
	if len(addr) >= 23 && strings.HasPrefix(addr[16:], "snippet") {
		return s.lookup(Xsnippet)
	}
	return s.lookup(Xshow)
}
