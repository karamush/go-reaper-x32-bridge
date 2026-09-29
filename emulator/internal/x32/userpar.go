package x32

import "x32emu/osc"

// The "/-stat/userpar/NN/value" parameters are the User Assign controls of a
// console: a real X32 pushes their value whenever a knob is turned or a button
// is pressed, which is how X32ReaperW (and other control surface bridges) get
// transport and marker commands from the desk.
//
// They are NOT part of the C reference emulator (X32.c has no such address in
// its tables - see docs/X32_OSC_PROTOCOL.md 14.2), so this table is hand written
// and, deliberately, not persisted: ".X32res.rc" stays byte compatible with the
// resource files of the C emulator.
//
// Numbering follows the hardware:
//
//	buttons  A 1..8, B 9..16, C 17..24, D 25..32   (Bank C = 17..24)
//	encoders 33..36 are the four rows of Bank C
//
// X32ReaperW maps its "Bank C button" numbers 5..12 through UserParIndex()
// (index = 12 + button number), sets a button to 0 on init, and acts on the
// release report (value 0), which is the "button up" transition.
const (
	UserParButtonFirst = 17                // first Bank C button
	UserParButtonLast  = 24                // last Bank C button
	UserParEncFirst    = 33                // first Bank C encoder
	UserParEncLast     = 36                // last Bank C encoder
	UserParCenter      = 64                // centered encoder value
	UserParAddrPrefix  = "/-stat/userpar/" // address is <prefix><NN>/value
	UserParAddrSuffix  = "/value"          //
)

// UserParIndex converts a Bank C button number (5..12) into its userpar index
// (17..24), as used by X32ReaperW ("/-stat/userpar/%2d/value", 12 + button).
func UserParIndex(button int) int { return 12 + button }

// Xuserpar is the hand written table of the user assign controls. The address is
// 23 characters long ("/-stat/userpar/18/value"), which matters: bridge clients
// parse the fixed byte offsets of the notification packet, so the length must
// not change.
var Xuserpar = newUserParTable()

func newUserParTable() Table {
	t := make(Table, 0, UserParButtonLast-UserParButtonFirst+1+UserParEncLast-UserParEncFirst+1)
	for i := UserParButtonFirst; i <= UserParButtonLast; i++ {
		t = append(t, Command{Addr: userParAddr(i), Type: TI32, Flags: FXET})
	}
	for i := UserParEncFirst; i <= UserParEncLast; i++ {
		c := Command{Addr: userParAddr(i), Type: TI32, Flags: FXET, ValI: UserParCenter}
		t = append(t, c)
	}
	return t
}

func userParAddr(i int) string {
	return UserParAddrPrefix + itoa(i) + UserParAddrSuffix
}

// userParEntry returns the table entry of a userpar index.
func userParEntry(i int) *Command {
	addr := userParAddr(i)
	for j := range Xuserpar {
		if Xuserpar[j].Addr == addr {
			return &Xuserpar[j]
		}
	}
	return nil
}

// UserPar returns the current value of a user assign control (false when the
// index is not one of the Bank C controls).
func UserPar(i int) (int32, bool) {
	if c := userParEntry(i); c != nil {
		return c.ValI, true
	}
	return 0, false
}

// PressUserPar simulates a physical user assign control event: it stores the new
// value and pushes "/-stat/userpar/NN/value ,i <v>" to every registered
// /xremote client, exactly like a console reporting a button or knob movement.
//
// This is the entry point for a GUI, a web front end or a bridge: there is no
// physical control on an emulator, so something has to generate the events.
// It returns false for an unknown index.
func (s *State) PressUserPar(i int, v int32) bool {
	c := userParEntry(i)
	if c == nil {
		return false
	}
	if c.ValI == v {
		return true // no change, like the console event filter of functParams
	}
	c.ValI = v
	b := osc.AppendInt32(osc.Start(c.Addr, "i"), v)
	s.NotifyAll(b)
	return true
}
