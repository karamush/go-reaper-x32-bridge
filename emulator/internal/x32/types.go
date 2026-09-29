// Package x32 implements the Behringer X32 console emulator from the
// X32-Behringer repository (X32.c v0.88) in Go: UDP/OSC server, command
// tables, /xremote clients, meters, persistence.
package x32

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Constants copied from X32.c.
const (
	Epsilon     = 0.0001  // float comparison epsilon
	BSize       = 512     // communication buffers
	XVersion    = "4.06"  // firmware version reported by /info, /xinfo, /node showdump
	MaxClients  = 4       // number of xremote clients supported
	XRemoteTime = 11      // seconds before an xremote client is considered gone
	MaxMeters   = 17      // /meters/0 ... /meters/16
	UDPPort     = "10023" // desk port (10024 is XAir18)
)

// Flags of a table entry (F_GET, F_SET, F_XET, F_NPR, F_FND in X32.c).
type Flags uint8

const (
	FGET Flags = 0x01 // F_GET
	FSET Flags = 0x02 // F_SET
	FXET Flags = FGET | FSET
	FNPR Flags = 0x04 // F_NPR (unused by the tables)
	FFND Flags = 0x08 // F_FND: group header / first of a series
)

// Send masks returned by command handlers (S_SND / S_REM in X32.c).
const (
	SSnd = 0x01 // answer the requesting client
	SRem = 0x02 // push the answer to the other /xremote clients
)

// Command is one entry of a command table and also holds the live value of
// that parameter (the tables are the console state, exactly like in X32.c).
type Command struct {
	Addr  string
	Type  Type
	Flags Flags
	Count int32 // for F_FND entries: number of following items (0 = implicit)
	Enum  *Enum // enum string list, nil when not applicable
	ValI  int32
	ValF  float32
	ValS  *string
}

// Table is a command table (X32command[] in X32.c).
type Table []Command

// Str returns the string value ("" when unset, like a NULL char* in C).
func (c *Command) Str() string {
	if c.ValS == nil {
		return ""
	}
	return *c.ValS
}

// SetStr stores a string value ("" clears it, matching free()+NULL in C).
func (c *Command) SetStr(v string) {
	if v == "" {
		c.ValS = nil
		return
	}
	s := v
	c.ValS = &s
}

// EnumIndex looks a value up in the entry enum list. The comparison ignores
// every leading space because the C tables store " ON" while OSC clients send
// "ON" (XslashSetList compares command->node[j]+1).
func (c *Command) EnumIndex(v string) (int32, bool) {
	if c.Enum == nil {
		return 0, false
	}
	for i, s := range c.Enum.Values {
		if s == "" {
			continue // sentinel
		}
		if strings.TrimLeft(s, " ") == v {
			return int32(i), true
		}
	}
	return 0, false
}

// EnumValue returns the raw enum string (with its leading space) at index i.
func (c *Command) EnumValue(i int32) string {
	if c.Enum == nil || i < 0 || int(i) >= len(c.Enum.Values) {
		return ""
	}
	return c.Enum.Values[i]
}

// enumFor returns a named enum list (used by the node renderer for arrays that
// are referenced directly in X32.c, e.g. Scolor or Sfxtyp1).
func enumFor(name string) []string {
	if e, ok := enums[name]; ok {
		return e.Values
	}
	return nil
}

// enumRaw returns values[idx] of a named enum list, "" when out of range.
func enumRaw(name string, idx int32) string {
	v := enumFor(name)
	if idx < 0 || int(idx) >= len(v) {
		return ""
	}
	return v[idx]
}

// ---------------------------------------------------------------- formatting
// The following helpers reproduce the Sxxx() functions of X32.c. Strings carry
// a leading space because they are concatenated in /node replies.

// Slevel converts a linear fader value [0..1] to a dB string (Slevel()).
func Slevel(f float32) string {
	if f <= 0 {
		return " -oo"
	}
	var fl float32
	switch {
	case f <= 0.0625:
		fl = 30./0.0625*f - 90.
	case f <= 0.25:
		fl = 30./(0.25-0.0625)*(f-0.0625) - 60.
	case f < 0.5:
		fl = 20./(0.5-0.25)*(f-0.25) - 30.
	default:
		fl = 20./(1.-0.5)*(f-0.5) - 10.
	}
	return sprintfFloat(" %+.1f", fl)
}

// Slinf converts a linear normalized value to its physical range (Slinf()).
func Slinf(fin, fmin, fmax float32, pre int) string {
	return sprintfFloat(" %."+strconv.Itoa(pre)+"f", fmin+(fmax-fmin)*fin)
}

// Slinfs is Slinf with an explicit sign (Slinfs()).
func Slinfs(fin, fmin, fmax float32, pre int) string {
	return sprintfFloat(" %+."+strconv.Itoa(pre)+"f", fmin+(fmax-fmin)*fin)
}

// Slogf converts a log normalized value to its physical range (Slogf()).
func Slogf(fin, fmin, fmax float32, pre int) string {
	v := float32(math.Exp(float64(fin)*math.Log(float64(fmax/fmin)) + math.Log(float64(fmin))))
	return sprintfFloat(" %."+strconv.Itoa(pre)+"f", v)
}

