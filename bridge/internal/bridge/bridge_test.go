package bridge

import (
	"testing"

	"x32emu/osc"
)

// newTestBridge builds a bridge without sockets: fromReaper / fromX32 only queue
// the outgoing datagrams, which the tests inspect directly.
func newTestBridge(t *testing.T, tweak func(*Config)) *Bridge {
	t.Helper()
	cfg := DefaultConfig()
	cfg.Verbose = false
	if tweak != nil {
		tweak(&cfg)
	}
	b, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	b.log = func(string, ...any) {}
	return b
}

func (b *Bridge) sentToReaper() []*osc.Message {
	var out []*osc.Message
	for _, m := range b.pending {
		if m.toReaper {
			out = append(out, osc.Decode(m.b))
		}
	}
	return out
}

func (b *Bridge) sentToX32() []*osc.Message {
	var out []*osc.Message
	for _, m := range b.pending {
		if !m.toReaper {
			out = append(out, osc.Decode(m.b))
		}
	}
	return out
}

func addrs(ms []*osc.Message) []string {
	out := make([]string, 0, len(ms))
	for _, m := range ms {
		out = append(out, m.Addr)
	}
	return out
}

func hasAddr(ms []*osc.Message, addr string) bool {
	for _, m := range ms {
		if m.Addr == addr {
			return true
		}
	}
	return false
}

func TestTrackStripMapping(t *testing.T) {
	b := newTestBridge(t, nil)
	if strip, ok := b.cfg.TrackToX32(1); !ok || strip != 1 {
		t.Fatalf("track 1 -> strip %d, %v", strip, ok)
	}
	if strip, ok := b.cfg.TrackToX32(32); !ok || strip != 32 {
		t.Fatalf("track 32 -> strip %d, %v", strip, ok)
	}
	if _, ok := b.cfg.TrackToX32(33); ok {
		t.Fatal("track 33 must be outside the channel bank")
	}
	if track, ok := b.cfg.X32ToTrack(32); !ok || track != 32 {
		t.Fatalf("strip 32 -> track %d, %v", track, ok)
	}
	if max := b.cfg.MaxBankOffset(); max != 0 {
		t.Fatalf("MaxBankOffset = %d, want 0 (tracks 1..32)", max)
	}
}

// A wider track range gives more banks: with 64 tracks the desk can show two
// banks of 32 channels.
func TestTrackStripMappingAcrossBanks(t *testing.T) {
	b := newTestBridge(t, func(c *Config) { c.TrkMin, c.TrkMax = 1, 64 })
	if max := b.cfg.MaxBankOffset(); max != 1 {
		t.Fatalf("MaxBankOffset = %d, want 1", max)
	}
	b.cfg.BankOffset = 1
	if strip, ok := b.cfg.TrackToX32(33); !ok || strip != 1 {
		t.Fatalf("bank 1: track 33 -> strip %d, %v", strip, ok)
	}
	if track, ok := b.cfg.X32ToTrack(1); !ok || track != 33 {
		t.Fatalf("bank 1: strip 1 -> track %d, %v", track, ok)
	}
	if track, ok := b.cfg.X32ToTrack(32); !ok || track != 64 {
		t.Fatalf("bank 1: strip 32 -> track %d, %v", track, ok)
	}
	if _, ok := b.cfg.TrackToX32(1); ok {
		t.Fatal("bank 1: track 1 must not be visible")
	}
}

func TestSectionMapping(t *testing.T) {
	b := newTestBridge(t, func(c *Config) {
		c.AuxMin, c.AuxMax = 33, 40
		c.FxrMin, c.FxrMax = 41, 48
		c.BusMin, c.BusMax = 49, 64
	})
	if track, ok := b.cfg.X32ToTrack(33); !ok || track != 33 {
		t.Fatalf("auxin 1 -> %d, %v", track, ok)
	}
	if strip, ok := b.cfg.TrackToX32(48); !ok || strip != StripFxrtnLast {
		t.Fatalf("track 48 -> strip %d, %v", strip, ok)
	}
	if strip, ok := b.cfg.TrackToX32(64); !ok || strip != StripBusLast {
		t.Fatalf("track 64 -> strip %d, %v", strip, ok)
	}
	if _, ok := b.cfg.TrackToX32(65); ok {
		t.Fatal("track 65 must not be mapped")
	}
	d := newTestBridge(t, nil)
	if _, ok := d.cfg.X32ToTrack(33); ok {
		t.Fatal("auxin must be disabled by default")
	}
}

