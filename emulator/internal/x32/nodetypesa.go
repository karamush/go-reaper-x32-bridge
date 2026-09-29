package x32

// Small accessors used by the /node renderers (never panic on short tables).
func gi(tab Table, i, off int) int32 {
	if i+off >= 0 && i+off < len(tab) {
		return tab[i+off].ValI
	}
	return 0
}

func gf(tab Table, i, off int) float32 {
	if i+off >= 0 && i+off < len(tab) {
		return tab[i+off].ValF
	}
	return 0
}

func gs(tab Table, i, off int) string {
	if i+off >= 0 && i+off < len(tab) {
		return tab[i+off].Str()
	}
	return ""
}

func ge(name string, idx int32) string { return enumRaw(name, idx) }

// renderGroup builds the value text of a group (the switch() at the end of
// function_node() in X32.c).
func (s *State) renderGroup(tab Table, i int) string {
	if v, ok := s.renderGroupA(tab, i); ok {
		return v
	}
	if v, ok := s.renderGroupB(tab, i); ok {
		return v
	}
	if v, ok := s.renderGroupC(tab, i); ok {
		return v
	}
	return ""
}

// renderGroupA renders the channel strip, bus, DCA and FX groups.
func (s *State) renderGroupA(tab Table, i int) (string, bool) {
	var b []byte
	add := func(ss ...string) {
		for _, x := range ss {
			b = append(b, x...)
		}
	}
	switch tab[i].Type {
	case TOFFON, TSSOLOSW:
		for j := 1; j < groupMax(&tab[i])+1; j++ {
			add(onOff(gi(tab, i, j) != 0))
		}
	case TCMONO:
		if gi(tab, i, 1) != 0 {
			add(" LCR")
		} else {
			add(" LR+M")
		}
		add(onOff(gi(tab, i, 2) != 0))
	case TCSOLO:
		add(Slevel(gf(tab, i, 1)), ge("Ssource", gi(tab, i, 2)), Slinf(gf(tab, i, 3), -18., 18., 1))
		for j := 4; j < 7; j++ {
			if gi(tab, i, j) != 0 {
				add(" AFL")
			} else {
				add(" PFL")
			}
		}
		for j := 7; j < 10; j++ {
			add(onOff(gi(tab, i, j) != 0))
		}
		add(Slinf(gf(tab, i, 10), -40., 0., 0))
		for j := 11; j < 14; j++ {
			add(onOff(gi(tab, i, j) != 0))
		}
		add(Slinf(gf(tab, i, 14), 0.3, 500., 1))
		for j := 15; j < 18; j++ {
			add(onOff(gi(tab, i, j) != 0))
		}
	case TCTALK:
		add(onOff(gi(tab, i, 1) != 0))
		if gi(tab, i, 2) != 0 {
			add(" EXT")
		} else {
			add(" INT")
		}
	case TCTALKAB:
		add(Slevel(gf(tab, i, 1)), onOff(gi(tab, i, 2) != 0), onOff(gi(tab, i, 3) != 0), Sbitmp(gi(tab, i, 4), 18))
	case TCOSC:
		add(Slevel(gf(tab, i, 1)), ge("f121", int32(120*gf(tab, i, 2)+0.5)), ge("f121", int32(120*gf(tab, i, 3)+0.5)))
		if gi(tab, i, 4) != 0 {
			add(" F2")
		} else {
			add(" F1")
		}
		add(ge("Sosct", gi(tab, i, 5)), Sint(gi(tab, i, 6)))
	case TCROUTSW:
		add(ge("Sroutin", gi(tab, i, 1)))
	case TCROUTIN, TCROUTPLAY:
		for j := 1; j < 5; j++ {
			add(ge("Sroutin", gi(tab, i, j)))
		}
		add(ge("Sroutax", gi(tab, i, 5)))
	case TCROUTAC:
		for j := 1; j < groupMax(&tab[i])+1; j++ {
			add(ge("Sroutac", gi(tab, i, j)))
		}
	case TCROUTOT:
		add(ge("Srouto1", gi(tab, i, 1)), ge("Srouto1", gi(tab, i, 2)), ge("Srouto2", gi(tab, i, 3)), ge("Srouto2", gi(tab, i, 4)))
	case TCCTRL:
		add(ge("Scolor", gi(tab, i, 1)))
	case TCENC:
		for j := 1; j < groupMax(&tab[i])+1; j++ {
			if gs(tab, i, j) != "" {
				add(" \"" + gs(tab, i, j) + "\"")
			} else {
				add(" \"-\"")
			}
		}
	case TCTAPE:
		add(Slinf(gf(tab, i, 1), -6., 24., 1), Slinf(gf(tab, i, 2), -6., 24., 1), onOff(gi(tab, i, 3) != 0))
	case TCMIX:
		add(onOff(gi(tab, i, 1) != 0), onOff(gi(tab, i, 2) != 0))
	case TCHCO:
		add(quoted(gs(tab, i, 1), " \"\""), Sint(gi(tab, i, 2)), ge("Scolor", gi(tab, i, 3)), Sint(gi(tab, i, 4)))
	case TCHDE:
		add(onOff(gi(tab, i, 1) != 0), Slinf(gf(tab, i, 2), 0.3, 500., 1))
	case TCHPR:
		add(Slinfs(gf(tab, i, 1), -18., 18., 1), onOff(gi(tab, i, 2) != 0), onOff(gi(tab, i, 3) != 0),
			ge("Sfslope", gi(tab, i, 4)), ge("f101", int32(100*gf(tab, i, 5)+0.5)))
	case TCHGA:
		add(onOff(gi(tab, i, 1) != 0), ge("Sgmode", gi(tab, i, 2)), Slinf(gf(tab, i, 3), -80., 0., 1),
			Slinf(gf(tab, i, 4), 3., 60., 1), Slinf(gf(tab, i, 5), 0., 120., 0),
			Slogf(gf(tab, i, 6), 0.02, 2000., 2), Slogf(gf(tab, i, 7), 5., 4000., 0), Sint(gi(tab, i, 8)))
	case TCHGF, TCHDF:
		add(onOff(gi(tab, i, 1) != 0), ge("Sgftype", gi(tab, i, 2)), ge("f201", int32(200*gf(tab, i, 3)+0.5)))
	case TCHDY:
		add(onOff(gi(tab, i, 1) != 0), ge("Sdmode", gi(tab, i, 2)), ge("Sddet", gi(tab, i, 3)), ge("Sdenv", gi(tab, i, 4)),
			Slinf(gf(tab, i, 5), -60., 0., 1), ge("Sdratio", gi(tab, i, 6)), Slinf(gf(tab, i, 7), 0., 5., 0),
			Slinf(gf(tab, i, 8), 0., 24., 1), Slinf(gf(tab, i, 9), 0., 120., 0), Slogf(gf(tab, i, 10), 0.02, 2000., 2),
			Slogf(gf(tab, i, 11), 5., 4000., 0), ge("Sdpos", gi(tab, i, 12)), Sint(gi(tab, i, 13)),
			Slinf(gf(tab, i, 14), 0., 100., 0), onOff(gi(tab, i, 15) != 0))
	case TCHIN:
		add(onOff(gi(tab, i, 1) != 0), ge("Sdpos", gi(tab, i, 2)), ge("Sinsel", gi(tab, i, 3)))
	case TCHEQ:
		add(ge("Setype", gi(tab, i, 1)), ge("f201", int32(200*gf(tab, i, 2)+0.5)),
			Slinfs(gf(tab, i, 3), -15., 15., 2), Slogf(gf(tab, i, 4), 10., 0.315, 1))
	case TCHMX:
		add(onOff(gi(tab, i, 1) != 0), Slevel(gf(tab, i, 2)), onOff(gi(tab, i, 3) != 0),
			Slinfs(gf(tab, i, 4), -100., 100., 0), onOff(gi(tab, i, 5) != 0), Slevel(gf(tab, i, 6)))
	case TCHMO:
		add(onOff(gi(tab, i, 1) != 0), Slevel(gf(tab, i, 2)), Slinfs(gf(tab, i, 3), -100., 100., 0),
			ge("Sctype", gi(tab, i, 4)), Sint(gi(tab, i, 5)))
	case TCHME:
		add(onOff(gi(tab, i, 1) != 0), Slevel(gf(tab, i, 2)))
	case TCHGRP:
		add(Sbitmp(gi(tab, i, 1), 8), Sbitmp(gi(tab, i, 2), 6))
	case TCHAMIX:
		add(ge("Samix", gi(tab, i, 1)), Slinfs(gf(tab, i, 2), -12., 12., 1))
	case TAXPR:
		add(Slinf(gf(tab, i, 1), -18., 18., 1), onOff(gi(tab, i, 2) != 0))
	case TBSCO:
		add(quoted(gs(tab, i, 1), " \"\""), Sint(gi(tab, i, 2)), ge("Scolor", gi(tab, i, 3)))
	case TMXPR:
		add(onOff(gi(tab, i, 1) != 0))
	case TMXDY:
		add(onOff(gi(tab, i, 1) != 0), ge("Sdmode", gi(tab, i, 2)), ge("Sddet", gi(tab, i, 3)), ge("Sdenv", gi(tab, i, 4)),
			Slinf(gf(tab, i, 5), -60., 0., 1), ge("Sdratio", gi(tab, i, 6)), Slinf(gf(tab, i, 7), 0., 5., 0),
			Slinf(gf(tab, i, 8), 0., 24., 1), Slinf(gf(tab, i, 9), 0., 120., 0), Slogf(gf(tab, i, 10), 0.02, 2000., 2),
			Slogf(gf(tab, i, 11), 5., 4000., 0), ge("Sdpos", gi(tab, i, 12)),
			Slinf(gf(tab, i, 14), 0., 100., 0), onOff(gi(tab, i, 15) != 0))
	case TMSMX:
		add(onOff(gi(tab, i, 1) != 0), Slevel(gf(tab, i, 2)), Slinfs(gf(tab, i, 4), -100., 100., 0))
	case TFXTYP1:
		add(ge("Sfxtyp1", gi(tab, i, 1)))
	case TFXTYP2:
		add(ge("Sfxtyp2", gi(tab, i, 1)))
	case TFXSRC:
		add(ge("Sfxsrc", gi(tab, i, 1)), ge("Sfxsrc", gi(tab, i, 2)))
	case TFXPAR1:
		add(s.fxParamsText(tab, i, false))
	case TFXPAR2:
		add(s.fxParamsText(tab, i, true))
	default:
		return "", false
	}
	return string(b), true
}
