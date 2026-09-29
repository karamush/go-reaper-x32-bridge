package bridge

import (
	"fmt"
	"net"
	"time"

	"x32emu/osc"
)

// Keepalive and timing constants. X32ReaperW re-registers its /xremote
// subscription every 9 s because a console drops a silent client after 11 s; the
// read loop uses the same 1 ms tick as the reference tool.
const (
	XRemoteRepeat = 9 * time.Second
	readTick      = time.Millisecond
)

// Bridge is the running two way bridge.
type Bridge struct {
	cfg    Config
	state  *stateCache
	meters MeterSink // filled by SetMeters (nil fields = no meter feed)

	x32Conn  *net.UDPConn // console side: unbound unless LocalPort is set
	reapConn *net.UDPConn // bound to cfg.Listen: REAPER sends here
	x32Addr  *net.UDPAddr
	reapAddr *net.UDPAddr

	pending []outMsg // datagrams waiting for their (optional) delay

	connected   bool // at least one answer of the console was received
	loopOn      bool // Bank C loop button state
	stop        bool // Stop() has been called
	verbose     bool
	log         func(format string, args ...any)
	onBankEvent func(offset int) // hook used by tests

	// masterMeter is the last level REAPER reported for the main bus (L and R).
	masterMeter [2]float32
}

// MeterSink receives the levels REAPER reports, already resolved to the console
// element they belong to: section is "ch" with a number of 1..32, "auxin" (1..8),
// "fxrtn" (1..8) or "bus" (1..16). Master receives the two channels of the main
// bus. Either field may be nil, which disables that half of the feed.
//
// The console itself takes care of building the frames: the emulator of this
// workspace fills its /meters/NN frames with these levels (see
// docs/X32_REAPER_BRIDGE.md and docs/X32_OSC_PROTOCOL.md §7.6).
type MeterSink struct {
	Track  func(section string, number int, v float32)
	Master func(left, right float32)
}

// SetMeters installs the sink used for the meters of the console. Call it before
// Run.
func (b *Bridge) SetMeters(s MeterSink) { b.meters = s }

// outMsg is one datagram scheduled for sending.
type outMsg struct {
	b        []byte
	toReaper bool
	bank     bool // use DelayBank instead of DelayGen
	at       time.Time
}

// New prepares a bridge. Run() opens the sockets.
func New(cfg Config) (*Bridge, error) {
	if err := cfg.Check(); err != nil {
		return nil, err
	}
	b := &Bridge{cfg: cfg, state: newStateCache(), verbose: cfg.Verbose}
	b.log = func(format string, args ...any) {
		if b.verbose {
			fmt.Printf(format, args...)
		}
	}
	return b, nil
}

// Config returns a copy of the configuration in use.
func (b *Bridge) Config() Config { return b.cfg }

// DeviceTrackCount returns the number of tracks the bridge asks REAPER to report
// (every track that any section of the desk can reach).
func (b *Bridge) DeviceTrackCount() int {
	max := b.cfg.TrkMax
	for _, v := range []int{b.cfg.AuxMax, b.cfg.FxrMax, b.cfg.BusMax} {
		if v > max {
			max = v
		}
	}
	for i := 0; i < 8; i++ {
		if b.cfg.RdcaMax[i] > max {
			max = b.cfg.RdcaMax[i]
		}
	}
	if max < b.cfg.DeviceTracks {
		max = b.cfg.DeviceTracks
	}
	if max < 1 {
		max = 1
	}
	if max > 128 {
		max = 128
	}
	return max
}

// Run opens the sockets, announces itself and loops until Stop is called.
func (b *Bridge) Run() error {
	var err error
	if b.x32Addr, err = net.ResolveUDPAddr("udp4", b.cfg.X32Addr); err != nil {
		return fmt.Errorf("X32 address %q: %w", b.cfg.X32Addr, err)
	}
	// The console side socket is deliberately left unbound when LocalPort is 0:
	// a console answers the source port of the request, and an ephemeral port is
	// what X32ReaperW uses too.
	local := &net.UDPAddr{Port: b.cfg.LocalPort}
	if b.cfg.LocalPort > 0 {
		if local, err = net.ResolveUDPAddr("udp4", "0.0.0.0:"+fmt.Sprint(b.cfg.LocalPort)); err != nil {
			return fmt.Errorf("local port %d: %w", b.cfg.LocalPort, err)
		}
	}
	if b.x32Conn, err = net.ListenUDP("udp4", local); err != nil {
		return fmt.Errorf("console socket: %w", err)
	}
	defer b.x32Conn.Close()
	if b.reapAddr, err = net.ResolveUDPAddr("udp4", net.JoinHostPort(b.cfg.Host, b.cfg.Port)); err != nil {
		return fmt.Errorf("REAPER address: %w", err)
	}
	if b.reapConn, err = net.ListenUDP("udp4", mustResolve(b.cfg.Listen)); err != nil {
		return fmt.Errorf("listen %s: %w", b.cfg.Listen, err)
	}
	defer b.reapConn.Close()

	b.log("X32 -> %s (local %s), REAPER <- %s, sending to %s\n",
		b.cfg.X32Addr, b.x32Conn.LocalAddr(), b.cfg.Listen, b.reapAddr)
	b.handshake()
	return b.loop()
}

