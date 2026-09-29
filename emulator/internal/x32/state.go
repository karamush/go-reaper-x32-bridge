package x32

import (
	"fmt"
	"net"
	"sync"
	"time"

	"x32emu/osc"
)

// client is one /xremote client (X32Client[] in X32.c). Clients are identified
// by IP and UDP port, like the sa_data comparison of the C code.
type client struct {
	valid bool
	addr  *net.UDPAddr
	exp   time.Time // X32Client[].xrem
}

// State holds everything the emulator needs: the command tables (which are the
// console state, exactly like in X32.c), the current packet buffers, the
// xremote clients and the meter engine.
type State struct {
	// Verbosity (mirrors -v -d -x -b -f -r -m of X32.c).
	Verbose    bool
	Debug      bool
	VerbRemote bool
	VerbBatch  bool
	VerbFormat bool
	VerbRenew  bool
	VerbMeter  bool

	IP   string
	Port string
	// Name is the console name given at startup (-name). The name stored in
	// "/-prefs/name" (resource file or a client) has priority over it.
	Name string
	// BindIP is the local address the UDP socket is bound to. It may differ from
	// IP (which is the address advertised by /info, /xinfo and /status): with
	// "0.0.0.0" the emulator listens on every interface but still advertises a
	// usable address.
	BindIP string
	Out    []byte       // s_buf equivalent
	Raw    []byte       // r_buf equivalent
	Msg    *osc.Message // decoded form of Raw
	cur    *cursor      // input cursor of the current "/" command
	From   *net.UDPAddr // client the current packet came from (Client_ip_pt)
	SendFn func([]byte, *net.UDPAddr) error
	Logf   func(string) // verbose log sink

	clients [MaxClients]client
	meters  [MaxMeters]meter

	// meterMu guards levels: the levels are fed by an outside source (the REAPER
	// bridge) which runs in another goroutine than the server loop.
	meterMu sync.Mutex
	levels  MeterLevels

	// nodeSingleCmd is the command handled by the last functParams() call; it is
	// used by functionNodeSingle() (node_single_command in X32.c).
	nodeSingleCmd *Command

	Shutdown bool
	Stop     bool
}

// NewState returns a state with default verbosity settings and no clients.
func NewState() *State {
	s := &State{
		Verbose: true,
		Port:    UDPPort,
		BindIP:  "0.0.0.0",
		Logf:    func(l string) { fmt.Print(l) },
	}
	now := time.Now()
	for i := range s.meters {
		s.meters[i].stop = now
		s.meters[i].next = now
		s.meters[i].delta = 50 * time.Millisecond
	}
	return s
}

// logf writes a verbose line (no-op when verbosity is off).
func (s *State) logf(format string, args ...any) {
	if s.Verbose && s.Logf != nil {
		s.Logf(fmt.Sprintf(format, args...))
	}
}

// dump logs a packet with the reference "->X" / "X->" prefixes.
func (s *State) dump(header string, b []byte, on bool) {
	if !on {
		return
	}
	s.logf("%s\n", osc.Dump(header, b, s.Debug))
}

// Send sends a datagram (SendFn is provided by the UDP server).
func (s *State) Send(b []byte, to *net.UDPAddr) error {
	if s.SendFn == nil || to == nil {
		return nil
	}
	return s.SendFn(b, to)
}

// Xsend mirrors Xsend() of X32.c: SSnd answers the requesting client, SRem
// pushes the current buffer to the other /xremote registered clients.
func (s *State) Xsend(who int) {
	if len(s.Out) == 0 {
		return
	}
	if who&SSnd != 0 {
		s.dump("X->", s.Out, s.Verbose)
		_ = s.Send(s.Out, s.From)
	}
	if who&SRem != 0 {
		s.notify(s.Out)
	}
}

// notify pushes a message to every registered /xremote client but the one the
// current request comes from (Xsend(S_REM) in X32.c).
func (s *State) notify(b []byte) {
	s.pushClients(b, s.From)
}

// NotifyAll pushes a message to every registered /xremote client. It is used for
// console side events (meter frames, user assign controls), which have no
// originating client and therefore must not exclude anybody.
func (s *State) NotifyAll(b []byte) {
	s.pushClients(b, nil)
}

// pushClients sends b to every valid /xremote client, skipping exclude when it
// is not nil.
func (s *State) pushClients(b []byte, exclude *net.UDPAddr) {
	now := time.Now()
	for i := range s.clients {
		c := &s.clients[i]
		if !c.valid || c.exp.Before(now) {
			continue
		}
		if exclude != nil && c.addr != nil && c.addr.String() == exclude.String() {
			continue // never echo back to the sender
		}
		s.dump("X->", b, s.Verbose)
		_ = s.Send(b, c.addr)
	}
}

// RegisterClient handles /xremote: refresh or add a client, replacing outdated
// entries (function_xremote in X32.c). It never answers.
func (s *State) RegisterClient(a *net.UDPAddr) {
	now := time.Now()
	exp := now.Add(XRemoteTime * time.Second)
	for i := range s.clients {
		if s.clients[i].valid && s.clients[i].addr != nil && a != nil &&
			s.clients[i].addr.String() == a.String() {
			s.clients[i].exp = exp
			return
		}
	}
	for i := range s.clients {
		if !s.clients[i].valid || s.clients[i].exp.Before(now) {
			s.clients[i] = client{valid: true, addr: a, exp: exp}
			return
		}
	}
	// no room for a new client (X32.c silently drops it)
}

// UnregisterClient handles /unsubscribe.
func (s *State) UnregisterClient(a *net.UDPAddr) {
	if a == nil {
		return
	}
	for i := range s.clients {
		if s.clients[i].valid && s.clients[i].addr != nil && s.clients[i].addr.String() == a.String() {
			s.clients[i].valid = false
			return
		}
	}
}

// ClientCount returns the number of currently valid xremote clients.
func (s *State) ClientCount() int {
	now := time.Now()
	n := 0
	for i := range s.clients {
		if s.clients[i].valid && s.clients[i].exp.After(now) {
			n++
		}
	}
	return n
}
