package main

import (
	"testing"
)

func TestServer(t *testing.T) {
	port := "7275"
	server := NewServer(port)
	client := NewClient(port)

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

	t.Run("Basic", func(t *testing.T) {
		if err := client.Do("ohai!"); err != nil {
			t.Error(err)
		}
	})
}
