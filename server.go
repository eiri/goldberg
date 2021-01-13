package main

import (
	"errors"
	"log"
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

	log.Printf("put item %q on queue", req.Item)

	resp.Message = "ok"
	return nil
}

type Server struct {
	Addr string
}

func NewServer(port string) *Server {
	rpc.RegisterName("queue", &Handler{})

	return &Server{Addr: ":" + port}
}

func (s *Server) ListenAndServe() error {
	log.Printf("Starting server at port %s\n", s.Addr)
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	defer func() {
		listener.Close()
		log.Println("Listener closed, shutting down the server")
	}()

	rpc.Accept(listener)
	return nil
}
