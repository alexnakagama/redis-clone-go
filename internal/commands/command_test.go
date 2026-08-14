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
