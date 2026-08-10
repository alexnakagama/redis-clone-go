package store

import (
	"testing"
)

func TestSetGet(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	value, exists := store.Get("name")
	if !exists {
		t.Fatal("expected key to exist")
	}

	if value != "alex" {
		t.Fatalf("expected alex got: %s", value)
	}
}
