// Package emu runs the X32 emulator of this workspace as a library, so that one
// program can be both the console and the bridge to REAPER (the x32reaper command
// of the bridge module does exactly that). It is a thin wrapper over the internal
// x32 package: command line tools and tests use it instead of touching the
// emulator internals, which keeps the socket handling and the console state in
// one place.
package emu

import (
	"net"
	"strconv"

	"x32emu/internal/x32"
)

// Defaults of a console, re-exported so that a program using this package needs
// no access to the emulator internals.
const (
	DefaultPort     = x32.UDPPort          // 10023
	DefaultName     = x32.DefaultName      // "REAPER"
	DefaultResource = x32.ResourceFileName // ".X32res.rc"
)

// MeterKind selects a family of console meters (see SetMeterLevel).
type MeterKind int

const (
	MeterCh    MeterKind = iota // input channels 1..32
	MeterAuxin                  // aux inputs 1..8
	MeterFxrtn                  // effect returns 1..8
	MeterBus                    // mix buses 1..16
	MeterMtx                    // matrices 1..6
	MeterMain                   // main bus: 1 = L, 2 = R, 3 = mono
)

func (k MeterKind) internal() x32.MeterKind {
	switch k {
	case MeterAuxin:
		return x32.MeterAuxin
	case MeterFxrtn:
		return x32.MeterFxrtn
	case MeterBus:
		return x32.MeterBus
	case MeterMtx:
		return x32.MeterMtx
	case MeterMain:
		return x32.MeterMain
	}
	return x32.MeterCh
}

// Config describes one emulator instance. The zero value is usable and gives the
// defaults of the x32emu command: every interface, port 10023, the name of
// x32.DefaultName and a resource file in the working directory.
type Config struct {
	// BindIP is the local address to bind ("" or "0.0.0.0" = every interface).
	// IP is the address /info, /xinfo and /status advertise; when it is empty the
	// first usable IPv4 address of the machine is used for a wildcard bind.
	// Port is the UDP port ("" = 10023, "0" = any free port, which is handy for
	// tests: Port() and Address() then report the port actually in use).
	BindIP string
	IP     string
	Port   string

	// Name is stored in "/-prefs/name" at startup, so it wins over the name a
	// resource file may carry. Empty keeps the resource file value (which falls
	// back to x32.DefaultName).
	Name string
	// Resource is the file read at startup ("" = none, like a console without a
	// show file: every /-prefs/ip/addr/N is filled with our own address).
	Resource string

	Verbose    bool
	Debug      bool
	VerbRemote bool
	VerbBatch  bool
	VerbFormat bool
	VerbRenew  bool
	VerbMeter  bool

	// Logf receives the protocol log lines (nil = standard output).
	Logf func(string)
}

// Emulator is an X32 emulator: the console state plus its UDP socket.
type Emulator struct {
	st  *x32.State
	srv *x32.Server

	// Fresh is true when no resource file was found at startup: the console then
	// runs on its built-in defaults (X32.c prints a notice in that case).
	Fresh bool
}

// New prepares the console state, loads the resource file and binds the UDP
// socket, so that the address is usable as soon as New returns.
func New(cfg Config) (*Emulator, error) {
	port := cfg.Port
	if port == "" {
		port = x32.UDPPort
	}
	bind := cfg.BindIP
	if bind == "" {
		bind = "0.0.0.0"
	}
	st := x32.NewState()
	st.Port = port
	st.BindIP = bind
	st.Verbose = cfg.Verbose
	st.Debug = cfg.Debug
	st.VerbRemote = cfg.VerbRemote
	st.VerbBatch = cfg.VerbBatch
	st.VerbFormat = cfg.VerbFormat
	st.VerbRenew = cfg.VerbRenew
	st.VerbMeter = cfg.VerbMeter
	st.IP = cfg.IP
	if st.IP == "" {
		if bind == "0.0.0.0" {
			st.IP = FirstIPv4()
		} else {
			st.IP = bind
		}
	}
	if cfg.Logf != nil {
		st.Logf = cfg.Logf
	}
	res := cfg.Resource
	if res == "" {
		res = x32.ResourceFileName
	}
	fresh := st.X32Init(res)
	if fresh {
		// no resource file: feed /-prefs/ip/addr/N with our own address, like
		// X32.c does when it cannot read the file
		st.SetPrefsIP()
	}
	if cfg.Name != "" {
		st.SetName(cfg.Name)
	}
	srv, err := x32.NewServer(bind, st)
	if err != nil {
		return nil, err
	}
	// A port of 0 asks the system for a free one: keep the port in use, so that
	// Address() and Port() are usable right away (tests rely on it).
	if la, ok := srv.Conn.LocalAddr().(*net.UDPAddr); ok && la.Port > 0 {
		st.Port = strconv.Itoa(la.Port)
	}
	return &Emulator{st: st, srv: srv, Fresh: fresh}, nil
}

// Run serves requests until the console is asked to stop (/shutdown) or Stop is
// called. It returns when the server loop ends.
func (e *Emulator) Run() error { return e.srv.Run() }

// Stop asks the server loop to return.
func (e *Emulator) Stop() { e.st.Stop = true }

// Close releases the UDP socket.
func (e *Emulator) Close() error { return e.srv.Close() }

// Address is the address a local client (the bridge) should use to reach the
// emulator: the loopback one when the socket is bound to every interface.
func (e *Emulator) Address() string {
	ip := e.st.BindIP
	if ip == "" || ip == "0.0.0.0" {
		ip = "127.0.0.1"
	}
	return net.JoinHostPort(ip, e.st.Port)
}

// Bound is the local address the UDP socket is bound to.
func (e *Emulator) Bound() string { return e.st.BindIP }

// IP is the address the emulator advertises in /info, /xinfo and /status.
func (e *Emulator) IP() string { return e.st.IP }

// Port is the UDP port in use.
func (e *Emulator) Port() string { return e.st.Port }

// SetMeterLevel stores the level of one console meter. Values are linear
// amplitudes of the console meter scale: 0 is -oo, 1 is 0 dBFS. It is safe to
// call from another goroutine than the one running Run (that is the point: the
// bridge feeds REAPER levels into the meter frames of the console).
func (e *Emulator) SetMeterLevel(k MeterKind, i int, v float32) {
	e.st.SetMeterLevel(k.internal(), i, v)
}

// MeterLevel reads one level back.
func (e *Emulator) MeterLevel(k MeterKind, i int) float32 {
	return e.st.MeterLevel(k.internal(), i)
}

// FirstIPv4 returns the first usable IPv4 address of the machine, or 127.0.0.1
// when there is none. It is what the emulator advertises when it is bound to the
// wildcard address.
func FirstIPv4() string {
	ifs, err := net.Interfaces()
	if err != nil {
		return "127.0.0.1"
	}
	for _, i := range ifs {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, a := range addrs {
			if ipn, ok := a.(*net.IPNet); ok && !ipn.IP.IsLoopback() {
				if v4 := ipn.IP.To4(); v4 != nil {
					return v4.String()
				}
			}
		}
	}
	return "127.0.0.1"
}