// Sbitmp renders an integer as a %bitmap string (Sbitmp()).
func Sbitmp(iin int32, l int) string {
	var sb strings.Builder
	sb.WriteString(" %")
	for i := l - 1; i > -1; i-- {
		if iin&(1<<uint(i)) != 0 {
			sb.WriteByte('1')
		} else {
			sb.WriteByte('0')
		}
	}
	return sb.String()
}

// Sint renders an integer with a leading space (Sint()).
func Sint(iin int32) string {
	return " " + strconv.FormatInt(int64(iin), 10)
}

// onOff renders " ON" / " OFF".
func onOff(b bool) string {
	if b {
		return " ON"
	}
	return " OFF"
}

// quoted renders " \"text\"" with the empty variant " \" \"", as X32.c does.
func quoted(s string, empty string) string {
	if s == "" {
		return empty
	}
	return " \"" + s + "\""
}

func sprintfFloat(format string, f float32) string {
	return fmt.Sprintf(format, f)
}

// XrFloat parses a float the way Xr_float() does: it accepts "12", "12.5" and
// the "12k5" (== 12500) notation used by several clients.
func XrFloat(s string) float32 {
	if i := strings.IndexByte(s, '.'); i >= 0 {
		return parseFloatPrefix(s)
	}
	i := strings.IndexByte(s, 'k')
	if i < 0 {
		return parseFloatPrefix(s)
	}
	l := len(s)
	ival := parseFloatPrefix(s[:i])
	idec := parseFloatPrefix(s[i+1:])
	f := ival * 1000.
	switch l - i {
	case 2:
		f += idec * 100.
	case 3:
		f += idec * 10.
	case 4:
		f += idec
	}
	return f
}

// parseFloatPrefix mimics sscanf("%f"): leading spaces are skipped and parsing
// stops at the first character that cannot belong to a number.
func parseFloatPrefix(s string) float32 {
	s = strings.TrimLeft(s, " \t")
	end := 0
	for end < len(s) {
		c := s[end]
		if (c >= '0' && c <= '9') || c == '+' || c == '-' || c == '.' || c == 'e' || c == 'E' {
			end++
			continue
		}
		break
	}
	v, err := strconv.ParseFloat(s[:end], 32)
	if err != nil {
		return 0
	}
	return float32(v)
}

// parseIntPrefix mimics sscanf("%d").
func parseIntPrefix(s string) int32 {
	s = strings.TrimLeft(s, " \t")
	end := 0
	if end < len(s) && (s[end] == '+' || s[end] == '-') {
		end++
	}
	for end < len(s) && s[end] >= '0' && s[end] <= '9' {
		end++
	}
	v, err := strconv.ParseInt(s[:end], 10, 32)
	if err != nil {
		return 0
	}
	return int32(v)
}

// levelToF converts a dB level into the normalized fader value, quantized on
// nsteps steps (XslashSetLevl / Xp_level in X32.c).
func levelToF(s string, nsteps int) float32 {
	if strings.HasPrefix(s, "-oo") {
		return 0
	}
	f := XrFloat(s)
	q := func(v float32, n int) float32 { return float32(int(v*(float32(n)+0.5))) / float32(n) }
	switch {
	case f < -60.:
		f = f*0.00208333333 + 0.1875
		f = q(f, nsteps)
		if f < 0 {
			f = 0
		}
	case f < -30.:
		f = 0.00625*f + 0.4375
		f = q(f, nsteps)
	case f < -10.:
		f = 0.0125*f + 0.625
		f = q(f, nsteps)
	case f <= 10.:
		f = f*0.025 + 0.75
		if f = q(f, nsteps); f > 1.0 {
			f = 1.0
		}
	default:
		f = 1.0
	}
	return f
}

// logToF converts a physical value into a normalized log value, quantized on
// nsteps steps (XslashSetLogf). lmaxmin == log(xmax/xmin).
func logToF(v, xmin, lmaxmin float32, nsteps int) float32 {
	f := float32(math.Log(float64(v/xmin))) / lmaxmin
	f = float32(math.Round(float64(f*float32(nsteps)))) / float32(nsteps)
	if f <= 0. {
		f = 0. // avoid -0.0 values (0x80000000)
	}
	if f > 1. {
		f = 1.
	}
	return f
}

// linToF converts a physical value into a normalized linear value, quantized
// on xstep steps (XslashSetLinf). lmaxmin == xmax-xmin.
func linToF(v, xmin, lmaxmin, xstep float32) float32 {
	f := (v - xmin) / lmaxmin
	st := lmaxmin / xstep
	f = float32(math.Round(float64(f*st))) / st
	if f <= 0. {
		f = 0.
	}
	if f > 1. {
		f = 1.
	}
	return f
}

// percentToI decodes a "%0101" bitmap into an integer. Note: X32.c has a bug
// here (it shifts the old value instead of the accumulator and only keeps bit
// 0); this port implements the inverse of Sbitmp(), which is what clients
// expect and what a real console does.
func percentToI(s string) int32 {
	var v int32
	for i := 1; i < len(s); i++ {
		v <<= 1
		if s[i] == '1' {
			v |= 1
		}
	}
	return v
}
