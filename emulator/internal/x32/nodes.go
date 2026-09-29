package x32

import "strconv"

// nodeEntry is one entry of the Xnode[] table of X32.c: the path prefix and the
// command table it maps to. Order matters (specific paths come first).
type nodeEntry struct {
	name  string
	table Table
}

func pad2(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}

func pad3(n int) string {
	switch {
	case n < 10:
		return "00" + strconv.Itoa(n)
	case n < 100:
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}

// buildNodeTable reproduces the Xnode[] array of X32.c.
func buildNodeTable() []nodeEntry {
	n := []nodeEntry{
		{"conf", Xconfig},
		{"main", Xmain},
		{"-pre", Xprefs},
		{"-sta", Xstat},
	}
	for i := 1; i <= 32; i++ {
		n = append(n, nodeEntry{"ch/" + pad2(i), Xchannelset[i-1]})
	}
	n = append(n, nodeEntry{"ch", Xchannel01})
	for i := 1; i <= 8; i++ {
		n = append(n, nodeEntry{"auxin/" + pad2(i), Xauxinset[i-1]})
	}
	n = append(n, nodeEntry{"auxin", Xauxin01})
	for i := 1; i <= 8; i++ {
		n = append(n, nodeEntry{"fxrtn/" + pad2(i), Xfxrtnset[i-1]})
	}
	n = append(n, nodeEntry{"fxrtn", Xfxrtn01})
	for i := 1; i <= 8; i++ {
		n = append(n, nodeEntry{"fx/" + strconv.Itoa(i), Xfxset[i-1]})
	}
	n = append(n, nodeEntry{"fx", Xfx1})
	for i := 1; i <= 16; i++ {
		n = append(n, nodeEntry{"bus/" + pad2(i), Xbusset[i-1]})
	}
	n = append(n, nodeEntry{"bus", Xbus01})
	for i := 1; i <= 6; i++ {
		n = append(n, nodeEntry{"mtx/" + pad2(i), Xmtxset[i-1]})
	}
	n = append(n, nodeEntry{"mtx", Xmtx01})
	n = append(n, nodeEntry{"dca", Xdca})
	n = append(n, nodeEntry{"outputs/main/01", Xoutput},
		nodeEntry{"outputs/main", Xoutput},
		nodeEntry{"outputs", Xoutput})
	for i := 0; i <= 127; i++ {
		n = append(n, nodeEntry{"headamp/" + pad3(i), Xheadmpset[i]})
	}
	n = append(n, nodeEntry{"headamp", Xheadamp},
		nodeEntry{"-ha", Xmisc},
		nodeEntry{"-usb", Xmisc},
		nodeEntry{"undo", nil},
		nodeEntry{"-action", nil},
		nodeEntry{"-show/showfile/snippet", Xsnippet},
		nodeEntry{"-show/showfile/scene", Xscene},
		nodeEntry{"-show", Xshow},
		nodeEntry{"-urec", Xurec},
		nodeEntry{"-libs/fx", Xlibsf},
		nodeEntry{"-libs/r", Xlibsr},
		nodeEntry{"-libs", Xlibsc})
	return n
}

var nodeTable = buildNodeTable()
