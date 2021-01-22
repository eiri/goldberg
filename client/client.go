package client

import (
	"log"
	"net/rpc"

	gh "github.com/eiri/goldberg/handler"
)

type Client struct {
	Addr string
}

func NewClient(port string) *Client {
	return &Client{Addr: ":" + port}
}

func (c *Client) Do(name, cmd string) error {
	log.Printf("Connecting to %s", c.Addr)
	client, err := rpc.Dial("tcp", c.Addr)
	if err != nil {
		return err
	}
	defer client.Close()

	log.Printf("-> %s %s", name, cmd)
	resp := new(gh.Response)
	req := &gh.Request{Name: name}
	call := client.Go("goldberg.Create", req, resp, nil)
	reply := <-call.Done
	if reply.Error != nil {
		return reply.Error
	}
	log.Printf("<- %s", resp.Message)
	return nil
}
