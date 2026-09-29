package x32

import (
	"strings"
)

// functionSlash is the port of function_slash(): the OSC address is exactly "/"
// and its first (and only) string argument holds "<path> [<value> ...]".
//
// The request is always echoed back to its sender when at least one parameter
// could be read; the individual parameters notify the other xremote clients as
// they are applied.
func (s *State) functionSlash() int {
	payload := ""
	if len(s.Msg.Args) > 0 && s.Msg.Args[0].Tag == 's' {
		payload = s.Msg.Args[0].S
	}
	echo := append([]byte(nil), s.Raw...)

	path := strings.TrimPrefix(payload, "/")
	// find the node table the path belongs to (first matching prefix)
	var tab Table
	for i := range nodeTable {
		if strings.HasPrefix(path, nodeTable[i].name) {
			tab = nodeTable[i].table
			break
		}
	}
	if tab == nil {
		return 0
	}
	// the path token stops at the first space
	tok := path
	if i := strings.IndexByte(tok, ' '); i >= 0 {
		tok = tok[:i]
	}
	for i := range tab {
		if len(tab[i].Addr) < 1 || tab[i].Addr[1:] != tok {
			continue
		}
		// data starts right after "<path> " (C: str_pt_in += strlen(cmd+1) + 1)
		rest := ""
		if len(path) >= len(tab[i].Addr) {
			rest = path[len(tab[i].Addr):]
		}
		s.cur = &cursor{s: rest}

		if tab[i].Flags == FFND {
			// walk to the group that owns the data
			for i+1 < len(tab) && tab[i+1].Flags == FFND {
				i++
			}
			a := newApplier(s, tab, i+1)
			s.slashGroup(a, tab[i].Type)
			if a.firstFailed() {
				return 0
			}
			s.Out = append(s.Out[:0], echo...)
			return SSnd
		}
		if tab[i].Flags&FSET != 0 {
			// single parameter, set from the textual value
			if s.cur.empty() {
				return 0
			}
			c := &tab[i]
			switch c.Type {
			case TI32, TP32:
				if t, ok := s.cur.word(); ok {
					s.setInt(c, parseIntPrefix(t))
				}
			case TE32:
				if t, ok := s.cur.word(); ok {
					s.setList(c, t)
				}
			case TF32:
				if t, ok := s.cur.word(); ok {
					s.setFloat(c, XrFloat(t))
				}
			case TS32:
				if t, ok := s.cur.str(); ok {
					s.setString(c, t)
				}
			default:
				// composed types cannot be set from a plain value
			}
			s.Out = append(s.Out[:0], echo...)
			return SSnd
		}
		return 0
	}
	return 0
}

// slashGroup applies the parameters of a "/" command to the group whose type is
// typ. It is the switch() of function_slash() in X32.c.
func (s *State) slashGroup(a *applier, typ Type) {
	switch typ {
	case TCHCO: // name, icon#, color, source
		a.str()
		a.int()
		a.list()
		a.int()
	case TCHDE: // delay on, time
		a.list()
		a.linf(0.3, 499.7, 0.1)
	case TCHPR: // trim, invert, hpon, hpslope, hpf
		a.linf(-18., 36., 0.25)
		a.list()
		a.list()
		a.list()
		a.logf(20., 2.9957322735, 100) // log(400/20)
	case TCHGA: // on, mode, thr, range, attack, hold, release, keysrc
		a.list()
		a.list()
		a.linf(-80., 80., 0.5)
		a.linf(3., 57., 1.)
		a.linf(0., 120., 1.)
		a.logf(0.02, 11.512925465, 100) // log(2000/0.02)
		a.logf(5., 6.684611728, 100)    // log(4000/5)
		a.int()
	case TCHGF, TCHDF: // filter on, type, frequency
		a.list()
		a.list()
		a.logf(20., 6.907755279, 200) // log(20000/20)
	case TCHDY: // on, mode, det, env, thr, ratio, knee, mgain, attack, hold, release, pos, keysrc, mix, auto
		a.list()
		a.list()
		a.list()
		a.list()
		a.linf(-60., 60., 0.5)
		a.list()
		a.linf(0., 5.0, 1.0)
		a.linf(0., 24.0, 0.5)
		a.linf(0., 120.0, 1.0)
		a.logf(0.02, 11.51292546, 100) // log(2000/0.02)
		a.logf(5., 6.684611728, 100)   // log(4000/5)
		a.list()
		a.list()
		a.linf(0., 100.0, 5.0)
		a.list()
	case TCHIN: // insert on, pos, sel
		a.list()
		a.list()
		a.list()
	case TCHEQ: // type, f, g, q
		a.list()
		a.logf(20., 6.907755279, 200) // log(20000/20)
		a.linf(-15., 30.0, 0.250)
		a.logf(10., -3.506557897, 71) // log(0.3/10)
	case TCHMX: // on, level, pan, panFollow, tap, tapFollow
		a.list()
		a.lvl(1023)
		a.list()
		a.linf(-100., 200., 2.)
		a.list()
		a.lvl(160)
	case TCHMO: // on, level, pan, type, panFollow
		a.list()
		a.lvl(1023)
		a.linf(-100., 200., 2.)
		a.list()
		a.list()
	case TCHME: // on, level
		a.list()
		a.lvl(1023)
	case TCHGRP: // dca, mute (bitmaps)
		a.per()
		a.per()
	case TCHAMIX: // group, weight
		a.list()
		a.linf(-12., 24., 0.5)
	case TAXPR: // trim, invert (note: X32.c writes the invert into the trim entry too)
		a.linf(-18., 36., 0.25)
		a.list() // the C code targets command[i+1] again for the invert
	case TBSCO: // name, icon#, color
		a.str()
		a.int()
		a.list()
	case TMXPR: // invert
		a.list()
	case TMXDY: // on, mode, det, env, thr, ratio, knee, mgain, attack, hold, release, pos, mix, auto
		a.list()
		a.list()
		a.list()
		a.list()
		a.linf(-60., 60., 0.5)
		a.list()
		a.linf(0., 5.0, 1.0)
		a.linf(0., 24.0, 0.5)
		a.linf(0., 120.0, 1.0)
		a.logf(0.02, 11.51292546, 100) // log(2000/0.02)
		a.logf(5., 6.684611728, 100)   // log(4000/5)
		a.list()
		a.linf(0., 100.0, 5.0)
		a.list()
	case TMSMX: // on, level, pan
		a.list()
		a.lvl(1023)
		a.linf(-100., 200., 2.)
	default:
		// other types are applied below
		s.slashGroup2(a, typ)
	}
}