func mustResolve(s string) *net.UDPAddr {
	a, err := net.ResolveUDPAddr("udp4", s)
	if err != nil {
		// Run() reports the error of ListenUDP instead
		return &net.UDPAddr{}
	}
	return a
}

// handshake announces the bridge to both sides. X32ReaperW probes the console
// with "/info" (a console answers exactly "/info"), subscribes with "/xremote"
// and reads the selected channel; on the REAPER side it declares how many tracks
// and sends the surface has, then asks for a full refresh.
func (b *Bridge) handshake() {
	b.toX32(osc.Start("/info", ""), 0)
	b.toX32(osc.Start("/xremote", ""), 0)
	b.toX32(osc.AppendString(nil, "/-stat/selidx"), 0)
	b.setupBankC()
	if b.cfg.DeviceTracks > 0 {
		b.toReaper(osc.Build("/device/track/count", "i", osc.Arg{Tag: 'i', I: int32(b.DeviceTrackCount())}), 0)
	}
	if b.cfg.DeviceSends > 0 {
		b.toReaper(osc.Build("/device/send/count", "i", osc.Arg{Tag: 'i', I: int32(b.cfg.DeviceSends)}), 0)
	}
	// "Control surface: refresh" - makes REAPER report the state of every track it
	// knows about, which is what fills the desk after a start or a bank change.
	b.refreshReaper()
}

// refreshReaper asks REAPER to re-send the state of its tracks.
func (b *Bridge) refreshReaper() {
	b.toReaper(osc.Build("/device/track/count", "i", osc.Arg{Tag: 'i', I: int32(b.DeviceTrackCount())}), 0)
	b.toReaper(actionMsg(41743), 0)
}

// setupBankC assigns the Bank C controls of the desk and initialises them, like
// X32UsrCtrlC() does: without it a console (or the emulator) would have no
// buttons to report.
func (b *Bridge) setupBankC() {
	if !b.cfg.Transport && !b.cfg.ChBank {
		return
	}
	names := [4]string{"MP13000", "MP14000", "MP15000", "MP16000"}
	for i, n := range names {
		b.toX32(osc.Build("/config/userctrl/C/enc/"+fmt.Sprint(i+1), "s", osc.Arg{Tag: 's', S: n}), 0)
	}
	buttons := [8]string{"MN16000", "MN16001", "MN16002", "MN16003", "MN16004", "MN16005", "MN16006", "MN16007"}
	for i, n := range buttons {
		b.toX32(osc.Build("/config/userctrl/C/btn/"+fmt.Sprint(i+5), "s", osc.Arg{Tag: 's', S: n}), 0)
	}
	for i := 33; i <= 36; i++ {
		b.toX32(userParValue(i, 64), 0)
	}
	for i := 17; i <= 24; i++ {
		b.toX32(userParValue(i, 0), 0)
	}
	b.toX32(osc.Build("/config/userctrl/C/color", "i", osc.Arg{Tag: 'i', I: 0}), 0)
	b.toX32(osc.Build("/-stat/userbank", "i", osc.Arg{Tag: 'i', I: 2}), 0)
}

func userParValue(idx int, v int32) []byte {
	return osc.Build("/-stat/userpar/"+fmt.Sprint(idx)+"/value", "i", osc.Arg{Tag: 'i', I: v})
}

// Stop asks the loop to return (the sockets are closed by Run).
func (b *Bridge) Stop() { b.stop = true }

// loop is the main select()/read loop of the bridge: it delivers the queued
// datagrams, keeps the /xremote subscription alive and reads both sockets.
func (b *Bridge) loop() error {
	buf := make([]byte, 4096)
	lastXRemote := time.Now().Add(-XRemoteRepeat + 2*time.Second)
	for !b.stop {
		b.flush()
		if time.Since(lastXRemote) >= XRemoteRepeat {
			b.toX32(osc.Start("/xremote", ""), 0)
			lastXRemote = time.Now()
		}
		_ = b.reapConn.SetReadDeadline(time.Now().Add(readTick))
		if n, _, err := b.reapConn.ReadFromUDP(buf); err == nil {
			pkt := append([]byte(nil), buf[:n]...)
			b.logDatagram("R->", pkt)
			b.fromReaper(pkt)
		}
		_ = b.x32Conn.SetReadDeadline(time.Now().Add(readTick))
		if n, _, err := b.x32Conn.ReadFromUDP(buf); err == nil {
			pkt := append([]byte(nil), buf[:n]...)
			b.logDatagram("X->", pkt)
			b.fromX32(pkt)
		}
	}
	return nil
}

