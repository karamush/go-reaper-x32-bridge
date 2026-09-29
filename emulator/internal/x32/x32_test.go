package x32

import (
	"net"
	"os"
	"path/filepath"
	"testing"

	"x32emu/osc"
)

type sent struct {
	b    []byte
	addr string
}

func newTestState(t *testing.T) (*State, *[]sent) {
	t.Helper()
	ResetValues()
	st := NewState()
	st.Verbose = false
	st.Logf = func(string) {}
	sends := &[]sent{}
	st.SendFn = func(b []byte, to *net.UDPAddr) error {
		cp := make([]byte, len(b))
		copy(cp, b)
		*sends = append(*sends, sent{b: cp, addr: to.String()})
		return nil
	}
	return st, sends
}

func udpAddr(t *testing.T, s string) *net.UDPAddr {
	t.Helper()
	a, err := net.ResolveUDPAddr("udp4", s)
	if err != nil {
		t.Fatal(err)
	}
	return a
}

// rawPacket builds "<addr>\0<pad>" without any type tag block ("old" notation).
func rawPacket(addr string) []byte {
	return osc.AppendString(nil, addr)
}

func TestInfoReply(t *testing.T) {
	st, sends := newTestState(t)
	st.IP = "192.168.0.10"
	st.HandlePacket(osc.Start("/info", ""), udpAddr(t, "192.168.0.20:9999"))
	if len(*sends) != 1 {
		t.Fatalf("expected one answer, got %d", len(*sends))
	}
	m := osc.Decode((*sends)[0].b)
	if m.Addr != "/info" || m.Tags != "ssss" || len(m.Args) != 4 {
		t.Fatalf("decoded %+v", m)
	}
	if m.Args[0].S != "V2.07" || m.Args[1].S != DefaultName || m.Args[2].S != "X32" || m.Args[3].S != "4.06" {
		t.Fatalf("info args = %+v", m.Args)
	}
}

// The console name comes from "/-prefs/name" (set by the -name flag at startup
// or by a client later); an empty name falls back to DefaultName, which this port
// sets to "REAPER" instead of "X32 Emulator".
func TestConsoleName(t *testing.T) {
	st, sends := newTestState(t)
	st.IP = "192.168.0.10"
	st.SetName("My Desk")
	st.HandlePacket(osc.Start("/info", ""), udpAddr(t, "192.168.0.20:9999"))
	if got := osc.Decode((*sends)[0].b).Args[1].S; got != "My Desk" {
		t.Fatalf("name = %q, want My Desk", got)
	}
	st.SetName("")
	st.HandlePacket(osc.Start("/info", ""), udpAddr(t, "192.168.0.20:9999"))
	if got := osc.Decode((*sends)[1].b).Args[1].S; got != DefaultName {
		t.Fatalf("name = %q, want %q", got, DefaultName)
	}
	// a client renames the console with the regular command
	st.HandlePacket(osc.Build("/-prefs/name", "s", osc.Arg{Tag: 's', S: "From Client"}),
		udpAddr(t, "192.168.0.20:9999"))
	if got := st.consoleName(); got != "From Client" {
		t.Fatalf("name = %q, want From Client", got)
	}
}

func TestStatusAndXinfo(t *testing.T) {
	st, sends := newTestState(t)
	st.IP = "192.168.0.10"
	st.HandlePacket(osc.Start("/status", ""), udpAddr(t, "192.168.0.20:9999"))
	m := osc.Decode((*sends)[0].b)
	if m.Addr != "/status" || m.Tags != "sss" || m.Args[0].S != "active" || m.Args[1].S != "192.168.0.10" {
		t.Fatalf("status = %+v", m)
	}
	st.HandlePacket(osc.Start("/xinfo", ""), udpAddr(t, "192.168.0.20:9999"))
	m = osc.Decode((*sends)[1].b)
	if m.Addr != "/xinfo" || m.Args[0].S != "192.168.0.10" {
		t.Fatalf("xinfo = %+v", m)
	}
}

func TestGetFaderOldNotation(t *testing.T) {
	st, sends := newTestState(t)
	st.HandlePacket(rawPacket("/ch/01/mix/fader"), udpAddr(t, "192.168.0.20:9999"))
	if len(*sends) != 1 {
		t.Fatalf("expected one answer, got %d", len(*sends))
	}
	b := (*sends)[0].b
	if len(b) != 28 {
		t.Fatalf("answer length = %d, want 28", len(b))
	}
	m := osc.Decode(b)
	if m.Addr != "/ch/01/mix/fader" || m.Tags != "f" || m.Args[0].F != 0 {
		t.Fatalf("answer = %+v", m)
	}
}

func TestSetFaderNotifiesRemoteClientsOnly(t *testing.T) {
	st, sends := newTestState(t)
	a, b := udpAddr(t, "192.168.0.20:1000"), udpAddr(t, "192.168.0.21:1001")
	st.HandlePacket(rawPacket("/xremote"), b)
	if len(*sends) != 0 {
		t.Fatalf("/xremote must not answer, got %d messages", len(*sends))
	}
	if st.ClientCount() != 1 {
		t.Fatalf("client count = %d", st.ClientCount())
	}
	set := osc.Start("/ch/01/mix/fader", "f")
	set = osc.AppendFloat32(set, 0.5)
	st.HandlePacket(set, a)
	if len(*sends) != 1 {
		t.Fatalf("expected one notification, got %d", len(*sends))
	}
	if (*sends)[0].addr != b.String() {
		t.Fatalf("notification sent to %s, want %s", (*sends)[0].addr, b)
	}
	if string((*sends)[0].b) != string(set) {
		t.Fatalf("notification is not the echo of the request")
	}
	if st.faderValue() != 0.5 {
		t.Fatalf("fader = %v", st.faderValue())
	}
}

