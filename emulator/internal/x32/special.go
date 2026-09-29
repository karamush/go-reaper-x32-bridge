package x32

import (
	"fmt"
	"x32emu/osc"
)

// SetPrefsIP stores the bound IP address into "/-prefs/ip/addr/0..3"
// (the code at the end of main() in X32.c).
func (s *State) SetPrefsIP() {
	var a, b, c, d int
	if n, err := fmt.Sscanf(s.IP, "%d.%d.%d.%d", &a, &b, &c, &d); err != nil || n != 4 {
		return
	}
	addr := "/-prefs/ip/addr/"
	for i := 0; i < 4; i++ {
		for j := range Xprefs {
			if Xprefs[j].Addr == addr+itoa(i) {
				Xprefs[j].ValI = int32([]int{a, b, c, d}[i])
				break
			}
		}
	}
}

// DefaultName is the console name reported by /info, /xinfo and /status when no
// name was given (the -name flag) and the resource file carries none. The C
// emulator hard codes "X32 Emulator" here; this port defaults to "REAPER",
// because it usually runs as the console of the REAPER bridge
// (see docs/X32_REAPER_BRIDGE.md).
const DefaultName = "REAPER"

// consoleName returns the console name ("/-prefs/name"), the name given at
// startup (State.Name) or DefaultName: the same fallback chain as /info, /xinfo
// and /status in X32.c, where a console without a name is called "X32 Emulator".
func (s *State) consoleName() string {
	for i := range Xprefs {
		if Xprefs[i].Addr == "/-prefs/name" {
			if n := Xprefs[i].Str(); n != "" {
				return n
			}
			break
		}
	}
	if s.Name != "" {
		return s.Name
	}
	return DefaultName
}

// SetName sets the name used while no client has set "/-prefs/name". It is what
// the -name flag uses at startup; the stored preference of a resource file (or a
// client renaming the console with "/-prefs/name ,s <name>") wins over it, as on
// a real console.
func (s *State) SetName(name string) { s.Name = name }

// cmd returns the i-th argument of the current request (zero values when absent
// or of a different type) - a safer equivalent of the fixed r_buf offsets used
// by X32.c.
func (s *State) argS(i int) string {
	if s.Msg == nil || i < 0 || i >= len(s.Msg.Args) {
		return ""
	}
	return s.Msg.Args[i].S
}

func (s *State) argI(i int) int32 {
	if s.Msg == nil || i < 0 || i >= len(s.Msg.Args) {
		return 0
	}
	return s.Msg.Args[i].I
}

// functionInfo handles /info (function_info in X32.c).
func (s *State) functionInfo() int {
	b := osc.Start("/info", "ssss")
	b = osc.AppendString(b, "V2.07")
	b = osc.AppendString(b, s.consoleName())
	b = osc.AppendString(b, "X32")
	b = osc.AppendString(b, XVersion)
	s.Out = b
	return SSnd
}

// functionXinfo handles /xinfo: this is what clients use to discover a console.
func (s *State) functionXinfo() int {
	b := osc.Start("/xinfo", "ssss")
	b = osc.AppendString(b, s.IP)
	b = osc.AppendString(b, s.consoleName())
	b = osc.AppendString(b, "X32")
	b = osc.AppendString(b, XVersion)
	s.Out = b
	return SSnd
}

// functionStatus handles /status.
func (s *State) functionStatus() int {
	b := osc.Start("/status", "sss")
	b = osc.AppendString(b, "active")
	b = osc.AppendString(b, s.IP)
	b = osc.AppendString(b, s.consoleName())
	s.Out = b
	return SSnd
}

// functionXremote handles /xremote: register or refresh the requesting client.
// It never answers.
func (s *State) functionXremote() int {
	s.RegisterClient(s.From)
	return 0
}

// functionUnsubscribe handles /unsubscribe.
func (s *State) functionUnsubscribe() int {
	s.UnregisterClient(s.From)
	return 0
}

// functionShowdump answers /showdump (and any address starting with "/sho"):
// the current show file as a "node ,s" message.
func (s *State) functionShowdump() int {
	name := ""
	for i := range Xshow {
		if Xshow[i].Addr == "/-show/showfile/show/name" {
			name = Xshow[i].Str()
			break
		}
	}
	b := osc.Start("node", "s")
	s.Out = osc.AppendString(b, "/-show/showfile/show \""+name+"\" 0 0 0 0 0 0 0 0 0 0 \""+XVersion+"\"")
	return SSnd
}

// functionShutdown handles the non standard /shutdown command: save the
// resource file and stop the process.
func (s *State) functionShutdown() int {
	s.X32Shutdown()
	return 0
}
