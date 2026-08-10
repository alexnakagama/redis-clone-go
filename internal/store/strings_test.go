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

func TestDelete(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	deleted := store.Delete("name")
	if !deleted {
		t.Fatal("expected key to be deleted")
	}

	_, exists := store.Get("name")
	if exists {
		t.Fatal("expected key to not exist")
	}
}

func TestStrLen(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	length, err := store.StrLen("name")
	if err != nil {
		t.Fatal(err)
	}

	if length != 4 {
		t.Fatalf("expected length 4 returned: %d", length)
	}
}

func TestStrLenEmpty(t *testing.T) {
}
