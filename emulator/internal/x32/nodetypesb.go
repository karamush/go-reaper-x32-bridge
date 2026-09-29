package x32

// renderGroupB renders the /config, /-prefs, /-stat, /-show and library groups.
func (s *State) renderGroupB(tab Table, i int) (string, bool) {
	var b []byte
	add := func(ss ...string) {
		for _, x := range ss {
			b = append(b, x...)
		}
	}
	switch tab[i].Type {
	case TOMAIN:
		add(Sint(gi(tab, i, 1)), ge("Smpos", gi(tab, i, 2)), onOff(gi(tab, i, 3) != 0))
	case TOMAIN2:
		add(Sint(gi(tab, i, 1)), ge("Smpos", gi(tab, i, 2)))
	case TOP16:
		for j := 1; j < 5; j++ {
			add(Sint(gi(tab, i, j)))
		}
	case TOMAIND:
		add(onOff(gi(tab, i, 1) != 0), Slinf(gf(tab, i, 2), 0.3, 500., 1))
	case THAMP:
		add(Slinf(gf(tab, i, 1), -12., 60., 1), onOff(gi(tab, i, 2) != 0))
	case TPREFS:
		add(quoted(gs(tab, i, 1), " \"\""),
			Slinf(gf(tab, i, 2), 10., 100., 0), Slinf(gf(tab, i, 3), 0., 100., 0),
			Slinf(gf(tab, i, 4), 10., 100., 0), Slinf(gf(tab, i, 5), 10., 100., 0),
			onOff(gi(tab, i, 6) != 0))
		if gi(tab, i, 7) != 0 {
			add(" 44k1")
		} else {
			add(" 48k")
		}
		add(ge("Psource", gi(tab, i, 8)))
		for j := 9; j < 16; j++ {
			add(onOff(gi(tab, i, j) != 0))
		}
		add(Sbitmp(gi(tab, i, 16), 4), onOff(gi(tab, i, 17) != 0), ge("PSCont", gi(tab, i, 18)))
		if gi(tab, i, 19) != 0 {
			add(" 12h")
		} else {
			add(" 24h")
		}
		add(onOff(gi(tab, i, 20) != 0), onOff(gi(tab, i, 21) != 0))
		if gi(tab, i, 22) != 0 {
			add(" INV")
		} else {
			add(" NORM")
		}
		add(quoted(gs(tab, i, 23), " \"\""))
	case TPIR:
		add(onOff(gi(tab, i, 1) != 0), ge("PRpro", gi(tab, i, 2)), ge("PRport", gi(tab, i, 3)), Sbitmp(gi(tab, i, 4), 12))
	case TPIQ:
		add(ge("XiQspk", gi(tab, i, 1)), ge("XiQeq", gi(tab, i, 2)), Sint(gi(tab, i, 3)))
	case TPCARD:
		add(ge("Pctype", gi(tab, i, 1)), ge("Pufmode", gi(tab, i, 2)), ge("Pusbmod", gi(tab, i, 3)))
		if gi(tab, i, 4) != 0 {
			add(" OUT")
		} else {
			add(" IN")
		}
		add(ge("Pcas", gi(tab, i, 5)))
		if gi(tab, i, 6) != 0 {
			add(" 64")
		} else {
			add(" 56")
		}
		add(ge("Pcmadi", gi(tab, i, 7)), ge("Pcmado", gi(tab, i, 8)), ge("Pmadsrc", gi(tab, i, 9)))
	case TPRTA:
		add(ge("Prtavis", gi(tab, i, 1)), Slinf(gf(tab, i, 2), 0., 60., 0), onOff(gi(tab, i, 3) != 0),
			Sint(gi(tab, i, 4)))
		if gi(tab, i, 5) != 0 {
			add("POST")
		} else {
			add(" PRE")
		}
		if gi(tab, i, 6) != 0 {
			add(" SPEC")
		} else {
			add(" BAR")
		}
		add(Sbitmp(gi(tab, i, 7), 6))
		if gi(tab, i, 8) != 0 {
			add(" PEAK")
		} else {
			add(" RMS")
		}
		add(Slogf(gf(tab, i, 9), 0.25, 16., 2), ge("Prtaph", gi(tab, i, 10)))
	case TPIP:
		add(onOff(gi(tab, i, 1) != 0))
	case TPKEY:
		add(Sint(gi(tab, i, 1)), quoted(gs(tab, i, 2), " \" \""))
	case TPADDR, TPMASK, TPGWAY:
		for j := 1; j < 5; j++ {
			add(Sint(gi(tab, i, j)))
		}
	case TSTAT:
		add(ge("Sselidx", gi(tab, i, 1)), Sint(gi(tab, i, 2)), Sint(gi(tab, i, 3)), onOff(gi(tab, i, 4) != 0),
			Sint(gi(tab, i, 5)), Sint(gi(tab, i, 6)), onOff(gi(tab, i, 7) != 0), onOff(gi(tab, i, 8) != 0),
			Sint(gi(tab, i, 9)), onOff(gi(tab, i, 10) != 0), onOff(gi(tab, i, 11) != 0), onOff(gi(tab, i, 12) != 0),
			onOff(gi(tab, i, 13) != 0))
		if gi(tab, i, 14) != 0 {
			add(" SPEC")
		} else {
			add(" BAR")
		}
		if gi(tab, i, 15) != 0 {
			add(" SPEC")
		} else {
			add(" BAR")
		}
		add(onOff(gi(tab, i, 16) != 0), onOff(gi(tab, i, 17) != 0), Sint(gi(tab, i, 18)), Sint(gi(tab, i, 19)),
			onOff(gi(tab, i, 20) != 0), onOff(gi(tab, i, 21) != 0), Sint(gi(tab, i, 22)), Sint(gi(tab, i, 23)))
	case TSSCREEN:
		add(ge("Sscrn", gi(tab, i, 1)), onOff(gi(tab, i, 2) != 0), onOff(gi(tab, i, 3) != 0))
	case TSCHA:
		add(ge("Schal", gi(tab, i, 1)))
	case TSMET:
		add(ge("Smetl", gi(tab, i, 1)))
	case TSROU:
		add(ge("Sroul", gi(tab, i, 1)))
	case TSSET:
		add(ge("Ssetl", gi(tab, i, 1)))
	case TSLIB:
		add(ge("Slibl", gi(tab, i, 1)))
	case TSFX:
		add(ge("Sfxl", gi(tab, i, 1)))
	case TSMON:
		add(ge("Smonl", gi(tab, i, 1)))
	case TSUSB:
		add(ge("Susbl", gi(tab, i, 1)))
	case TSSCE:
		add(ge("Sscel", gi(tab, i, 1)))
	case TSASS:
		add(ge("Sassl", gi(tab, i, 1)))
	default:
		return "", false
	}
	return string(b), true
}
