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
	store := NewStore()

	err := store.Set("name", "")
	if err != nil {
		t.Fatal(err)
	}

	length, err := store.StrLen("name")
	if err != nil {
		t.Fatal(err)
	}

	if length != 0 {
		t.Fatalf("expected length 0 returned: %d", length)
	}
}

func TestIncr(t *testing.T) {
	store := NewStore()

	err := store.Set("age", "18")
	if err != nil {
		t.Fatal(err)
	}

	num, err := store.Incr("age")
	if err != nil {
		t.Fatal(err)
	}

	if num != 19 {
		t.Fatalf("Expected age 19 returned: %d", num)
	}
}

func TestDecr(t *testing.T) {
	store := NewStore()

	err := store.Set("age", "18")
	if err != nil {
		t.Fatal(err)
	}

	num, err := store.Decr("age")
	if err != nil {
		t.Fatal(err)
	}

	if num != 19 {
		t.Fatalf("Expected age 17 returned: %d", num)
	}
}
