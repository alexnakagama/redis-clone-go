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

func TestProcessDelete(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"DEL", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "OK\n" {
		t.Errorf("expected OK, got: %q", response)
	}

	if closeConn {
		t.Error("DEL should not close the connection")
	}

	setResponse, _, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if setResponse != "(nil)\n" {
		t.Errorf("expected (nil), got: %q", setResponse)
	}
}

func TestProcessDeleteNonExisting(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"DEL", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "(nil)\n" {
		t.Errorf("expected (nil), got: %q", response)
	}

	if closeConn {
		t.Errorf("DEL should not close the connection")
	}
}

func TestProcessDeleteMissingKey(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"DEL"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: missing key\n" {
		t.Errorf("expected ERROR: missing key, got: %q", response)
	}

	if closeConn {
		t.Errorf("DEL should not close the connection")
	}
}

func TestProcessExists(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"EXISTS", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "1\n" {
		t.Errorf("expected 1, got: %q", response)
	}

	if closeConn {
		t.Errorf("EXISTS should not close the connection")
	}
}

func TestProcessExistsNonExisting(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"EXISTS", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "0\n" {
		t.Errorf("expected 0, got: %q", response)
	}

	if closeConn {
		t.Errorf("EXISTS should not close the connection")
	}
}

func TestProcessExistsMissingKey(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"EXISTS"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: expected 2 arguments\n" {
		t.Errorf("expected ERROR: expected 2 arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("EXISTS should not close the connection")
	}
}

func TestProcessSize(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"SIZE"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "1\n" {
		t.Errorf("expected 1, got: %q", response)
	}

	if closeConn {
		t.Errorf("SIZE should not close the connection")
	}
}

func TestProcessSizeNonExisting(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"SIZE"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "0\n" {
		t.Errorf("expected 0, got: %q", response)
	}

	if closeConn {
		t.Errorf("SIZE should not close the connection")
	}
}

func TestProcessClear(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	getResponse, _, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if getResponse != "alex\n" {
		t.Errorf("expected alex, got: %q", getResponse)
	}

	clearResponse, closeConn, err := Process([]string{"CLEAR"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if clearResponse != "OK\n" {
		t.Errorf("expected OK, got: %q", clearResponse)
	}

	if closeConn {
		t.Errorf("CLEAR should not close the connection")
	}

	response, _, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "(nil)\n" {
		t.Errorf("expected (nil), got: %q", response)
	}
}

func TestProcessClearTooManyArgs(t *testing.T) {}
