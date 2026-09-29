// Package bridge implements a two way OSC bridge between an X32 console (or the
// emulator of the emulator module) and REAPER. Its model is X32ReaperW v2.86 by
// Patrick-Gilles Maillot (upstream X32-Behringer project,
// https://github.com/pmaillot/X32-Behringer, see docs/X32_REAPER_BRIDGE.md):
// the desk acts as a control surface for REAPER, REAPER moves the desk, and a
// "bank" of 8/16/32 channels follows the DAW track bank.
//
// This file holds the configuration and the address/range mapping rules.
package bridge

import "fmt"

// Filter bits (Xrsend / Xxsend masks of X32ReaperW, §5). The same ten bits are
// used in both directions and mean the same parameter classes.
const (
	BitPan    uint16 = 0x0001
	BitFader  uint16 = 0x0002
	BitName   uint16 = 0x0004
	BitMute   uint16 = 0x0008
	BitSelect uint16 = 0x0010
	BitSend   uint16 = 0x0020
	BitSolo   uint16 = 0x0040
	BitFX     uint16 = 0x0080
	BitMPan   uint16 = 0x0100
	BitMFader uint16 = 0x0200

	// AllBits enables every class (the default, "-1" in X32ReaperW).
	AllBits = BitPan | BitFader | BitName | BitMute | BitSelect |
		BitSend | BitSolo | BitFX | BitMPan | BitMFader
)

// X32 strip ranges (see §3.3 of docs/X32_OSC_PROTOCOL.md): the console numbers
// every fader of the desk from 1 to 80.
const (
	StripAuxinFirst = 33
	StripAuxinLast  = 40
	StripFxrtnFirst = 41
	StripFxrtnLast  = 48
	StripBusFirst   = 49
	StripBusLast    = 64
	StripDcaFirst   = 73
	StripDcaLast    = 80

	MaxSends = 16 // /ch/NN/mix/01..16/level
)

// Config holds every tuning knob of the bridge. The zero value is not usable:
// use DefaultConfig and override what is needed.
type Config struct {
	// Sockets. X32Addr is the console/emulator to talk to; Listen is our local
	// socket (REAPER sends here); Host/Port is where REAPER listens (we send
	// there). LocalPort optionally pins the local port of the console side
	// socket (0 = let the OS choose one, like X32ReaperW does).
	X32Addr   string
	Listen    string
	Host      string
	Port      string
	LocalPort int

	// Bank and track ranges (the "preset" values of X32ReaperW, §5).
	BankSize   int
	BankOffset int
	SendOffset int // TrackSendOffset

	TrkMin, TrkMax int
	AuxMin, AuxMax int
	FxrMin, FxrMax int
	BusMin, BusMax int
	DcaMin, DcaMax int
	RdcaMin        [8]int // REAPER track range driven by X32 DCA 1..8
	RdcaMax        [8]int

	// Modes.
	Transport bool // Bank C buttons/encoders drive the REAPER transport
	ChBank    bool // Bank C drives the channel bank instead of the transport
	Master    bool // /main/st is mirrored to /master
	EqCmp     bool // reserved (EQ/DYN mirroring, phase B)

	// Bank C assignment when ChBank is set (button numbers 5..12, as in
	// X32ReaperW: /-stat/userpar index = 12 + button).
	BankUp       int
	BankDn       int
	MarkerButton int

	// Filtering: which classes are allowed to travel in which direction.
	ToX32    uint16 // REAPER -> console
	ToReaper uint16 // console -> REAPER

	Verbose   bool
	DelayBank int // ms between the messages of a bank recall (Xdelayb)
	DelayGen  int // ms between ordinary messages (Xdelayg)

	// Device counts announced to REAPER at startup so that its OSC feedback is
	// not limited to the eight tracks of the standard .ReaperOSC template.
	DeviceTracks int
	DeviceSends  int
}

// DefaultConfig returns the settings used when no flag is given: the emulator on
// the loopback interface, REAPER listening on 8000 and sending to 9000 (the
// values of REAPER's Control/OSC/web page), a bank of 32 channels mapped on
// REAPER tracks 1..32 and every filter open.
func DefaultConfig() Config {
	return Config{
		X32Addr: "127.0.0.1:10023",
		Listen:  "0.0.0.0:9000",
		Host:    "127.0.0.1",
		Port:    "8000",

		BankSize:   32,
		BankOffset: 0,
		SendOffset: 0,

		TrkMin: 1, TrkMax: 32,

		Transport: true,
		ChBank:    false,
		Master:    true,

		BankUp:       9,
		BankDn:       10,
		MarkerButton: 7,

		ToX32:    AllBits,
		ToReaper: AllBits,

		// The logs are off by default: -v turns them on.
		Verbose:   false,
		DelayBank: 0,
		DelayGen:  0,

		DeviceTracks: 32,
		DeviceSends:  MaxSends,
	}
}

