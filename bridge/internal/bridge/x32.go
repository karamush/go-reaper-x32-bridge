package bridge

import (
	"strconv"
	"strings"

	"x32emu/osc"
)

// ---------------------------------------------------------------- address helpers

// fields splits an OSC address into its non empty parts: "/ch/03/mix/fader"
// becomes ["ch", "03", "mix", "fader"]. Parsing the address instead of using the
// fixed byte offsets of X32ReaperW makes one or two digit numbers work equally
// well (the C code only copes with two digit indexes).
func fields(addr string) []string {
	raw := strings.Split(strings.Trim(addr, "/"), "/")
	out := raw[:0]
	for _, f := range raw {
		if f != "" {
			out = append(out, f)
		}
	}
	return out
}

// atoi is strconv.Atoi with a "0 when not a number" policy.
func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

// num formats an index on two digits, as every X32 address does.
func num(i int) string {
	if i < 10 {
		return "0" + strconv.Itoa(i)
	}
	return strconv.Itoa(i)
}

// Argument accessors: they return ok == false when the packet does not carry the
// argument the tag string promised (a truncated datagram stops the decoder).
func argF(m *osc.Message, i int) (float32, bool) {
	if m.Tags != "f" || i >= len(m.Args) {
		return 0, false
	}
	return m.Args[i].F, true
}

func argI(m *osc.Message, i int) (int32, bool) {
	if m.Tags != "i" || i >= len(m.Args) {
		return 0, false
	}
	return m.Args[i].I, true
}

func argS(m *osc.Message, i int) (string, bool) {
	if m.Tags != "s" || i >= len(m.Args) {
		return "", false
	}
	return m.Args[i].S, true
}

// ---------------------------------------------------------------- console -> REAPER

// x32ToReaper turns one message of the console into the REAPER messages it
// implies. It returns nil when the message is of no interest to the bridge (or
// filtered out by the masks).
func (b *Bridge) x32ToReaper(m *osc.Message) [][]byte {
	f := fields(m.Addr)
	if len(f) < 2 {
		return nil
	}
	switch f[0] {
	case "ch", "auxin", "fxrtn", "bus", "mtx":
		return b.stripToReaper(f, m)
	case "dca":
		return b.dcaToReaper(f, m)
	case "main":
		return b.mainToReaper(f, m)
	case "-stat":
		return b.statToReaper(f, m)
	}
	return nil
}

// stripToReaper handles "/<section>/<strip>/..." for the channel, auxin, fxrtn,
// bus and matrix sections.
func (b *Bridge) stripToReaper(f []string, m *osc.Message) [][]byte {
	strip, err := strconv.Atoi(f[1])
	if err != nil {
		return nil
	}
	track, ok := b.cfg.X32ToTrack(strip)
	if !ok {
		return nil
	}
	st := b.state.get(track)
	switch {
	case len(f) == 4 && f[2] == "mix" && f[3] == "fader":
		v, ok := argF(m, 0)
		if !ok || b.cfg.ToReaper&BitFader == 0 {
			return nil
		}
		st.fader, st.valid = v, true
		return [][]byte{b.reaperFloat(track, "volume", v)}
	case len(f) == 4 && f[2] == "mix" && f[3] == "pan":
		v, ok := argF(m, 0)
		if !ok || b.cfg.ToReaper&BitPan == 0 {
			return nil
		}
		st.pan, st.valid = v, true
		return [][]byte{b.reaperFloat(track, "pan", v)}
	case len(f) == 4 && f[2] == "mix" && f[3] == "on":
		v, ok := argI(m, 0)
		if !ok || b.cfg.ToReaper&BitMute == 0 {
			return nil
		}
		// the desk sends 1 for "channel on"; REAPER's mute is the opposite
		mute := float32(1)
		if v == 1 {
			mute = 0
		}
		st.mute, st.valid = mute > 0.5, true
		return [][]byte{b.reaperFloat(track, "mute", mute)}
	case len(f) == 4 && f[2] == "config" && f[3] == "name":
		v, ok := argS(m, 0)
		if !ok || b.cfg.ToReaper&BitName == 0 {
			return nil
		}
		st.name, st.valid = v, true
		return [][]byte{b.reaperString(track, "name", v)}
	case len(f) == 5 && f[2] == "mix" && f[4] == "level":
		v, ok := argF(m, 0)
		if !ok || b.cfg.ToReaper&BitSend == 0 {
			return nil
		}
		bus, err := strconv.Atoi(f[3])
		if err != nil || bus < 1 || bus > MaxSends {
			return nil
		}
		// REAPER increases the logical send numbers when a track also feeds a
		// hardware output (X32ReaperW "TrackSendOffset")
		send := bus + b.cfg.SendOffset
		st.sends[bus-1], st.valid = v, true
		return [][]byte{b.reaperFloat(track, "send/"+strconv.Itoa(send)+"/volume", v)}
	}
	return nil
}

