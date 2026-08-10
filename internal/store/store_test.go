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
}
