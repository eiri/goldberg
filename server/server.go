package server

import (
	"net"
	"net/rpc"

	gh "github.com/eiri/goldberg/handler"
)

type Server struct {
	Addr string
	*rpc.Server
}

func NewServer(port string) *Server {
	s := rpc.NewServer()
	h := gh.New()
	s.RegisterName("goldberg", h)

	return &Server{":" + port, s}
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go s.ServeConn(conn)
	}
}
