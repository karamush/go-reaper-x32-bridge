package x32

import (
	"strings"

	"x32emu/osc"
)

// cursor walks the "<path> <value> [<value> ...]" payload of a "/" command,
// reproducing the way the C parsers of X32.c consume the input string.
type cursor struct{ s string }

// empty reports whether all parameters have been consumed.
func (c *cursor) empty() bool { return strings.TrimLeft(c.s, " ") == "" }

// word returns the next space delimited token.
func (c *cursor) word() (string, bool) {
	c.s = strings.TrimLeft(c.s, " ")
	if c.s == "" {
		return "", false
	}
	if i := strings.IndexAny(c.s, " \n"); i >= 0 {
		t := c.s[:i]
		c.s = c.s[i:]
		return t, true
	}
	t := c.s
	c.s = ""
	return t, true
}

// str returns the next string token, supporting "..." and '...' quoting
// (XslashSetString).
func (c *cursor) str() (string, bool) {
	c.s = strings.TrimLeft(c.s, " ")
	if c.s == "" {
		return "", false
	}
	if c.s[0] == '"' || c.s[0] == '\'' {
		q := c.s[0]
		if e := strings.IndexByte(c.s[1:], q); e >= 0 {
			v := c.s[1 : 1+e]
			c.s = c.s[2+e:]
			return v, true
		}
		v := c.s[1:]
		c.s = ""
		return v, true
	}
	if i := strings.IndexAny(c.s, " \n"); i >= 0 {
		v := c.s[:i]
		c.s = c.s[i:]
		return v, true
	}
	v := c.s
	c.s = ""
	return v, true
}

// ---------------------------------------------------------------- value setters
// Each setter stores the value and immediately notifies the other xremote
// clients with a single parameter message, exactly like Xfprint()+Xsend(S_REM)
// in X32.c.

// NotifyValue builds and sends the "addr,tag value" message of one parameter.
func (s *State) NotifyValue(c *Command) {
	if c == nil {
		return
	}
	b := osc.AppendString(nil, c.Addr)
	switch c.Type {
	case TF32:
		b = osc.AppendString(b, ",f")
		b = osc.AppendFloat32(b, c.ValF)
	case TS32:
		b = osc.AppendString(b, ",s")
		b = osc.AppendString(b, c.Str())
	default:
		b = osc.AppendString(b, ",i")
		b = osc.AppendInt32(b, c.ValI)
	}
	s.notify(b)
}

func (s *State) setInt(c *Command, v int32) {
	if v != c.ValI {
		c.ValI = v
		s.NotifyValue(c)
	}
}

func (s *State) setFloat(c *Command, v float32) {
	if v < c.ValF-Epsilon || v > c.ValF+Epsilon {
		c.ValF = v
		s.NotifyValue(c)
	}
}

func (s *State) setString(c *Command, v string) {
	if c.Str() != v {
		c.SetStr(v)
		s.NotifyValue(c)
	}
}

func (s *State) setList(c *Command, v string) {
	idx, ok := c.EnumIndex(v)
	if !ok {
		return
	}
	if idx != c.ValI {
		c.ValI = idx
		s.NotifyValue(c)
	}
}

// ---------------------------------------------------------------- applier
// applier applies the parameters of a "/" command to the entries following the
// group header, mirroring the switch() of function_slash() (X32.c), where every
// parser consumes exactly one parameter and stops at the end of the input.
type applier struct {
	s   *State
	tab Table
	idx int
	n   int  // successfully parsed parameters
	err bool // a parameter could not be parsed
}

func newApplier(s *State, tab Table, idx int) *applier {
	return &applier{s: s, tab: tab, idx: idx}
}

// firstFailed reproduces the C behaviour: a failure on the very first parameter
// keeps the whole command silent (return 0), while a later failure still echoes
// the request back to the sender (return S_SND).
func (a *applier) firstFailed() bool { return a.err && a.n == 0 }

func (a *applier) next() (*Command, bool) {
	if a.err || a.idx >= len(a.tab) {
		a.err = true
		return nil, false
	}
	c := &a.tab[a.idx]
	a.idx++
	return c, true
}

func (a *applier) done(ok bool) bool {
	if ok {
		a.n++
	} else {
		a.err = true
	}
	return ok
}

// int consumes an integer parameter.
func (a *applier) int() bool {
	c, ok := a.next()
	if !ok {
		return false
	}
	t, ok := a.s.cur.word()
	if !ok {
		return a.done(false)
	}
	a.s.setInt(c, parseIntPrefix(t))
	return a.done(true)
}

// str consumes a string parameter (quoted or not).
func (a *applier) str() bool {
	c, ok := a.next()
	if !ok {
		return false
	}
	t, ok := a.s.cur.str()
	if !ok {
		return a.done(false)
	}
	a.s.setString(c, t)
	return a.done(true)
}

// list consumes an enum parameter.
func (a *applier) list() bool {
	c, ok := a.next()
	if !ok {
		return false
	}
	t, ok := a.s.cur.word()
	if !ok {
		return a.done(false)
	}
	a.s.setList(c, t)
	return a.done(true)
}

// per consumes an integer or a "%0101" bitmap parameter.
func (a *applier) per() bool {
	c, ok := a.next()
	if !ok {
		return false
	}
	t, ok := a.s.cur.word()
	if !ok {
		return a.done(false)
	}
	if strings.HasPrefix(t, "%") {
		a.s.setInt(c, percentToI(t))
	} else {
		a.s.setInt(c, parseIntPrefix(t))
	}
	return a.done(true)
}

// lvl consumes a dB level parameter, quantized on steps steps.
func (a *applier) lvl(steps int) bool {
	c, ok := a.next()
	if !ok {
		return false
	}
	t, ok := a.s.cur.word()
	if !ok {
		return a.done(false)
	}
	a.s.setFloat(c, levelToF(t, steps))
	return a.done(true)
}

// logf consumes a physical value mapped on a log scale.
func (a *applier) logf(min, lmaxmin float32, steps int) bool {
	c, ok := a.next()
	if !ok {
		return false
	}
	t, ok := a.s.cur.word()
	if !ok {
		return a.done(false)
	}
	a.s.setFloat(c, logToF(XrFloat(t), min, lmaxmin, steps))
	return a.done(true)
}

// linf consumes a physical value mapped on a linear scale.
func (a *applier) linf(min, lmaxmin, step float32) bool {
	c, ok := a.next()
	if !ok {
		return false
	}
	t, ok := a.s.cur.word()
	if !ok {
		return a.done(false)
	}
	a.s.setFloat(c, linToF(XrFloat(t), min, lmaxmin, step))
	return a.done(true)
}

// float consumes a plain (already normalized) float parameter.
func (a *applier) float() bool {
	c, ok := a.next()
	if !ok {
		return false
	}
	t, ok := a.s.cur.word()
	if !ok {
		return a.done(false)
	}
	a.s.setFloat(c, XrFloat(t))
	return a.done(true)
}
