package bridge

import (
	"strconv"
	"strings"

	"x32emu/osc"
)

// ---------------------------------------------------------------- REAPER side

// reaperAddr builds the address of a REAPER parameter: track parameters are
// absolute ("/track/<n>/<leaf>"), master and action addresses are passed through.
func reaperAddr(track int, leaf string) string {
	switch {
	case strings.HasPrefix(leaf, "master/"), strings.HasPrefix(leaf, "action"):
		return "/" + leaf
	default:
		return "/track/" + strconv.Itoa(track) + "/" + leaf
	}
}

// reaperFloat builds a ",f <v>" message for a REAPER parameter.
func (b *Bridge) reaperFloat(track int, leaf string, v float32) []byte {
	return osc.Build(reaperAddr(track, leaf), "f", osc.Arg{Tag: 'f', F: v})
}

// reaperString builds a ",s <v>" message for a REAPER parameter.
func (b *Bridge) reaperString(track int, leaf, v string) []byte {
	return osc.Build(reaperAddr(track, leaf), "s", osc.Arg{Tag: 's', S: v})
}

// actionMsg builds a bare "/action/<id>" address: X32ReaperW sends the action
// numbers without any argument block and every REAPER version accepts it.
func actionMsg(id int) []byte {
	return osc.AppendString(nil, "/action/"+strconv.Itoa(id))
}

// soloMsg builds the "/-stat/solosw/NN ,i v" message of a console strip.
func soloMsg(strip int, v int32) []byte {
	return osc.Build("/-stat/solosw/"+num(strip), "i", osc.Arg{Tag: 'i', I: v})
}

// ---------------------------------------------------------------- console side

// x32Addr builds the address of a console parameter. section is one of ch,
// auxin, fxrtn, bus, mtx, dca or main/st.
func x32Addr(section string, strip int, leaf string) string {
	switch section {
	case "dca":
		return "/dca/" + strconv.Itoa(strip) + "/" + leaf
	case "main/st":
		return "/main/st/" + leaf
	default:
		return "/" + section + "/" + num(strip) + "/" + leaf
	}
}

// x32Int builds a ",i <v>" message for a console parameter.
func x32Int(section string, strip int, leaf string, v int32) []byte {
	return osc.Build(x32Addr(section, strip, leaf), "i", osc.Arg{Tag: 'i', I: v})
}

// x32Float builds a ",f <v>" message for a console parameter.
func x32Float(section string, strip int, leaf string, v float32) []byte {
	return osc.Build(x32Addr(section, strip, leaf), "f", osc.Arg{Tag: 'f', F: v})
}

// x32String builds a ",s <v>" message for a console parameter.
func x32String(section string, strip int, leaf, v string) []byte {
	return osc.Build(x32Addr(section, strip, leaf), "s", osc.Arg{Tag: 's', S: v})
}

// sectionNumber converts an absolutely numbered console strip (1..80) into the
// number used inside its section: channels keep 1..32, aux inputs and effect
// returns 1..8 and buses 1..16.
func sectionNumber(section string, strip int) int {
	switch section {
	case "auxin":
		return strip - StripAuxinFirst + 1
	case "fxrtn":
		return strip - StripFxrtnFirst + 1
	case "bus":
		return strip - StripBusFirst + 1
	}
	return strip
}

// sectionOf returns the console section and the leaf prefix matching a REAPER
// track: channel strips of the current bank are "/ch", the other sections keep
// their own numbering.
func (b *Bridge) sectionOf(track int) (string, int, bool) {
	if strip, ok := b.cfg.TrackToX32(track); ok {
		switch {
		case strip <= b.cfg.BankSize:
			return "ch", strip, true
		case strip <= StripAuxinLast:
			return "auxin", strip, true
		case strip <= StripFxrtnLast:
			return "fxrtn", strip, true
		default:
			return "bus", strip, true
		}
	}
	// a track that is not in the current channel bank is still reachable through
	// the auxin/fxrtn/bus ranges when those are configured
	return "", 0, false
}
