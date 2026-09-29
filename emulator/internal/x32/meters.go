package x32

import (
	"encoding/binary"
	"net"
	"time"

	"x32emu/osc"
)

// meter is one active meter subscription (the XActiveMeters/XClientMeters/
// XTimerMeters/XInterMeters/XDeltaMeters arrays of X32.c).
type meter struct {
	active bool
	stop   time.Time     // XTimerMeters: meters stay alive 10s after a request
	next   time.Time     // XInterMeters: time of the next frame
	delta  time.Duration // XDeltaMeters: 50ms * time factor
	addr   *net.UDPAddr  // XClientMeters
	buf    []byte        // Xbuf_meters / Lbuf_meters
	data   int           // offset of the value area inside buf
}

// meterSizes is the number of values carried by "/meters/<i>" (Xprepmeter
// calls in function_meters() of X32.c).
var meterSizes = [MaxMeters]int{70, 96, 49, 22, 82, 27, 4, 16, 6, 32, 32, 5, 4, 48, 80, 50, 48}

// functionMeters handles /meters: it subscribes (or refreshes) the requesting
// client for one meter group. The frames themselves are sent by PumpMeters.
//
// The emulator has no audio: like X32.c, the values are all zero.
func (s *State) functionMeters() int {
	addr := s.argS(0)
	for i := range Xmeters {
		if Xmeters[i].Addr != addr {
			continue
		}
		// the C code locates the time factor at r_buf[k + 24] where k depends on
		// the meter group (0.88 fix for /meters/5 and /meters/6)
		k := 0
		switch i {
		case 5:
			k = 12
		case 6:
			k = 8
		}
		tf := int32(0)
		off := k + 24
		if off+4 <= len(s.Raw) {
			tf = int32(binary.BigEndian.Uint32(s.Raw[off:]))
		}
		if tf < 1 || tf > 99 {
			tf = 1
		}
		s.prepMeter(i, meterSizes[i], tf)
		return 0
	}
	return 0
}

// prepMeter builds the frame of meter group i and (re)activates it for 10
// seconds (Xprepmeter in X32.c).
func (s *State) prepMeter(i, l int, tf int32) {
	head := osc.Start("/meters/"+itoa(i), "b")
	buf := make([]byte, len(head)+4+4+4*l)
	copy(buf, head)
	binary.BigEndian.PutUint32(buf[len(head):], uint32((l+1)*4)) // blob size
	binary.LittleEndian.PutUint32(buf[len(head)+4:], uint32(l))  // value count (little endian!)
	m := &s.meters[i]
	m.buf = buf
	m.active = true
	now := time.Now()
	m.stop = now.Add(10 * time.Second)
	m.next = now
	m.addr = s.From
	m.delta = 50 * time.Millisecond * time.Duration(tf)
	// the value area follows the address+tags, the blob size and the count
	m.data = len(head) + 8
}

// PumpMeters sends every pending meter frame. It must be called regularly
// (the main loop of X32.c calls it on every select() timeout).
func (s *State) PumpMeters() {
	now := time.Now()
	for i := range s.meters {
		m := &s.meters[i]
		if !m.active {
			continue
		}
		if now.After(m.stop) {
			m.active = false
			continue
		}
		if now.Before(m.next) {
			continue
		}
		s.fillMeterFrame(i, m)
		if m.addr != nil {
			_ = s.Send(m.buf, m.addr)
		}
		m.next = m.next.Add(m.delta)
		if m.next.Before(now) {
			m.next = now.Add(m.delta)
		}
	}
}

func itoa(i int) string {
	if i < 10 {
		return string(rune('0' + i))
	}
	return string(rune('0'+i/10)) + string(rune('0'+i%10))
}
