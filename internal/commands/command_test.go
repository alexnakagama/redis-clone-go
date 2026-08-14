package commands

import (
	"testing"

	"github.com/alexnakagama/redis-clone-go/internal/store"
)

func TestProcessPing(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"PING"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "PONG\n" {
		t.Errorf("expected %q, got: %q", response, "PONG\n")
	}

	if closeConn {
		t.Error("PING should not close the connection")
	}
}

func TestProcessSet(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "OK\n" {
		t.Errorf("expected %q, got: %q", response, "OK\n")
	}

	if closeConn {
		t.Error("SET should not close the connection")
	}
}

func TestProcessGet(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "alex\n" {
		t.Errorf("expected alex, got: %q", response)
	}

	if closeConn {
		t.Error("GET should not close the connection")
	}
}

func TestProcessGetNonExisting(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "(nil)\n" {
		t.Errorf("expected (nil), got: %q", response)
	}

	if closeConn {
		t.Error("GET should not close the connection")
	}
}

func TestProcessGetMissingKey(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"GET"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: missing key\n" {
		t.Errorf("expected ERROR: missing key, got: %q", response)
	}

	if closeConn {
		t.Error("GET should not close the connection")
	}
}

func TestProcessDelete(t *testing.T) {}
