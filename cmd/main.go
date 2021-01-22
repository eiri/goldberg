package main

import (
	"flag"
	"log"
	"os"

	gc "github.com/eiri/goldberg/client"
	gs "github.com/eiri/goldberg/server"
)

func main() {
	var server bool
	var port, name, cmd string
	flag.BoolVar(&server, "server", false, "run as a server")
	flag.BoolVar(&server, "s", false, "run as a server")
	flag.StringVar(&port, "port", "7275", "server's port")
	flag.StringVar(&port, "p", "7275", "server's port")
	flag.StringVar(&name, "name", "", "name of the queue")
	flag.StringVar(&name, "n", "", "name of the queue")
	flag.StringVar(&cmd, "cmd", "", "command to run")
	flag.Parse()

	if server {
		s := gs.NewServer(port)
		if err := s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}

	client := gc.NewClient(port)
	if err := client.Do(name, cmd); err != nil {
		log.Fatalln("fatal error:", err)
	}
}
