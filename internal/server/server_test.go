package server

import (
	"net"
	"testing"
)

func TestServerStart(t *testing.T) {
	server, err := NewServer(":6380")
	if err != nil {
		t.Fatal(err)
	}

	go server.Start()
	defer server.Close()

	conn, err := net.Dial("tcp", "localhost:6380")
	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close()
}
