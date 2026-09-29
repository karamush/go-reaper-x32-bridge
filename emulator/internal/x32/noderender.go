package x32

import (
	"strconv"
	"strings"

	"x32emu/osc"
)

// functionNode is the port of function_node(): "/node ,s <path>" either returns
// the whole group of values as one text string, or - when <path> designates a
// single parameter - re-dispatches it as a normal command and wraps the answer.
func (s *State) functionNode() int {
	payload := ""
	if len(s.Msg.Args) > 0 && s.Msg.Args[0].Tag == 's' {
		payload = s.Msg.Args[0].S
	}
	var tab Table
	for i := range nodeTable {
		if strings.HasPrefix(payload, nodeTable[i].name) {
			tab = nodeTable[i].table
			break
		}
	}
	if tab == nil {
		return 0
	}
	for i := range tab {
		if tab[i].Flags == FFND {
			if !strings.HasPrefix(tab[i].Addr[1:], payload) {
				continue
			}
			for i+1 < len(tab) && tab[i+1].Flags == FFND {
				i++
			}
			out := osc.Start("node", "s")
			txt := tab[i].Addr + s.renderGroup(tab, i)
			s.Out = osc.AppendString(out, txt+"\n")
			return SSnd
		}
		if strings.HasPrefix(tab[i].Addr[1:], payload) {
			// change the request into the equivalent single command:
			// "/node ,s -prefs/rta/visibility" becomes "/-prefs/rta/visibility"
			start := 12
			if start < len(s.Raw) && s.Raw[start] == '/' {
				start = 13
			}
			pkt := append([]byte{'/'}, s.Raw[start:]...)
			// the handler must work on the rebuilt command
			s.Raw = pkt
			s.Msg = osc.Decode(pkt)
			key := make([]byte, 4)
			copy(key, pkt)
			for h := range headers {
				if headers[h].key == string(key) {
					headers[h].fn(s)
					break
				}
			}
			return s.functionNodeSingle()
		}
	}
	return 0
}

// functionNodeSingle wraps the answer of a single command into a "node ,s"
// message (function_node_single in X32.c).
func (s *State) functionNodeSingle() int {
	c := s.nodeSingleCmd
	if c == nil {
		return 0
	}
	addr := s.Msg.Addr
	out := osc.Start("node", "s")
	s.Out = osc.AppendString(out, addr+s.nodeSingleValue(c)+"\n")
	return SSnd
}

// nodeSingleValue renders the value of one parameter (function_node_single).
func (s *State) nodeSingleValue(c *Command) string {
	if c.Enum != nil {
		if v := c.EnumValue(c.ValI); v != "" {
			return v
		}
	}
	switch c.Type {
	case TI32:
		return Sint(c.ValI)
	case TF32:
		return " " + strconv.FormatFloat(float64(c.ValF), 'f', 6, 32)
	case TS32:
		return c.Str() // no leading space, like X32.c
	case TP32:
		var sb strings.Builder
		sb.WriteString(" %")
		il := 31
		for il > 0 && c.ValI&(1<<uint(il)) == 0 {
			il--
		}
		for ; il >= 0; il-- {
			if c.ValI&(1<<uint(il)) != 0 {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
		}
		return sb.String()
	}
	return " TODO"
}

// groupMax returns the number of data entries of a group header, using the
// explicit count when the table provides one and the C default (1) otherwise.
func groupMax(c *Command) int {
	if c.Count > 0 {
		return int(c.Count)
	}
	return 0
}
