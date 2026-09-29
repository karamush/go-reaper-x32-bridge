package x32

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"x32emu/osc"
)

// ---------------------------------------------------------------- /-stat/userpar
//
// The user assign controls are reported by a real console (and used by bridge
// software such as X32ReaperW) through "/-stat/userpar/NN/value". The table is
// hand written, so its contract is checked here: address length, packet layout
// and notification behaviour.

func userParSet(idx int, v int32) []byte {
	return osc.AppendInt32(osc.Start(userParAddr(idx), "i"), v)
}

func TestUserParIndex(t *testing.T) {
	// X32ReaperW uses "/-stat/userpar/%2d/value" with (12 + button number).
	if got := UserParIndex(5); got != 17 {
		t.Fatalf("UserParIndex(5) = %d, want 17", got)
	}
	if got := UserParIndex(12); got != 24 {
		t.Fatalf("UserParIndex(12) = %d, want 24", got)
	}
}

// The address must stay 23 characters long: bridge clients parse the fixed byte
// offsets of the notification (tag at 24, integer at 28).
func TestUserParAddressLayout(t *testing.T) {
	if n := len(userParAddr(18)); n != 23 {
		t.Fatalf("address length = %d, want 23", n)
	}
	b := userParSet(18, 127)
	if len(b) != 32 {
		t.Fatalf("packet length = %d, want 32", len(b))
	}
	if b[23] != 0 || b[24] != ',' || b[25] != 'i' {
		t.Fatalf("tag block at the wrong offset: %q", b[20:28])
	}
	if v := int32(binary.BigEndian.Uint32(b[28:])); v != 127 {
		t.Fatalf("integer at offset 28 = %d, want 127", v)
	}
}

func TestUserParGet(t *testing.T) {
	st, sends := newTestState(t)
	a := udpAddr(t, "192.168.0.20:1000")
	st.HandlePacket(rawPacket(userParAddr(17)), a)
	m := osc.Decode((*sends)[0].b)
	if m.Addr != userParAddr(17) || m.Tags != "i" || m.Args[0].I != 0 {
		t.Fatalf("button reply = %+v", m)
	}
	// encoders start centered, like on the desk
	st.HandlePacket(rawPacket(userParAddr(33)), a)
	m = osc.Decode((*sends)[1].b)
	if m.Addr != userParAddr(33) || m.Tags != "i" || m.Args[0].I != UserParCenter {
		t.Fatalf("encoder reply = %+v", m)
	}
}

// A SET from a third party (a GUI, a second bridge) reaches the subscribed
// client as the echo of the request - exactly what X32ReaperW listens to.
func TestUserParSetNotifiesSubscribedClients(t *testing.T) {
	st, sends := newTestState(t)
	a, b := udpAddr(t, "192.168.0.20:1000"), udpAddr(t, "192.168.0.21:1001")
	st.HandlePacket(rawPacket("/xremote"), b)
	*sends = (*sends)[:0]

	// button down (127) then button up (0): X32ReaperW acts on the release
	down := userParSet(18, 127)
	st.HandlePacket(down, a)
	if len(*sends) != 1 || (*sends)[0].addr != b.String() {
		t.Fatalf("expected one notification to %s, got %+v", b, *sends)
	}
	if string((*sends)[0].b) != string(down) {
		t.Fatalf("notification is not the echo of the request")
	}
	*sends = (*sends)[:0]
	st.HandlePacket(userParSet(18, 0), a)
	if len(*sends) != 1 {
		t.Fatalf("release must notify once, got %d", len(*sends))
	}
	if v := int32(binary.BigEndian.Uint32((*sends)[0].b[28:])); v != 0 {
		t.Fatalf("release value = %d, want 0", v)
	}
	// no change: no answer at all
	*sends = (*sends)[:0]
	st.HandlePacket(userParSet(18, 0), a)
	if len(*sends) != 0 {
		t.Fatalf("an unchanged value must stay silent, got %d messages", len(*sends))
	}
}

// PressUserPar is the console side event: it reaches every xremote client,
// including the one that sent the previous request.
func TestPressUserParBroadcastsToAll(t *testing.T) {
	st, sends := newTestState(t)
	a, b := udpAddr(t, "192.168.0.20:1000"), udpAddr(t, "192.168.0.21:1001")
	st.HandlePacket(rawPacket("/xremote"), a)
	st.HandlePacket(rawPacket("/xremote"), b)
	*sends = (*sends)[:0]
	// a is the "last sender", it must still be notified
	st.From = a

	if !st.PressUserPar(18, 127) {
		t.Fatal("PressUserPar(18) failed")
	}
	if len(*sends) != 2 {
		t.Fatalf("expected two notifications, got %d", len(*sends))
	}
	for _, s := range *sends {
		m := osc.Decode(s.b)
		if m.Addr != userParAddr(18) || m.Tags != "i" || m.Args[0].I != 127 {
			t.Fatalf("notification = %+v", m)
		}
	}
	if v, ok := UserPar(18); !ok || v != 127 {
		t.Fatalf("UserPar(18) = %d, %v", v, ok)
	}
	// an unchanged value produces no event, and unknown indices are rejected
	*sends = (*sends)[:0]
	st.PressUserPar(18, 127)
	if len(*sends) != 0 {
		t.Fatalf("unchanged value must not be reported, got %d", len(*sends))
	}
	if st.PressUserPar(19, 0) == false {
		t.Fatal("PressUserPar(19) should be accepted")
	}
	if st.PressUserPar(1, 0) || st.PressUserPar(99, 0) {
		t.Fatal("indices outside the Bank C range must be rejected")
	}
}

// The userpar table is not persisted: ".X32res.rc" keeps the exact layout of the
// C emulator, which knows nothing about these parameters.
func TestUserParIsNotPersisted(t *testing.T) {
	for _, tab := range persistTables() {
		for i := range tab {
			if strings.HasPrefix(tab[i].Addr, UserParAddrPrefix) {
				t.Fatalf("%s must not be persisted", tab[i].Addr)
			}
		}
	}
	st, _ := newTestState(t)
	st.PressUserPar(18, 127)
	path := filepath.Join(t.TempDir(), ".X32res.rc")
	if err := st.SaveResource(path); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "userpar") {
		t.Fatal("resource file must not mention userpar")
	}
}
