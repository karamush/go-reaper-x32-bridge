package x32

import (
	"net"
	"time"
)

// Server owns the UDP socket and the main loop of the emulator (the select()
// loop of X32.c: non blocking receive with a 10ms timeout, meter frames sent on
// every iteration).
type Server struct {
	Conn  *net.UDPConn
	State *State
}

// NewServer binds the UDP socket on ip:Port and wires the send function.
func NewServer(ip string, st *State) (*Server, error) {
	addr, err := net.ResolveUDPAddr("udp4", net.JoinHostPort(ip, st.Port))
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		return nil, err
	}
	st.SendFn = func(b []byte, to *net.UDPAddr) error {
		_, err := conn.WriteToUDP(b, to)
		return err
	}
	return &Server{Conn: conn, State: st}, nil
}

// Run reads and handles packets until the state asks to stop.
func (srv *Server) Run() error {
	buf := make([]byte, BSize)
	for !srv.State.Stop {
		_ = srv.Conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
		n, addr, err := srv.Conn.ReadFromUDP(buf)
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				srv.State.PumpMeters()
				continue
			}
			return err
		}
		pkt := make([]byte, n)
		copy(pkt, buf[:n])
		srv.State.HandlePacket(pkt, addr)
		srv.State.PumpMeters()
	}
	return nil
}

// Close releases the socket.
func (srv *Server) Close() error { return srv.Conn.Close() }
