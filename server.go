package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

type Server struct {
	Addr string
}

func NewServer(port string) *Server {
	return &Server{Addr: ":" + port}
}

func (s *Server) Handler(conn net.Conn) {
	log.Println("New connection")

	defer func() {
		log.Println("Closing connection")
		conn.Close()
	}()

	timeoutDuration := 5 * time.Second
	reader := bufio.NewReader(conn)

	for {
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("%s", bytes)
	}
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

	var e error
	for {
		conn, err := listener.Accept()
		if err != nil {
			e = err
			break
		}
		go s.Handler(conn)
	}
	return e
}