func TestFaderBothWays(t *testing.T) {
	b := newTestBridge(t, nil)
	b.fromX32(osc.Build("/ch/03/mix/fader", "f", osc.Arg{Tag: 'f', F: 0.75}))
	r := b.sentToReaper()
	if len(r) != 1 || r[0].Addr != "/track/3/volume" || r[0].Args[0].F != 0.75 {
		t.Fatalf("REAPER got %v", addrs(r))
	}
	b.pending = nil
	b.fromReaper(osc.Build("/track/3/volume", "f", osc.Arg{Tag: 'f', F: 0.25}))
	x := b.sentToX32()
	if len(x) != 1 || x[0].Addr != "/ch/03/mix/fader" || x[0].Args[0].F != 0.25 {
		t.Fatalf("desk got %v", addrs(x))
	}
	if st := b.state.peek(3); st == nil || st.fader != 0.25 {
		t.Fatalf("cached state = %+v", st)
	}
}

func TestMuteIsInverted(t *testing.T) {
	b := newTestBridge(t, nil)
	b.fromReaper(osc.Build("/track/4/mute", "f", osc.Arg{Tag: 'f', F: 1}))
	x := b.sentToX32()
	if len(x) != 1 || x[0].Addr != "/ch/04/mix/on" || x[0].Args[0].I != 0 {
		t.Fatalf("desk got %v", x)
	}
	b.pending = nil
	b.fromX32(osc.Build("/ch/04/mix/on", "i", osc.Arg{Tag: 'i', I: 1}))
	r := b.sentToReaper()
	if len(r) != 1 || r[0].Addr != "/track/4/mute" || r[0].Args[0].F != 0 {
		t.Fatalf("REAPER got %v", r)
	}
	b.pending = nil
	b.fromX32(osc.Build("/ch/04/mix/on", "i", osc.Arg{Tag: 'i', I: 0}))
	if r = b.sentToReaper(); len(r) != 1 || r[0].Args[0].F != 1 {
		t.Fatalf("REAPER got %v", r)
	}
}

func TestSendLevelAndOffset(t *testing.T) {
	b := newTestBridge(t, func(c *Config) { c.SendOffset = 1 })
	b.fromX32(osc.Build("/ch/02/mix/05/level", "f", osc.Arg{Tag: 'f', F: 0.5}))
	r := b.sentToReaper()
	if len(r) != 1 || r[0].Addr != "/track/2/send/6/volume" {
		t.Fatalf("REAPER got %v", addrs(r))
	}
	b.pending = nil
	b.fromReaper(osc.Build("/track/2/send/6/volume", "f", osc.Arg{Tag: 'f', F: 0.75}))
	x := b.sentToX32()
	if len(x) != 1 || x[0].Addr != "/ch/02/mix/05/level" || x[0].Args[0].F != 0.75 {
		t.Fatalf("desk got %v", addrs(x))
	}
}

func TestSelectAndSolo(t *testing.T) {
	b := newTestBridge(t, nil)
	// the desk reports the selected strip as a zero based index
	b.fromX32(osc.Build("/-stat/selidx", "i", osc.Arg{Tag: 'i', I: 5}))
	r := b.sentToReaper()
	if len(r) != 2 || r[0].Addr != "/action/40297" || r[1].Addr != "/track/6/select" {
		t.Fatalf("REAPER got %v", addrs(r))
	}
	b.pending = nil
	b.fromX32(osc.Build("/-stat/solosw/07", "i", osc.Arg{Tag: 'i', I: 1}))
	if r = b.sentToReaper(); len(r) != 1 || r[0].Addr != "/track/7/solo" || r[0].Args[0].F != 1 {
		t.Fatalf("REAPER got %v", addrs(r))
	}
	b.pending = nil
	b.fromReaper(osc.Build("/track/7/solo", "f", osc.Arg{Tag: 'f', F: 1}))
	x := b.sentToX32()
	if len(x) != 1 || x[0].Addr != "/-stat/solosw/07" || x[0].Args[0].I != 1 {
		t.Fatalf("desk got %v", addrs(x))
	}
}

