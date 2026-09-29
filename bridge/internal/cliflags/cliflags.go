// Package cliflags defines the command line flags shared by the commands of this
// module: x32bridge (the bridge alone, for a real console or a separately started
// emulator) and x32reaper (the emulator and the bridge in one process).
//
// The defaults come from bridge.DefaultConfig(), so the two commands accept the
// same tuning knobs with the same names and the same values.
package cliflags

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"x32bridge/internal/bridge"
)

// maskNames maps the class names of the -tox32 / -toreaper flags to their filter
// bit (the same ten classes as X32ReaperW).
var maskNames = map[string]uint16{
	"pan":    bridge.BitPan,
	"fader":  bridge.BitFader,
	"name":   bridge.BitName,
	"mute":   bridge.BitMute,
	"select": bridge.BitSelect,
	"send":   bridge.BitSend,
	"solo":   bridge.BitSolo,
	"fx":     bridge.BitFX,
	"mpan":   bridge.BitMPan,
	"mfader": bridge.BitMFader,
}

// parseMask turns "all", "none" or a comma separated list of class names into a
// filter mask.
func parseMask(s string) (uint16, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	switch s {
	case "all", "":
		return bridge.AllBits, nil
	case "none":
		return 0, nil
	}
	var mask uint16
	for _, name := range strings.Split(s, ",") {
		name = strings.TrimSpace(name)
		bit, ok := maskNames[name]
		if !ok {
			return 0, fmt.Errorf("unknown class %q", name)
		}
		mask |= bit
	}
	return mask, nil
}

// parseRanges turns "1-4,5-8,9-12,13-16" into the eight DCA track ranges.
func parseRanges(s string, lo, hi *[8]int) error {
	parts := strings.Split(s, ",")
	for i := 0; i < 8; i++ {
		if i >= len(parts) {
			break
		}
		p := strings.TrimSpace(parts[i])
		if p == "" || p == "-" {
			continue
		}
		fields := strings.SplitN(p, "-", 2)
		a, err := strconv.Atoi(strings.TrimSpace(fields[0]))
		if err != nil {
			return fmt.Errorf("range %q: %v", p, err)
		}
		b := a
		if len(fields) == 2 {
			if b, err = strconv.Atoi(strings.TrimSpace(fields[1])); err != nil {
				return fmt.Errorf("range %q: %v", p, err)
			}
		}
		if a < 1 || b < a {
			return fmt.Errorf("range %q is not usable", p)
		}
		lo[i], hi[i] = a, b
	}
	return nil
}

type Extra struct {
	toX32    string
	toReaper string
	rdca     string
}

// Register defines every bridge flag on fs. The values are written straight into
// cfg, so the caller only has to call Apply afterwards (for the text flags).
func Register(fs *flag.FlagSet, cfg *bridge.Config) *Extra {
	fs.StringVar(&cfg.X32Addr, "x32", cfg.X32Addr, "console/emulator address host:port")
	fs.StringVar(&cfg.Listen, "listen", cfg.Listen, "local address REAPER sends to (host:port)")
	fs.StringVar(&cfg.Host, "host", cfg.Host, "REAPER host (where we send)")
	fs.StringVar(&cfg.Port, "port", cfg.Port, "REAPER OSC listen port (where we send)")
	fs.IntVar(&cfg.LocalPort, "localport", cfg.LocalPort, "pin the local port of the console side socket (0 = ephemeral)")

	fs.IntVar(&cfg.BankSize, "bank", cfg.BankSize, "channel bank size (8, 16 or 32)")
	fs.IntVar(&cfg.BankOffset, "bankof", cfg.BankOffset, "initial channel bank")
	fs.IntVar(&cfg.SendOffset, "sendoff", cfg.SendOffset, "REAPER track send offset")

	fs.IntVar(&cfg.TrkMin, "trkmin", cfg.TrkMin, "first REAPER track of the channel bank")
	fs.IntVar(&cfg.TrkMax, "trkmax", cfg.TrkMax, "last REAPER track of the channel bank")
	fs.IntVar(&cfg.AuxMin, "auxmin", cfg.AuxMin, "first REAPER track mapped to auxin 1")
	fs.IntVar(&cfg.AuxMax, "auxmax", cfg.AuxMax, "last REAPER track mapped to auxin 8")
	fs.IntVar(&cfg.FxrMin, "fxrmin", cfg.FxrMin, "first REAPER track mapped to fxrtn 1")
	fs.IntVar(&cfg.FxrMax, "fxrmax", cfg.FxrMax, "last REAPER track mapped to fxrtn 8")
	fs.IntVar(&cfg.BusMin, "busmin", cfg.BusMin, "first REAPER track mapped to bus 1")
	fs.IntVar(&cfg.BusMax, "busmax", cfg.BusMax, "last REAPER track mapped to bus 16")
	fs.IntVar(&cfg.DcaMin, "dcamin", cfg.DcaMin, "first REAPER track of DCA 1")
	fs.IntVar(&cfg.DcaMax, "dcamax", cfg.DcaMax, "last REAPER track of DCA 1")

	fs.BoolVar(&cfg.Transport, "transport", cfg.Transport, "Bank C buttons/encoders drive the REAPER transport")
	fs.BoolVar(&cfg.ChBank, "chbank", cfg.ChBank, "Bank C drives the channel bank instead of the transport")
	fs.BoolVar(&cfg.Master, "master", cfg.Master, "mirror the stereo master bus")
	fs.IntVar(&cfg.BankUp, "bankup", cfg.BankUp, "Bank C button (5..12) for bank up when -chbank")
	fs.IntVar(&cfg.BankDn, "bankdn", cfg.BankDn, "Bank C button (5..12) for bank down when -chbank")
	fs.IntVar(&cfg.MarkerButton, "markerbtn", cfg.MarkerButton, "Bank C button (5..12) inserting a marker when -transport=false")

	fs.BoolVar(&cfg.Verbose, "v", cfg.Verbose, "log every datagram exchanged with REAPER and the console")
	fs.IntVar(&cfg.DelayBank, "delayb", cfg.DelayBank, "ms between the messages of a bank recall")
	fs.IntVar(&cfg.DelayGen, "delayg", cfg.DelayGen, "ms between ordinary messages")

	e := &Extra{}
	fs.StringVar(&e.toReaper, "toreaper", "all", "console -> REAPER classes: all, none or a list (pan,fader,name,mute,select,send,solo,fx,mpan,mfader)")
	fs.StringVar(&e.toX32, "tox32", "all", "REAPER -> console classes (same names)")
	fs.StringVar(&e.rdca, "rdca", "", "REAPER track ranges per DCA, e.g. \"1-4,5-8,9-12,13-16\"")
	return e
}

// Apply converts the text flags into their configuration values.
func (e *Extra) Apply(cfg *bridge.Config) error {
	var err error
	if cfg.ToX32, err = parseMask(e.toX32); err != nil {
		return fmt.Errorf("-tox32: %w", err)
	}
	if cfg.ToReaper, err = parseMask(e.toReaper); err != nil {
		return fmt.Errorf("-toreaper: %w", err)
	}
	if e.rdca != "" {
		if err := parseRanges(e.rdca, &cfg.RdcaMin, &cfg.RdcaMax); err != nil {
			return fmt.Errorf("-rdca: %w", err)
		}
	}
	return nil
}

// WasSet reports whether a flag was given on the command line (used to know
// whether the default of -x32 still applies).
func WasSet(fs *flag.FlagSet, name string) bool {
	found := false
	fs.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