// logDatagram logs one incoming datagram. Bundles are split: REAPER sends all of
// its feedback in bundles, and dumping only the "#bundle" header would hide what
// the bridge actually received.
func (b *Bridge) logDatagram(prefix string, pkt []byte) {
	if !b.verbose {
		return
	}
	for _, sub := range splitBundle(pkt) {
		b.log("  %s %s\n", prefix, osc.Dump("<", sub, false))
	}
}

// flush sends every queued datagram whose delay has elapsed.
func (b *Bridge) flush() {
	if len(b.pending) == 0 {
		return
	}
	now := time.Now()
	keep := b.pending[:0]
	for _, m := range b.pending {
		if m.at.After(now) {
			keep = append(keep, m)
			continue
		}
		var err error
		if m.toReaper {
			_, err = b.reapConn.WriteToUDP(m.b, b.reapAddr)
		} else {
			_, err = b.x32Conn.WriteToUDP(m.b, b.x32Addr)
		}
		if err != nil {
			b.log("send error: %v\n", err)
		}
	}
	b.pending = keep
}

// toX32 queues a datagram for the console.
func (b *Bridge) toX32(p []byte, delayMs int) {
	if len(p) == 0 {
		return
	}
	b.log("  ->X %s\n", osc.Dump(">", p, false))
	b.pending = append(b.pending, outMsg{b: p, at: time.Now().Add(time.Duration(delayMs) * time.Millisecond)})
}

// toReaper queues a datagram for REAPER.
func (b *Bridge) toReaper(p []byte, delayMs int) {
	if len(p) == 0 {
		return
	}
	b.log("  ->R %s\n", osc.Dump(">", p, false))
	b.pending = append(b.pending,
		outMsg{b: p, toReaper: true, at: time.Now().Add(time.Duration(delayMs) * time.Millisecond)})
}

// fromX32 processes one datagram of the console.
func (b *Bridge) fromX32(pkt []byte) {
	b.connected = true
	m := osc.Decode(pkt)
	for _, out := range b.x32ToReaper(m) {
		b.toReaper(out, 0)
	}
}

// fromReaper processes one datagram of REAPER (which may be a bundle).
func (b *Bridge) fromReaper(pkt []byte) {
	for _, sub := range splitBundle(pkt) {
		m := osc.Decode(sub)
		if m.Addr == "" {
			continue
		}
		for _, out := range b.reaperToX32(m) {
			b.toX32(out, 0)
		}
	}
}

// bankStep moves the channel bank of the desk and restores it on the console.
// Everything it sends is queued directly (the console messages and the REAPER
// refresh), so it always returns nil.
func (b *Bridge) bankStep(delta int) [][]byte {
	off := b.cfg.BankOffset + delta
	max := b.cfg.MaxBankOffset()
	if off < 0 {
		off = 0
	}
	if off > max {
		off = max
	}
	if off == b.cfg.BankOffset {
		return nil
	}
	b.cfg.BankOffset = off
	b.log("channel bank -> %d (tracks %d..%d)\n",
		off, b.cfg.TrkMin+off*b.cfg.BankSize, b.cfg.TrkMin+(off+1)*b.cfg.BankSize-1)
	b.recallBank()
	if b.onBankEvent != nil {
		b.onBankEvent(off)
	}
	return nil
}

// recallBank pushes the state of the new bank to the desk and asks REAPER for a
// full refresh: a console has no idea of the tracks it is not showing, and
// REAPER reports them on change only, so the bridge restores what it knows and
// forces REAPER to send the rest.
func (b *Bridge) recallBank() {
	for strip := 1; strip <= b.cfg.BankSize; strip++ {
		track, ok := b.cfg.X32ToTrack(strip)
		if !ok {
			continue
		}
		st := b.state.peek(track)
		fader, pan, on, name, solo := float32(0), float32(0.5), int32(1), "", int32(0)
		if st != nil {
			fader, pan, name = st.fader, st.pan, st.name
			if st.mute {
				on = 0
			}
			if st.solo {
				solo = 1
			}
		}
		if b.cfg.ToX32&BitFader != 0 {
			b.toX32(x32Float("ch", strip, "mix/fader", fader), b.cfg.DelayBank)
		}
		if b.cfg.ToX32&BitPan != 0 {
			b.toX32(x32Float("ch", strip, "mix/pan", pan), b.cfg.DelayBank)
		}
		if b.cfg.ToX32&BitMute != 0 {
			b.toX32(x32Int("ch", strip, "mix/on", on), b.cfg.DelayBank)
		}
		if b.cfg.ToX32&BitName != 0 {
			b.toX32(x32String("ch", strip, "config/name", name), b.cfg.DelayBank)
		}
		if b.cfg.ToX32&BitSolo != 0 {
			b.toX32(soloMsg(strip, solo), b.cfg.DelayBank)
		}
	}
	b.feedBankMeters()
	b.refreshReaper()
}
