package store

import (
	"testing"
)

func TestLPush(t *testing.T) {}

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
