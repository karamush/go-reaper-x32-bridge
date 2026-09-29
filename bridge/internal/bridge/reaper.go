package bridge

import (
	"encoding/binary"
	"math"
	"strconv"

	"x32emu/osc"
)

// splitBundle splits an OSC bundle into its messages. REAPER answers status
// requests with "#bundle" datagrams; anything that is not a bundle (or a
// truncated one) is returned as is.
func splitBundle(b []byte) [][]byte {
	if len(b) < 16 || string(b[:7]) != "#bundle" || b[7] != 0 {
		return [][]byte{b}
	}
	out := [][]byte{}
	p := 16 // skip "#bundle\0" and the 8 byte time tag
	for p+4 <= len(b) {
		n := int(binary.BigEndian.Uint32(b[p:]))
		p += 4
		if n <= 0 || p+n > len(b) {
			break
		}
		out = append(out, b[p:p+n])
		p += n
	}
	return out
}

// argNum returns the single numeric argument of a message whatever its declared
// type: REAPER sends floats for most parameters, but binaries and integers are
// legal too ("TRACK_MUTE b/track/@/mute").
func argNum(m *osc.Message) (float32, bool) {
	if len(m.Args) == 0 {
		return 0, false
	}
	a := m.Args[0]
	switch m.Tags {
	case "f":
		return a.F, true
	case "i":
		return float32(a.I), true
	case "b":
		if len(a.B) >= 4 {
			return float32(math.Float32frombits(binary.BigEndian.Uint32(a.B))), true
		}
	}
	return 0, false
}

// reaperToX32 turns one REAPER message into the console messages it implies.
func (b *Bridge) reaperToX32(m *osc.Message) [][]byte {
	f := fields(m.Addr)
	if len(f) == 0 {
		return nil
	}
	switch f[0] {
	case "track":
		return b.reaperTrack(f, m)
	case "master":
		return b.reaperMaster(f, m)
	case "play", "pause", "stop", "record", "repeat":
		return b.reaperTransport(f[0])
	case "device", "reaper":
		return nil // device state announcements: nothing to do on the desk
	}
	return nil
}

// reaperTrack handles "/track/<n>/..." (volume, pan, mute, solo, select, name and
// send levels).
//
// The state of every track REAPER reports is cached, even when the track is not
// in the channel bank the desk is showing: this is what makes a bank change
// instant, because the console can then be restored from the cache instead of
// waiting for REAPER to report again.
func (b *Bridge) reaperTrack(f []string, m *osc.Message) [][]byte {
	if len(f) < 3 {
		return nil
	}
	track := atoi(f[1])
	st := b.state.get(track)
	section, strip, visible := b.sectionOf(track)
	switch {
	case len(f) == 3 && f[2] == "volume":
		v, ok := argNum(m)
		if !ok || b.cfg.ToX32&BitFader == 0 {
			return nil
		}
		st.fader, st.valid = v, true
		if !visible {
			return nil
		}
		return [][]byte{x32Float(section, strip, "mix/fader", v)}
	case len(f) == 3 && f[2] == "pan":
		v, ok := argNum(m)
		if !ok || b.cfg.ToX32&BitPan == 0 {
			return nil
		}
		st.pan, st.valid = v, true
		if !visible {
			return nil
		}
		return [][]byte{x32Float(section, strip, "mix/pan", v)}
	case len(f) == 3 && f[2] == "mute":
		v, ok := argNum(m)
		if !ok || b.cfg.ToX32&BitMute == 0 {
			return nil
		}
		st.mute, st.valid = v > 0.5, true
		if !visible {
			return nil
		}
		// REAPER muted == desk channel off
		on := int32(1)
		if v > 0.5 {
			on = 0
		}
		return [][]byte{x32Int(section, strip, "mix/on", on)}
	case len(f) == 3 && f[2] == "solo":
		v, ok := argNum(m)
		if !ok || b.cfg.ToX32&BitSolo == 0 {
			return nil
		}
		st.solo, st.valid = v > 0.5, true
		if !visible {
			return nil
		}
		solo := int32(0)
		if v > 0.5 {
			solo = 1
		}
		return [][]byte{soloMsg(strip, solo)}
	case len(f) == 3 && f[2] == "select":
		if b.cfg.ToX32&BitSelect == 0 || !visible {
			return nil
		}
		// the desk walks its selection with a zero based index
		return [][]byte{osc.Build("/-stat/selidx", "i", osc.Arg{Tag: 'i', I: int32(strip - 1)})}
	case len(f) == 3 && f[2] == "name":
		v, ok := argS(m, 0)
		if !ok || b.cfg.ToX32&BitName == 0 {
			return nil
		}
		st.name, st.valid = v, true
		if !visible {
			return nil
		}
		return [][]byte{x32String(section, strip, "config/name", v)}
	case len(f) == 5 && f[2] == "send" && f[4] == "volume":
		v, ok := argNum(m)
		if !ok || b.cfg.ToX32&BitSend == 0 {
			return nil
		}
		bus := atoi(f[3]) - b.cfg.SendOffset
		if bus < 1 || bus > MaxSends {
			return nil
		}
		st.sends[bus-1], st.valid = v, true
		if !visible {
			return nil
		}
		return [][]byte{x32Float(section, strip, "mix/"+num(bus)+"/level", v)}
	case len(f) == 3 && f[2] == "vu":
		// REAPER's track meter; the console meter frames are built from it by the
		// emulator (see MeterSink)
		if v, ok := argNum(m); ok {
			b.feedTrackMeter(track, v)
		}
		return nil
	case len(f) == 4 && f[2] == "vu" && (f[3] == "L" || f[3] == "R"):
		v, ok := argNum(m)
		if !ok {
			return nil
		}
		// a channel meter of the desk is one bar: use the loudest side
		if f[3] == "L" {
			st.vuL = v
		} else {
			st.vuR = v
		}
		b.feedTrackMeter(track, maxf(st.vuL, st.vuR))
		return nil
	}
	return nil
}

