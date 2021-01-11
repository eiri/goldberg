package main

import (
	"context"
	"log"
	"net"
	"time"
)

type Client struct {
	Addr string
}

func NewClient(port string) *Client {
	return &Client{Addr: ":" + port}
}

func (c *Client) Do(cmd string) error {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	log.Printf("Connecting to %s", c.Addr)
	conn, err := d.DialContext(ctx, "tcp", c.Addr)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	log.Printf("=> %s", cmd)
	_, err = conn.Write([]byte(cmd + "\n"))
	return err
}
