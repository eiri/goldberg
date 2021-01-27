package server_test

import (
	"net"
	"net/rpc"
	"strconv"
	"testing"
	"time"

	gh "github.com/eiri/goldberg/handler"
	gs "github.com/eiri/goldberg/server"
)

func TestServer(t *testing.T) {
	port := "7275"
	name := "test"
	items := []string{"apple", "banana", "cherry", "date", "elderberry", "fig"}

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

	// wait for port alloc
	try := 1
	for {
		conn, err := net.DialTimeout("tcp", ":"+port, 100*time.Millisecond)
		if try > 3 {
			t.Fatalf("Failed to start RPC server: %s", err)
		}
		if conn != nil {
			conn.Close()
			break
		}
		time.Sleep(100 * time.Millisecond)
		try++
	}

	client, err := rpc.Dial("tcp", ":"+port)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	req := &gh.Request{Name: name}
	resp := &gh.Response{}

	t.Run("Create", func(t *testing.T) {
		if err = client.Call("goldberg.Create", req, resp); err != nil {
			t.Error(err)
		}
		expect := "ok"
		if resp.Message != expect {
			t.Errorf("Expected response %q, got %q", expect, resp.Message)
		}
	})

	t.Run("Enqueue", func(t *testing.T) {
		for i, item := range items {
			r := &gh.Request{Name: name, Item: item}
			if err = client.Call("goldberg.PushBack", r, resp); err != nil {
				t.Error(err)
			}
			expect := "ok"
			if resp.Message != expect {
				t.Errorf("Expected response %q, got %q", expect, resp.Message)
			}

			if err = client.Call("goldberg.Len", req, resp); err != nil {
				t.Error(err)
			}
			expect = strconv.Itoa(i + 1)
			if resp.Message != expect {
				t.Errorf("Expected response %q, got %q", expect, resp.Message)
			}
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		for i, item := range items {
			if err = client.Call("goldberg.PopFront", req, resp); err != nil {
				t.Error(err)
			}
			if resp.Message != item {
				t.Errorf("Expected response %q, got %q", item, resp.Message)
			}

			if err = client.Call("goldberg.Len", req, resp); err != nil {
				t.Error(err)
			}
			expect := strconv.Itoa(len(items) - i - 1)
			if resp.Message != expect {
				t.Errorf("Expected response %q, got %q", expect, resp.Message)
			}
		}
	})

	t.Run("Empty", func(t *testing.T) {
		if err = client.Call("goldberg.Len", req, resp); err != nil {
			t.Error(err)
		}
		expect := "0"
		if resp.Message != expect {
			t.Errorf("Expected response %q, got %q", expect, resp.Message)
		}
	})
}
