package main

import (
	"errors"
	"net"
	"net/rpc"
)

type Request struct {
	Item string
}

type Response struct {
	Message string
}

type Handler struct{}

func (h *Handler) Execute(req Request, resp *Response) error {
	if req.Item == "" {
		return errors.New("An item must be specified")
	}

	resp.Message = "ok"
	return nil
}

type Server struct {
	Addr string
	*rpc.Server
}

func NewServer(port string) *Server {
	s := rpc.NewServer()
	s.RegisterName("queue", &Handler{})

	return &Server{":" + port, s}
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	defer func() {
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go s.ServeConn(conn)
	}
	return nil
}