// dcaToReaper handles "/dca/<n>/{fader,on,config/name}": one DCA drives the
// whole REAPER track range declared for it (RdcaMin/RdcaMax).
func (b *Bridge) dcaToReaper(f []string, m *osc.Message) [][]byte {
	n := atoi(f[1])
	if n < 1 || n > 8 {
		return nil
	}
	tracks := b.cfg.DcaTracks(n)
	if len(tracks) == 0 {
		return nil
	}
	var out [][]byte
	switch {
	case len(f) == 3 && f[2] == "fader":
		v, ok := argF(m, 0)
		if !ok || b.cfg.ToReaper&BitFader == 0 {
			return nil
		}
		for _, t := range tracks {
			out = append(out, b.reaperFloat(t, "volume", v))
		}
	case len(f) == 3 && f[2] == "on":
		v, ok := argI(m, 0)
		if !ok || b.cfg.ToReaper&BitMute == 0 {
			return nil
		}
		mute := float32(1)
		if v == 1 {
			mute = 0
		}
		for _, t := range tracks {
			out = append(out, b.reaperFloat(t, "mute", mute))
		}
	case len(f) == 4 && f[2] == "config" && f[3] == "name":
		v, ok := argS(m, 0)
		if !ok || b.cfg.ToReaper&BitName == 0 {
			return nil
		}
		for _, t := range tracks {
			out = append(out, b.reaperString(t, "name", v))
		}
	}
	return out
}

// mainToReaper mirrors the stereo master bus onto REAPER's master track.
func (b *Bridge) mainToReaper(f []string, m *osc.Message) [][]byte {
	if !b.cfg.Master || len(f) != 4 || f[2] != "mix" {
		return nil
	}
	switch f[3] {
	case "fader":
		v, ok := argF(m, 0)
		if !ok || b.cfg.ToReaper&BitMFader == 0 {
			return nil
		}
		return [][]byte{b.reaperFloat(0, "master/volume", v)}
	case "pan":
		v, ok := argF(m, 0)
		if !ok || b.cfg.ToReaper&BitMPan == 0 {
			return nil
		}
		return [][]byte{b.reaperFloat(0, "master/pan", v)}
	}
	return nil
}

// statToReaper handles the console status notifications: the selected channel,
// the solo switches and the user assign controls (Bank C is the transport
// remote, see §7.7 of docs/X32_OSC_PROTOCOL.md).
func (b *Bridge) statToReaper(f []string, m *osc.Message) [][]byte {
	if len(f) < 2 {
		return nil
	}
	switch f[1] {
	case "selidx":
		v, ok := argI(m, 0)
		if !ok || b.cfg.ToReaper&BitSelect == 0 {
			return nil
		}
		// the desk reports zero based strips
		track, ok := b.cfg.X32ToTrack(int(v) + 1)
		if !ok {
			return nil
		}
		return [][]byte{actionMsg(40297), b.reaperFloat(track, "select", 1)}
	case "solosw":
		if len(f) != 3 {
			return nil
		}
		v, ok := argI(m, 0)
		if !ok || b.cfg.ToReaper&BitSolo == 0 {
			return nil
		}
		solo := float32(0)
		if v == 1 {
			solo = 1
		}
		strip := atoi(f[2])
		if strip >= StripDcaFirst && strip <= StripDcaLast {
			var out [][]byte
			for _, t := range b.cfg.DcaTracks(strip - StripDcaFirst + 1) {
				out = append(out, b.reaperFloat(t, "solo", solo))
			}
			return out
		}
		track, ok := b.cfg.X32ToTrack(strip)
		if !ok {
			return nil
		}
		return [][]byte{b.reaperFloat(track, "solo", solo)}
	case "userpar":
		// the Bank C buttons are only useful when one of the two modes is on
		if len(f) != 4 || f[3] != "value" || (!b.cfg.Transport && !b.cfg.ChBank) {
			return nil
		}
		v, ok := argI(m, 0)
		if !ok {
			return nil
		}
		return b.userParToReaper(atoi(f[2]), int(v))
	}
	return nil
}