func TestMasterAndNameAndDca(t *testing.T) {
	b := newTestBridge(t, func(c *Config) { c.RdcaMin[0], c.RdcaMax[0] = 10, 12 })
	b.fromReaper(osc.Build("/master/volume", "f", osc.Arg{Tag: 'f', F: 0.6}))
	x := b.sentToX32()
	if len(x) != 1 || x[0].Addr != "/main/st/mix/fader" {
		t.Fatalf("desk got %v", addrs(x))
	}
	b.pending = nil
	b.fromX32(osc.Build("/main/st/mix/fader", "f", osc.Arg{Tag: 'f', F: 0.7}))
	r := b.sentToReaper()
	if len(r) != 1 || r[0].Addr != "/master/volume" {
		t.Fatalf("REAPER got %v", addrs(r))
	}
	b.pending = nil
	b.fromX32(osc.Build("/ch/01/config/name", "s", osc.Arg{Tag: 's', S: "Kick"}))
	r = b.sentToReaper()
	if len(r) != 1 || r[0].Addr != "/track/1/name" || r[0].Args[0].S != "Kick" {
		t.Fatalf("REAPER got %v", addrs(r))
	}
	// one DCA drives its whole REAPER track range
	b.pending = nil
	b.fromX32(osc.Build("/dca/1/fader", "f", osc.Arg{Tag: 'f', F: 0.4}))
	r = b.sentToReaper()
	if len(r) != 3 || r[0].Addr != "/track/10/volume" || r[2].Addr != "/track/12/volume" {
		t.Fatalf("REAPER got %v", addrs(r))
	}
}

func TestBankCTransport(t *testing.T) {
	b := newTestBridge(t, nil)
	// a button fires on the release (value 0), not on the press
	b.fromX32(userParValue(18, 127))
	if r := b.sentToReaper(); len(r) != 0 {
		t.Fatalf("the press must be ignored, got %v", addrs(r))
	}
	b.pending = nil
	b.fromX32(userParValue(18, 0))
	r := b.sentToReaper()
	if len(r) != 1 || r[0].Addr != "/play" || r[0].Args[0].F != 1 {
		t.Fatalf("REAPER got %v", addrs(r))
	}
	for _, tc := range []struct {
		idx  int
		addr string
	}{
		{17, "/action/40042"}, // home
		{19, "/pause"},
		{20, "/action/40043"}, // end
		{22, "/repeat"},
		{23, "/stop"},
		{24, "/record"},
	} {
		b.pending = nil
		b.fromX32(userParValue(tc.idx, 0))
		r = b.sentToReaper()
		if len(r) != 1 || r[0].Addr != tc.addr {
			t.Fatalf("button %d -> %v, want %s", tc.idx, addrs(r), tc.addr)
		}
	}
	// the loop button toggles between the two REAPER loop actions
	b.pending = nil
	b.fromX32(userParValue(21, 0))
	r = b.sentToReaper()
	if len(r) != 1 || r[0].Addr != "/action/40222" {
		t.Fatalf("first loop press -> %v", addrs(r))
	}
	b.pending = nil
	b.fromX32(userParValue(21, 0))
	if r = b.sentToReaper(); len(r) != 1 || r[0].Addr != "/action/40223" {
		t.Fatalf("second loop press -> %v", addrs(r))
	}
	// encoders scrub and re-center themselves
	b.pending = nil
	b.fromX32(userParValue(35, 100))
	r = b.sentToReaper()
	x := b.sentToX32()
	if len(r) != 1 || r[0].Addr != "/action/40173" {
		t.Fatalf("marker encoder forward -> %v", addrs(r))
	}
	if len(x) != 1 || x[0].Addr != "/-stat/userpar/35/value" || x[0].Args[0].I != 64 {
		t.Fatalf("encoder was not re-centered: %v", addrs(x))
	}
}

func TestTransportFeedback(t *testing.T) {
	b := newTestBridge(t, nil)
	b.fromReaper(osc.Build("/play", "f", osc.Arg{Tag: 'f', F: 1}))
	x := b.sentToX32()
	if len(x) != 1 || x[0].Addr != "/-stat/userpar/18/value" || x[0].Args[0].I != 127 {
		t.Fatalf("desk got %v", addrs(x))
	}
	b.pending = nil
	b.fromReaper(osc.Build("/stop", "f", osc.Arg{Tag: 'f', F: 1}))
	x = b.sentToX32()
	if len(x) != 3 {
		t.Fatalf("stop must release the transport buttons, got %v", addrs(x))
	}
}

