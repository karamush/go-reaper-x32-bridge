package x32

// slashGroup2 handles the remaining group types of function_slash().
func (s *State) slashGroup2(a *applier, typ Type) {
	switch typ {
	case TFXTYP1, TFXTYP2: // effect type
		a.list()
	case TFXSRC: // source (2 parameters)
		a.list()
		a.list()
	case TFXPAR1:
		s.applyFxParams(a, false)
	case TFXPAR2:
		s.applyFxParams(a, true)
	case TOMAIN: // src, pos, invert
		a.int()
		a.list()
		a.list()
	case TOMAIN2: // src, pos
		a.int()
		a.list()
	case TOP16: // 3 lists + int
		a.list()
		a.list()
		a.list()
		a.int()
	case TOMAIND: // on, time
		a.list()
		a.linf(0.3, 499.7, 0.1)
	case THAMP: // gain, phantom
		a.linf(-12., 72., 0.5)
		a.list()
	case TPREFS: // the whole /-prefs block
		a.str()
		a.linf(10., 90., 5.)
		a.linf(0., 100., 2.)
		a.linf(10., 90., 5.)
		a.linf(10., 90., 10.)
		for i := 0; i < 10; i++ {
			a.list()
		}
		a.per()
		for i := 0; i < 6; i++ {
			a.list()
		}
		a.str()
		a.list()
	case TPIR: // remote enable, protocol, port, ioenable
		a.list()
		a.list()
		a.list()
		a.per()
	case TPIQ: // iQmodel, IQeqset, IQsound
		a.list()
		a.list()
		a.int()
	case TPCARD: // 13 card settings
		for i := 0; i < 13; i++ {
			a.list()
		}
	case TPRTA: // rta settings
		a.list()
		a.linf(0., 60., 6.)
		a.list()
		a.int()
		a.list()
		a.list()
		a.per()
		a.list()
		a.logf(0.25, 4.158883083, 19) // log(16/0.25)
		a.list()
	case TPIP: // ip on
		a.list()
	case TPKEY: // recall key
		a.int()
	case TPADDR, TPMASK, TPGWAY: // 4 bytes
		a.int()
		a.int()
		a.int()
		a.int()
	case TSTAT: // the whole /-stat block
		a.list()
		a.int()
		a.int()
		a.list()
		a.int()
		a.int()
		a.list()
		a.list()
		a.int()
		a.list()
		a.int()
		a.list()
		a.list()
		a.list()
		a.list()
		a.list()
		a.list()
		a.int()
		a.int()
		a.list()
		a.list()
		a.int()
		a.int()
	case TSSCREEN: // screen, mutegrp, utils
		a.list()
		a.list()
		a.list()
	case TSCHA, TSMET, TSROU, TSSET, TSLIB, TSFX, TSMON, TSUSB, TSSCE, TSASS:
		a.list()
	case TSSOLOSW: // solo switch + 79 parameters
		a.list()
		for j := 2; j < 81; j++ {
			a.list()
		}
	case TSOSC:
		a.list()
	case TSTALK:
		a.list()
		a.list()
	case TUSB:
		a.str()
		a.str()
	case TSNAM: // show file
		a.str()
		a.int()
		a.int()
		a.int()
		a.int()
		a.int()
		a.int()
		a.int()
		a.int()
		a.int()
		a.int()
	case TSCUE: // cue point
		a.int()
		a.str()
		a.int()
		a.int()
		a.int()
		a.int()
		a.int()
		a.int()
		a.int()
	case TSSCN: // scene
		a.str()
		a.str()
		a.per()
		a.int()
	case TSSNP: // snippet
		a.str()
		a.int()
		a.int()
		a.int()
		a.int()
		a.int()
	case TUREC: // sd recorder
		a.int()
		a.int()
		a.int()
		a.int()
		a.int()
		a.list()
		a.int()
		a.int()
		a.int()
		a.int()
		a.list()
		a.list()
		a.str()
		a.str()
		a.str()
		a.int()
	case TSLIBS: // library entry
		a.int()
		a.str()
		a.int()
		a.per()
		a.int()
	case TD48: // DP48
		a.per()
		a.int()
	case TD48A:
		for j := 1; j < 49; j++ {
			a.int()
		}
	case TD48G:
		for j := 1; j < 13; j++ {
			a.str()
		}
	case TUROUO:
		for j := 1; j < 49; j++ {
			a.int()
		}
	case TUROUI:
		for j := 1; j < 33; j++ {
			a.int()
		}
	default:
		// SAES, STAPE, HA, ACTION, SOSC: no textual parameters in X32.c
	}
}