// userParToReaper converts a user assign event into REAPER actions. The mapping
// is the Bank C layout of X32ReaperW (§3.2): buttons 17…24 are the transport and
// the four encoders 33…36 scrub by beat, measure, marker and item. Buttons are
// acted upon when they come back **up** (value 0), encoders on every turn
// (value < 64 = backwards, >= 64 = forwards), after which they are re-centered.
func (b *Bridge) userParToReaper(idx, v int) [][]byte {
	switch {
	case idx >= 17 && idx <= 24:
		if v != 0 { // "ignore non 0 value": the action fires on the button release
			return nil
		}
		button := idx - 12
		if b.cfg.ChBank && button == b.cfg.BankUp {
			return b.bankStep(+1)
		}
		if b.cfg.ChBank && button == b.cfg.BankDn {
			return b.bankStep(-1)
		}
		if !b.cfg.Transport {
			if button == b.cfg.MarkerButton {
				return [][]byte{actionMsg(40157)} // insert marker at the cursor
			}
			return nil
		}
		switch button {
		case 5:
			return [][]byte{actionMsg(40042)} // home
		case 6:
			return [][]byte{osc.Build("/play", "f", osc.Arg{Tag: 'f', F: 1})}
		case 7:
			return [][]byte{osc.Build("/pause", "f", osc.Arg{Tag: 'f', F: 1})}
		case 8:
			return [][]byte{actionMsg(40043)} // end of project
		case 9:
			return b.toggleLoop()
		case 10:
			return [][]byte{osc.Build("/repeat", "f", osc.Arg{Tag: 'f', F: 1})}
		case 11:
			return [][]byte{osc.Build("/stop", "f", osc.Arg{Tag: 'f', F: 1})}
		case 12:
			return [][]byte{osc.Build("/record", "f", osc.Arg{Tag: 'f', F: 1})}
		}
	case idx >= 33 && idx <= 36:
		back := v < 64
		var id int
		switch idx {
		case 33:
			id = pick(back, 40842, 40841) // previous / next beat
		case 34:
			id = pick(back, 40840, 40839) // previous / next measure
		case 35:
			id = pick(back, 40172, 40173) // previous / next marker
		case 36:
			id = pick(back, 40318, 40319) // item left / right
		}
		// an encoder is endless: re-center it so that the next turn reports a
		// direction again (X32ReaperW writes userpar = 64 after every move). This
		// message goes to the desk, not to REAPER, so it is queued directly.
		b.toX32(osc.AppendInt32(osc.Start("/-stat/userpar/"+strconv.Itoa(idx)+"/value", "i"), 64), 0)
		return [][]byte{actionMsg(id)}
	}
	return nil
}

// toggleLoop implements the loop button of Bank C (X32ReaperW keeps a toggle
// state and sends the two REAPER loop actions in turn).
func (b *Bridge) toggleLoop() [][]byte {
	if b.loopOn {
		b.loopOn = false
		return [][]byte{actionMsg(40223)} // end loop
	}
	b.loopOn = true
	return [][]byte{actionMsg(40222)} // start loop
}

func pick(cond bool, yes, no int) int {
	if cond {
		return yes
	}
	return no
}
