package store

import (
	"testing"
)

func TestNewStore(t *testing.T) {
	store := NewStore()

	if store == nil {
		t.Fatal("expected store not to be nil")
	}
}

func TestExists(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	exists := store.Exists("name")
	if !exists {
		t.Fatal("expected key to exist")
	}
}

func TestSize(t *testing.T) {
}