func TestSetSameValueIsSilent(t *testing.T) {
	st, sends := newTestState(t)
	a := udpAddr(t, "192.168.0.20:1000")
	set := osc.Start("/ch/01/mix/fader", "f")
	set = osc.AppendFloat32(set, 0.5)
	st.HandlePacket(set, a)
	*sends = (*sends)[:0]
	st.HandlePacket(set, a) // same value: nothing must be sent at all
	if len(*sends) != 0 {
		t.Fatalf("unchanged SET must be silent, got %d messages", len(*sends))
	}
}

func TestNodeGroupReply(t *testing.T) {
	st, sends := newTestState(t)
	req := osc.Build("/node", "s", osc.Arg{Tag: 's', S: "ch/01/config"})
	st.HandlePacket(req, udpAddr(t, "192.168.0.20:9999"))
	if len(*sends) != 1 {
		t.Fatalf("expected one answer, got %d", len(*sends))
	}
	m := osc.Decode((*sends)[0].b)
	if m.Addr != "node" || m.Tags != "s" || len(m.Args) != 1 {
		t.Fatalf("node reply = %+v", m)
	}
	if want := "/ch/01/config \"\" 0 OFF 0\n"; m.Args[0].S != want {
		t.Fatalf("node text = %q, want %q", m.Args[0].S, want)
	}
}

func TestSlashCommandSetsString(t *testing.T) {
	st, sends := newTestState(t)
	a, b := udpAddr(t, "192.168.0.20:1000"), udpAddr(t, "192.168.0.21:1001")
	st.HandlePacket(rawPacket("/xremote"), b)
	*sends = (*sends)[:0]
	req := osc.Build("/", "s", osc.Arg{Tag: 's', S: "/ch/01/config/name \"My Ch\""})
	st.HandlePacket(req, a)
	// one notification for the other client and the echo for the sender
	if len(*sends) != 2 {
		t.Fatalf("expected notification + echo, got %d", len(*sends))
	}
	if (*sends)[0].addr != b.String() {
		t.Fatalf("first message should go to the other client")
	}
	if got := st.channelName(); got != "My Ch" {
		t.Fatalf("name = %q", got)
	}
	n := osc.Decode((*sends)[0].b)
	if n.Addr != "/ch/01/config/name" || n.Tags != "s" || n.Args[0].S != "My Ch" {
		t.Fatalf("notification = %+v", n)
	}
	if string((*sends)[1].b) != string(req) {
		t.Fatalf("echo missing")
	}
}

// An unquoted string stops at the first space, exactly like XslashSetString.
func TestSlashCommandUnquotedStringStopsAtSpace(t *testing.T) {
	st, _ := newTestState(t)
	a := udpAddr(t, "192.168.0.20:1000")
	req := osc.Build("/", "s", osc.Arg{Tag: 's', S: "/ch/01/config/name My Ch"})
	st.HandlePacket(req, a)
	if got := st.channelName(); got != "My" {
		t.Fatalf("name = %q, want \"My\"", got)
	}
}

func TestMetersFrame(t *testing.T) {
	st, sends := newTestState(t)
	a := udpAddr(t, "192.168.0.20:1000")
	req := osc.Start("/meters", "siii")
	req = osc.AppendString(req, "/meters/1")
	req = osc.AppendInt32(req, 0)
	req = osc.AppendInt32(req, 0)
	req = osc.AppendInt32(req, 0)
	st.HandlePacket(req, a)
	if len(*sends) != 0 {
		t.Fatalf("the meter request itself is not answered")
	}
	st.PumpMeters()
	if len(*sends) != 1 {
		t.Fatalf("expected one meter frame, got %d", len(*sends))
	}
	frame := (*sends)[0].b
	if len(frame) != 24+4*meterSizes[1] {
		t.Fatalf("frame length = %d", len(frame))
	}
	m := osc.Decode(frame)
	if m.Addr != "/meters/1" || m.Tags != "b" {
		t.Fatalf("frame = %+v", m)
	}
	if len(m.Args[0].B) != 4*(meterSizes[1]+1) {
		t.Fatalf("blob length = %d", len(m.Args[0].B))
	}
}

func TestPersistence(t *testing.T) {
	st, _ := newTestState(t)
	st.setFaderForTest(0.25)
	st.setNameForTest("My Channel")
	path := filepath.Join(t.TempDir(), ".X32res.rc")
	if err := st.SaveResource(path); err != nil {
		t.Fatal(err)
	}
	ResetValues()
	st2, _ := newTestState(t)
	if st2.X32Init(path) {
		t.Fatalf("resource file not found")
	}
	if got := st2.faderValue(); got != 0.25 {
		t.Fatalf("restored fader = %v", got)
	}
	if got := st2.channelName(); got != "My Channel" {
		t.Fatalf("restored name = %q", got)
	}
	_ = os.Remove(path)
}

// ---------------------------------------------------------------- test helpers
func (s *State) tableValue(addr string) *Command {
	for _, tab := range persistTables() {
		for i := range tab {
			if tab[i].Addr == addr {
				return &tab[i]
			}
		}
	}
	return nil
}

func (s *State) faderValue() float32 {
	if c := s.tableValue("/ch/01/mix/fader"); c != nil {
		return c.ValF
	}
	return -1
}

func (s *State) setFaderForTest(v float32) {
	if c := s.tableValue("/ch/01/mix/fader"); c != nil {
		c.ValF = v
	}
}

func (s *State) channelName() string {
	if c := s.tableValue("/ch/01/config/name"); c != nil {
		return c.Str()
	}
	return ""
}

func (s *State) setNameForTest(v string) {
	if c := s.tableValue("/ch/01/config/name"); c != nil {
		c.SetStr(v)
	}
}
