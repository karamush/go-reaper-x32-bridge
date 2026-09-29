package x32

// renderGroupC renders the monitor, show file, recorder and library groups.
func (s *State) renderGroupC(tab Table, i int) (string, bool) {
	var b []byte
	add := func(ss ...string) {
		for _, x := range ss {
			b = append(b, x...)
		}
	}
	switch tab[i].Type {
	case TSOSC:
		add(onOff(gi(tab, i, 1) != 0))
	case TSTALK:
		add(onOff(gi(tab, i, 1) != 0), onOff(gi(tab, i, 2) != 0))
	case TUSB:
		add(quoted(gs(tab, i, 1), " \"\""), quoted(gs(tab, i, 2), " \"\""))
	case TSNAM: // show file header: name + 10 counters + version
		add(quoted(gs(tab, i, 1), " \"\""))
		for j := 2; j < 12; j++ {
			add(Sint(gi(tab, i, j)))
		}
		add(" \"" + XVersion + "\"")
	case TSCUE: // cue point
		add(Sint(gi(tab, i, 1)), quoted(gs(tab, i, 2), " \"\""))
		for j := 3; j < 10; j++ {
			add(Sint(gi(tab, i, j)))
		}
	case TSSCN: // scene name, notes, safes, hasdata
		add(quoted(gs(tab, i, 1), " \"\""), quoted(gs(tab, i, 2), " \"\""), Sbitmp(gi(tab, i, 3), 9), Sint(gi(tab, i, 4)))
	case TSSNP: // snippet
		add(quoted(gs(tab, i, 1), " \"\""))
		for j := 2; j < 7; j++ {
			add(Sint(gi(tab, i, j)))
		}
	case TSAES:
		add(quoted(gs(tab, i, 1), " \" \""), quoted(gs(tab, i, 2), " \" \""), Sint(gi(tab, i, 3)))
	case TSTAPE:
		add(ge("Stapl", gi(tab, i, 1)), quoted(gs(tab, i, 2), " \"\""), Sint(gi(tab, i, 3)), Sint(gi(tab, i, 4)))
	case THA, TACTION:
		// nothing is rendered by X32.c
	case TUREC:
		for j := 1; j < 6; j++ {
			add(Sint(gi(tab, i, j)))
		}
		add(ge("Ubat", gi(tab, i, 6)))
		for j := 7; j < 11; j++ {
			add(Sint(gi(tab, i, j)))
		}
		add(ge("Usdc", gi(tab, i, 11)), ge("Usdc", gi(tab, i, 12)),
			quoted(gs(tab, i, 13), " \"\""), quoted(gs(tab, i, 14), " \"\""), quoted(gs(tab, i, 15), " \"\""),
			Sint(gi(tab, i, 16)))
	case TSLIBS: // library entry: number, name, type, flags, hasdata
		add(Sint(gi(tab, i, 1)), quoted(gs(tab, i, 2), " \"\""), Sint(gi(tab, i, 3)),
			Sbitmp(gi(tab, i, 4), 16), Sint(gi(tab, i, 5)))
	case TD48:
		add(Sbitmp(gi(tab, i, 1), 4), Sint(gi(tab, i, 2)))
	case TD48A:
		for j := 1; j < 49; j++ {
			add(Sint(gi(tab, i, j)))
		}
	case TD48G:
		for j := 1; j < 13; j++ {
			add(quoted(gs(tab, i, j), " \"\""))
		}
	case TUROUO:
		for j := 1; j < 49; j++ {
			add(Sint(gi(tab, i, j)))
		}
	case TUROUI:
		for j := 1; j < 33; j++ {
			add(Sint(gi(tab, i, j)))
		}
	default:
		return "", false
	}
	return string(b), true
}