func TestFiltersBlockAClass(t *testing.T) {
	b := newTestBridge(t, func(c *Config) { c.ToReaper &^= BitFader })
	b.fromX32(osc.Build("/ch/01/mix/fader", "f", osc.Arg{Tag: 'f', F: 0.5}))
	if r := b.sentToReaper(); len(r) != 0 {
		t.Fatalf("fader must be filtered, got %v", addrs(r))
	}
	// other classes still pass
	b.fromX32(osc.Build("/ch/01/mix/pan", "f", osc.Arg{Tag: 'f', F: 0.5}))
	if r := b.sentToReaper(); len(r) != 1 {
		t.Fatalf("pan must pass, got %v", addrs(r))
	}
	b2 := newTestBridge(t, func(c *Config) { c.ToX32 &^= BitName })
	b2.fromReaper(osc.Build("/track/1/name", "s", osc.Arg{Tag: 's', S: "x"}))
	if x := b2.sentToX32(); len(x) != 0 {
		t.Fatalf("name must be filtered, got %v", addrs(x))
	}
}

func TestBundleIsSplit(t *testing.T) {
	b := newTestBridge(t, nil)
	sub1 := osc.Build("/track/1/volume", "f", osc.Arg{Tag: 'f', F: 0.1})
	sub2 := osc.Build("/track/2/volume", "f", osc.Arg{Tag: 'f', F: 0.2})
	bundle := append([]byte("#bundle\x00"), make([]byte, 8)...)
	for _, sub := range [][]byte{sub1, sub2} {
		bundle = append(bundle, byte(len(sub)>>24), byte(len(sub)>>16), byte(len(sub)>>8), byte(len(sub)))
		bundle = append(bundle, sub...)
	}
	b.fromReaper(bundle)
	x := b.sentToX32()
	if len(x) != 2 || x[0].Args[0].F != 0.1 || x[1].Args[0].F != 0.2 {
		t.Fatalf("bundle was not split: %v", addrs(x))
	}
}

// TestBankStepRestoresTheDesk checks the bank logic: moving the bank re-sends
// every strip of the new bank to the console and asks REAPER for a refresh.
func TestBankStepRestoresTheDesk(t *testing.T) {
	b := newTestBridge(t, func(c *Config) {
		c.BankSize = 8
		c.TrkMin, c.TrkMax = 1, 24
		c.ChBank = true
		c.Transport = false
	})
	// the desk has seen track 9 (strip 1 of bank 1)
	b.fromReaper(osc.Build("/track/9/volume", "f", osc.Arg{Tag: 'f', F: 0.3}))
	if _, ok := b.cfg.TrackToX32(9); ok {
		t.Fatal("track 9 must not be visible in bank 0")
	}
	b.pending = nil
	b.fromX32(userParValue(12+b.cfg.BankUp, 0)) // Bank C "up" button
	if b.cfg.BankOffset != 1 {
		t.Fatalf("bank offset = %d, want 1", b.cfg.BankOffset)
	}
	x := b.sentToX32()
	if !hasAddr(x, "/ch/01/mix/fader") || !hasAddr(x, "/ch/08/config/name") {
		t.Fatalf("the new bank was not restored: %v", addrs(x))
	}
	r := b.sentToReaper()
	if !hasAddr(r, "/action/41743") {
		t.Fatalf("REAPER was not asked to refresh: %v", addrs(r))
	}
	// the cached fader of track 9 lands on strip 1
	for _, m := range x {
		if m.Addr == "/ch/01/mix/fader" && m.Args[0].F != 0.3 {
			t.Fatalf("cached fader = %v, want 0.3", m.Args[0].F)
		}
	}
	// and back down
	b.pending = nil
	b.fromX32(userParValue(12+b.cfg.BankDn, 0))
	if b.cfg.BankOffset != 0 {
		t.Fatalf("bank offset = %d, want 0", b.cfg.BankOffset)
	}
}
