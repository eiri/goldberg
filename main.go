package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var server bool
	var port, cmd string
	flag.BoolVar(&server, "server", false, "run as a server")
	flag.BoolVar(&server, "s", false, "run as a server")
	flag.StringVar(&port, "port", "7275", "server's port")
	flag.StringVar(&port, "p", "7275", "server's port")
	flag.StringVar(&cmd, "cmd", "", "command to run")
	flag.Parse()

	if server {
		s := NewServer(port)
		if err := s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}

	client := NewClient(port)
	if err := client.Do(cmd); err != nil {
		log.Fatalln(err)
	}
}
