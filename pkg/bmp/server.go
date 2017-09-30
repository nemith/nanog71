package bmp

import (
	"bufio"
	"net"

	"github.com/osrg/gobgp/packet/bmp"
)

// Message a single BMP message
type Message struct {
	SourceAddr string
	*bmp.BMPMessage
}

// Handler will hand a single BMP message
type Handler interface {
	HandleBMPMessage(*Message)
}

// HandlerFunc is an adapter to use a function as a Handler.
type HandlerFunc func(*Message)

// HandleBMPMessage calls f(msg)
func (f HandlerFunc) HandleBMPMessage(msg *Message) {
	f(msg)
}

// Server defines parameters for a BMP server.  Zero value for Server is valid
type Server struct {
	Addr    string  // TCP address to listen on, :11019 if empty
	Handler Handler // handler to invoke on new messages
}

// ListenAndServer will listen on the TCP network address srv.Adr and then
// handle the incoming connections.
func (s *Server) ListenAndServe() error {
	addr := s.Addr
	if addr != "" {
		addr = ":11019"
	}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go handleConn(conn, s.Handler)
	}
}

func ListenAndServe(addr string) error {
	s := Server{Addr: addr}
	return s.ListenAndServe()
}

func handleConn(conn net.Conn, handler Handler) {
	scanner := bufio.NewScanner(conn)
	scanner.Split(bmp.SplitBMP)

	for scanner.Scan() {
		msg, err := bmp.ParseBMPMessage(scanner.Bytes())
		if err != nil {
			continue
		}
		handler.HandleBMPMessage(&Message{
			SourceAddr: conn.RemoteAddr().String(),
			BMPMessage: msg,
		})
	}
}