// feedTrackMeter remembers the meter level of one REAPER track and forwards it to
// the console through the meter sink, resolved to the strip the track is shown
// on. A track that is not mapped (or not in the current channel bank) is cached
// but not fed.
func (b *Bridge) feedTrackMeter(track int, v float32) {
	st := b.state.get(track)
	st.vu = v
	if b.meters.Track == nil {
		return
	}
	section, strip, ok := b.sectionOf(track)
	if !ok {
		return
	}
	b.meters.Track(section, sectionNumber(section, strip), v)
}

// feedBankMeters pushes the cached meter levels of the current channel bank, so
// that switching banks does not leave the bars of the new bank at the level of
// the tracks that were shown before.
func (b *Bridge) feedBankMeters() {
	if b.meters.Track == nil {
		return
	}
	for strip := 1; strip <= b.cfg.BankSize; strip++ {
		track, ok := b.cfg.X32ToTrack(strip)
		if !ok {
			continue
		}
		if st := b.state.peek(track); st != nil {
			b.meters.Track("ch", strip, st.vu)
		}
	}
}

// maxf returns the larger of two levels (used for the L/R meters of a track,
// which the desk shows as a single bar).
func maxf(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

// reaperMaster handles "/master/{volume,pan,vu,vu/L,vu/R}".
func (b *Bridge) reaperMaster(f []string, m *osc.Message) [][]byte {
	if len(f) < 2 {
		return nil
	}
	// the meters of the main bus reach the console without going through the
	// parameter masks: they are not a parameter the desk can be told to ignore
	if f[1] == "vu" {
		v, ok := argNum(m)
		if ok {
			switch {
			case len(f) == 2: // /master/vu
				b.masterMeter[0], b.masterMeter[1] = v, v
			case len(f) == 3 && f[2] == "L": // /master/vu/L
				b.masterMeter[0] = v
			case len(f) == 3 && f[2] == "R": // /master/vu/R
				b.masterMeter[1] = v
			default:
				return nil
			}
			if b.meters.Master != nil {
				b.meters.Master(b.masterMeter[0], b.masterMeter[1])
			}
		}
		return nil
	}
	if !b.cfg.Master || len(f) != 2 {
		return nil
	}
	switch f[1] {
	case "volume":
		v, ok := argNum(m)
		if !ok || b.cfg.ToX32&BitMFader == 0 {
			return nil
		}
		return [][]byte{x32Float("main/st", 0, "mix/fader", v)}
	case "pan":
		v, ok := argNum(m)
		if !ok || b.cfg.ToX32&BitMPan == 0 {
			return nil
		}
		return [][]byte{x32Float("main/st", 0, "mix/pan", v)}
	}
	return nil
}

// reaperTransport mirrors REAPER's transport state on the Bank C buttons of the
// desk. A console button is "127 while it is held down, 0 when released", so the
// active state is reported as 127 and stopping releases the play and pause
// buttons as well.
func (b *Bridge) reaperTransport(what string) [][]byte {
	if !b.cfg.Transport {
		return nil
	}
	value := func(idx int, v int32) []byte {
		return osc.Build("/-stat/userpar/"+strconv.Itoa(idx)+"/value", "i", osc.Arg{Tag: 'i', I: v})
	}
	switch what {
	case "play":
		return [][]byte{value(18, 127)}
	case "pause":
		return [][]byte{value(19, 127)}
	case "repeat":
		return [][]byte{value(22, 127)}
	case "record":
		return [][]byte{value(24, 127)}
	case "stop":
		return [][]byte{value(23, 0), value(18, 0), value(19, 0)}
	}
	return nil
}
