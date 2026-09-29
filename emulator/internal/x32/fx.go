package x32

import "strings"

// fxLookupCode returns the parameter type string of an effect (Sflookup /
// Sflookup2 in X32.c): one character per parameter, 'i' = integer, 'f' = float,
// 'e' = enum, 's' = string, '-' = unused.
func fxLookupCode(slotType int32, second bool) string {
	name := "Sflookup"
	if second {
		name = "Sflookup2"
	}
	v := enumFor(name)
	if slotType < 0 || int(slotType) >= len(v) {
		return ""
	}
	return strings.TrimRight(v[slotType], "-")
}

// fxSlotType returns the effect type index loaded in an FX slot (the value of
// "/fx/<n>/type").
func fxSlotType(tab Table, slot int) int32 {
	addr := "/fx/" + string(rune('0'+slot)) + "/type"
	for i := range tab {
		if tab[i].Addr == addr {
			return tab[i].ValI
		}
	}
	return 0
}

// fxArgType resolves the type of an FX parameter (FXc_lookup in X32.c).
func (s *State) fxArgType(tab Table, idx int) Type {
	addr := tab[idx].Addr
	if len(addr) < 12 {
		return TNIL
	}
	par := int(addr[10]-'0')*10 + int(addr[11]-'0') - 1
	if par < 0 || par > 63 {
		return TNIL
	}
	slot := int(addr[4] - '0')
	if slot < 1 || slot > 8 {
		return TNIL
	}
	code := fxLookupCode(fxSlotType(tab, slot), slot > 4)
	if par >= len(code) {
		return TF32
	}
	switch code[par] {
	case 'i':
		return TI32
	case 'e':
		return TE32
	case 's':
		return TS32
	}
	return TF32
}

// applyFxParams consumes the parameters of an FX effect, using the parameter
// count and types of the effect as defined by Sflookup / Sflookup2.
//
// NOTE: the physical scaling of each parameter (SetFxPar1 in X32.c) is not
// implemented yet - values are stored normalized, which is what clients send in
// the OSC form. The textual "/" form therefore expects normalized values too.
func (s *State) applyFxParams(a *applier, second bool) {
	if a.idx == 0 || a.idx > len(a.tab) {
		a.err = true
		return
	}
	slot := int(a.tab[a.idx-1].Addr[4] - '0')
	code := fxLookupCode(fxSlotType(a.tab, slot), second)
	for k := 0; k < len(code); k++ {
		switch code[k] {
		case 'i', 'e':
			if !a.int() {
				return
			}
		default:
			if !a.float() {
				return
			}
		}
	}
}

// fxParamsText renders every parameter of an FX slot for /node replies
// (GetFxPar1 in X32.c formats each value with the scale of the effect).
func (s *State) fxParamsText(tab Table, base int, second bool) string {
	if base < 0 || base >= len(tab) {
		return ""
	}
	slot := int(tab[base].Addr[4] - '0')
	code := fxLookupCode(fxSlotType(tab, slot), second)
	var sb strings.Builder
	for k := 0; k < len(code) && base+1+k < len(tab); k++ {
		c := &tab[base+1+k]
		switch code[k] {
		case 'i', 'e':
			sb.WriteString(Sint(c.ValI))
		default:
			sb.WriteString(Slinf(c.ValF, 0., 1., 4))
		}
	}
	return sb.String()
}
