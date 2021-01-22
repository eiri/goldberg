package server_test

import (
	"net/rpc"
	"testing"

	gh "github.com/eiri/goldberg/handler"
	gs "github.com/eiri/goldberg/server"
)

func TestServer(t *testing.T) {
	port := "7275"
	server := gs.NewServer(port)

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

	t.Run("Create", func(t *testing.T) {
		client, err := rpc.Dial("tcp", ":"+port)
		if err != nil {
			t.Fatal(err)
		}
		defer client.Close()

		req := &gh.Request{Name: "test"}
		resp := new(gh.Response)
		err = client.Call("goldberg.Create", req, resp)
		if err != nil {
			t.Error(err)
		}
		expect := "ok"
		if resp.Message != expect {
			t.Errorf("Expected response %q, got %q", expect, resp.Message)
		}
	})

	t.Run("PushBack", func(t *testing.T) {
		client, err := rpc.Dial("tcp", ":"+port)
		if err != nil {
			t.Fatal(err)
		}
		defer client.Close()

		req := &gh.Request{Name: "test", Item: "apple"}
		resp := new(gh.Response)
		err = client.Call("goldberg.PushBack", req, resp)
		if err != nil {
			t.Error(err)
		}
		expect := "ok"
		if resp.Message != expect {
			t.Errorf("Expected response %q, got %q", expect, resp.Message)
		}
	})

	t.Run("Len", func(t *testing.T) {
		client, err := rpc.Dial("tcp", ":"+port)
		if err != nil {
			t.Fatal(err)
		}
		defer client.Close()

		req := &gh.Request{Name: "test"}
		resp := new(gh.Response)
		err = client.Call("goldberg.Len", req, resp)
		if err != nil {
			t.Error(err)
		}
		expect := "1"
		if resp.Message != expect {
			t.Errorf("Expected response %q, got %q", expect, resp.Message)
		}
	})

	t.Run("PopFront", func(t *testing.T) {
		client, err := rpc.Dial("tcp", ":"+port)
		if err != nil {
			t.Fatal(err)
		}
		defer client.Close()

		req := &gh.Request{Name: "test"}
		resp := new(gh.Response)
		err = client.Call("goldberg.PopFront", req, resp)
		if err != nil {
			t.Error(err)
		}
		expect := "apple"
		if resp.Message != expect {
			t.Errorf("Expected response %q, got %q", expect, resp.Message)
		}

		err = client.Call("goldberg.Len", req, resp)
		if err != nil {
			t.Error(err)
		}
		expect = "0"
		if resp.Message != expect {
			t.Errorf("Expected response %q, got %q", expect, resp.Message)
		}
	})

}
