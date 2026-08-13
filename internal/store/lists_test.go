package store

import (
	"testing"
)

func TestLPush(t *testing.T) {
	store := NewStore()

	names := []string{"alex", "juan"}

	listLen, err := store.RPush("name", names)
	if err != nil {
		t.Fatal(err)
	}

	if listLen != 2 {
		t.Fatalf("expected 2, got: %d", listLen)
	}

	name := []string{"fran"}

	length, err := store.LPush("name", name)
	if err != nil {
		t.Fatal(err)
	}

	if length != 3 {
		t.Fatalf("expected 3, got: %d", length)
	}

	list, err := store.LRange("name", 0, 2)
	if err != nil {
		t.Fatal(err)
	}

	if list[0] != "fran" {
		t.Fatalf("expected fran, got: %s", list[0])
	}

	if list[1] != "alex" {
		t.Fatalf("expected alex, got: %s", list[1])
	}

	if list[2] != "juan" {
		t.Fatalf("expected juan got: %s", list[2])
	}
}

func TestRPush(t *testing.T) {
	store := NewStore()

	names := []string{"alex", "juan"}

	listLen, err := store.RPush("name", names)
	if err != nil {
		t.Fatal(err)
	}

	if listLen != 2 {
		t.Fatalf("expected 2, got: %d", listLen)
	}

	list, err := store.LRange("name", 0, 1)
	if err != nil {
		t.Fatal(err)
	}

	if list[0] != "alex" {
		t.Fatalf("expected alex, got: %s", list[0])
	}

	if list[1] != "juan" {
		t.Fatalf("expected juan, got: %s", list[1])
	}
}

func TestLLen(t *testing.T) {
	store := NewStore()

	names := []string{"alex", "jose"}

	_, err := store.RPush("name", names)
	if err != nil {
		t.Fatal(err)
	}

	length, err := store.LLen("name")
	if err != nil {
		t.Fatal(err)
	}

	if length != 2 {
		t.Fatalf("expected 2, got: %d", length)
	}
}

func TestLPop(t *testing.T) {}

func TestRPop(t *testing.T) {}
