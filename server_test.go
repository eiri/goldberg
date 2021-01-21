package main

import (
	"net/rpc"
	"testing"
)

func TestServer(t *testing.T) {
	port := "7275"
	server := NewServer(port)

	ech := make(chan error)
	go func() {
		defer close(ech)
		if err := server.ListenAndServe(); err != nil {
			ech <- err
		}
	}()
	select {
	case err := <-ech:
		t.Fatal(err)
	default:
	}

	t.Run("Basic call", func(t *testing.T) {
		client, err := rpc.Dial("tcp", ":"+port)
		if err != nil {
			t.Fatal(err)
		}
		defer client.Close()

		req := &Request{Item: "item"}
		resp := new(Response)
		err = client.Call("queue.Execute", req, resp)
		if err != nil {
			t.Error(err)
		}
		expect := "ok"
		if resp.Message != expect {
			t.Errorf("Expected response %q, got %q", expect, resp.Message)
		}
	})
}
