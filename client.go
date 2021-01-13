package main

import (
	"log"
	"net/rpc"
)

type Client struct {
	Addr string
}

func NewClient(port string) *Client {
	return &Client{Addr: ":" + port}
}

func (c *Client) Do(cmd string) error {
	log.Printf("Connecting to %s", c.Addr)
	client, err := rpc.Dial("tcp", c.Addr)
	if err != nil {
		return err
	}
	defer client.Close()

	log.Printf("-> %s", cmd)
	resp := new(Response)
	call := client.Go("queue.Execute", &Request{Item: cmd}, resp, nil)
	reply := <-call.Done
	if reply.Error != nil {
		return reply.Error
	}
	log.Printf("<- %s", resp.Message)
	return nil
}