// Check validates the ranges and fills in the defaults of the optional fields.
func (c *Config) Check() error {
	if c.BankSize < 1 || c.BankSize > 32 {
		return fmt.Errorf("bank size %d out of range (1..32)", c.BankSize)
	}
	if c.BankOffset < 0 {
		return fmt.Errorf("bank offset %d must not be negative", c.BankOffset)
	}
	if c.SendOffset < 0 || c.SendOffset > MaxSends {
		return fmt.Errorf("send offset %d out of range (0..%d)", c.SendOffset, MaxSends)
	}
	if c.TrkMax > 0 && c.TrkMin < 1 {
		return fmt.Errorf("track range %d..%d is not usable", c.TrkMin, c.TrkMax)
	}
	if c.TrkMax > 0 && c.TrkMax < c.TrkMin {
		return fmt.Errorf("track range %d..%d is inverted", c.TrkMin, c.TrkMax)
	}
	if c.Listen == "" {
		return fmt.Errorf("listen address must not be empty")
	}
	if c.X32Addr == "" {
		return fmt.Errorf("X32 address must not be empty")
	}
	if c.Host == "" || c.Port == "" {
		return fmt.Errorf("REAPER destination must be host:port")
	}
	if c.DeviceTracks == 0 {
		c.DeviceTracks = c.BankSize
	}
	if c.DeviceSends == 0 {
		c.DeviceSends = MaxSends
	}
	return nil
}

// X32ToTrack converts a console strip number into a REAPER track number. ok is
// false when the strip has no counterpart (or the corresponding section is
// disabled). Channel strips are resolved through the bank: strip s of bank o
// holds REAPER track TrkMin + o*BankSize + s - 1. The auxin, fxrtn and bus
// sections are mapped one to one on their own track ranges.
func (c *Config) X32ToTrack(strip int) (int, bool) {
	switch {
	case strip >= 1 && strip <= c.BankSize:
		if c.TrkMax <= 0 {
			return 0, false
		}
		track := c.TrkMin + c.BankOffset*c.BankSize + strip - 1
		return track, track <= c.TrkMax
	case strip >= StripAuxinFirst && strip <= StripAuxinLast:
		if c.AuxMax <= 0 {
			return 0, false
		}
		return c.AuxMin + strip - StripAuxinFirst, true
	case strip >= StripFxrtnFirst && strip <= StripFxrtnLast:
		if c.FxrMax <= 0 {
			return 0, false
		}
		return c.FxrMin + strip - StripFxrtnFirst, true
	case strip >= StripBusFirst && strip <= StripBusLast:
		if c.BusMax <= 0 {
			return 0, false
		}
		return c.BusMin + strip - StripBusFirst, true
	}
	return 0, false
}

// TrackToX32 is the inverse of X32ToTrack. ok is false when the track does not
// belong to any section, or when it falls outside the current channel bank.
func (c *Config) TrackToX32(track int) (int, bool) {
	switch {
	case c.TrkMax > 0 && track >= c.TrkMin && track <= c.TrkMax:
		strip := track - c.TrkMin + 1 - c.BankOffset*c.BankSize
		if strip < 1 || strip > c.BankSize {
			return 0, false
		}
		return strip, true
	case c.AuxMax > 0 && track >= c.AuxMin && track <= c.AuxMax:
		return StripAuxinFirst + track - c.AuxMin, true
	case c.FxrMax > 0 && track >= c.FxrMin && track <= c.FxrMax:
		return StripFxrtnFirst + track - c.FxrMin, true
	case c.BusMax > 0 && track >= c.BusMin && track <= c.BusMax:
		return StripBusFirst + track - c.BusMin, true
	}
	return 0, false
}

// DcaTracks returns the REAPER tracks driven by X32 DCA n (1..8). A DCA whose
// RdcaMin <= RdcaMax covers that whole range (X32ReaperW ver 2.0: "multiple
// REAPER tracks can be controlled from X32 DCA channels").
func (c *Config) DcaTracks(n int) []int {
	if n < 1 || n > 8 {
		return nil
	}
	lo, hi := c.RdcaMin[n-1], c.RdcaMax[n-1]
	if hi <= 0 || hi < lo {
		return nil
	}
	out := make([]int, 0, hi-lo+1)
	for t := lo; t <= hi; t++ {
		out = append(out, t)
	}
	return out
}

// DcaForTrack returns the X32 DCA number (1..8) driving a REAPER track, 0 when
// there is none.
func (c *Config) DcaForTrack(track int) int {
	for i := 0; i < 8; i++ {
		if c.RdcaMax[i] > 0 && track >= c.RdcaMin[i] && track <= c.RdcaMax[i] {
			return i + 1
		}
	}
	return 0
}

// MaxBankOffset returns the highest bank offset that still holds a mapped track.
func (c *Config) MaxBankOffset() int {
	if c.TrkMax <= 0 || c.TrkMax < c.TrkMin {
		return 0
	}
	n := (c.TrkMax - c.TrkMin + 1) / c.BankSize
	if (c.TrkMax-c.TrkMin+1)%c.BankSize != 0 {
		n++
	}
	if n < 1 {
		n = 1
	}
	return n - 1
}
