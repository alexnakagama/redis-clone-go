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

func TestAppend(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	length, err := store.Append("name", "a")
	if err != nil {
		t.Fatal(err)
	}

	if length != 5 {
		t.Fatalf("expected length 5 returned: %d", length)
	}

	value, exists := store.Get("name")
	if !exists {
		t.Fatal("expected key to exist")
	}

	if value != "alexa" {
		t.Fatalf("expected value alexa, returned: %s", value)
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

func TestIncrBy(t *testing.T) {
	store := NewStore()

	err := store.Set("age", "1")
	if err != nil {
		t.Fatal(err)
	}

	num, err := store.IncrBy("age", 20)
	if err != nil {
		t.Fatal(err)
	}

	if num != 21 {
		t.Fatalf("expected 21 returned: %d", num)
	}
}

func TestDecrBy(t *testing.T) {
	store := NewStore()

	err := store.Set("age", "30")
	if err != nil {
		t.Fatal(err)
	}

	num, err := store.DecrBy("age", 20)
	if err != nil {
		t.Fatal(err)
	}

	if num != 10 {
		t.Fatalf("expected 10 retured: %d", num)
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

	if num != 17 {
		t.Fatalf("Expected age 17 returned: %d", num)
	}
}

func TestMGet(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	keys := []string{"name"}

	values := store.MGet(keys)
	if len(values) != 1 {
		t.Fatalf("expected length 1, got: %d", len(values))
	}

	if values[0] != "alex" {
		t.Fatalf("expected alex, got: %s", values[0])
	}
}

func TestMGetMultiple(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	err = store.Set("username", "mike")
	if err != nil {
		t.Fatal(err)
	}

	keys := []string{"name", "username"}

	values := store.MGet(keys)
	if len(values) != 2 {
		t.Fatalf("expected length 2, got: %d", len(values))
	}

	if values[0] != "alex" || values[1] != "mike" {
		t.Fatalf("expected [alex mike], got: %v", values)
	}
}

func TestGetDel(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	value, err := store.GetDel("name")
	if err != nil {
		t.Fatal(err)
	}

	if value != "alex" {
		t.Fatalf("expected alex, got: %s", value)
	}

	_, exists := store.Get("name")
	if exists {
		t.Fatalf("expected key to not exist")
	}
}

func TestSetNX(t *testing.T) {
	store := NewStore()

	ok := store.SetNX("name", "alex")
	if !ok {
		t.Fatalf("expected SetNX to return true")
	}

	value, exists := store.Get("name")
	if !exists {
		t.Fatalf("expected key to exist")
	}

	if value != "alex" {
		t.Fatalf("expected alex: got: %s", value)
	}
}

func TestSetNXKeyExist(t *testing.T) {
	store := NewStore()

	store.Set("name", "alex")

	ok := store.SetNX("name", "juan")
	if ok {
		t.Fatalf("expected SetNX to return false")
	}

	value, _ := store.Get("name")
	if value != "alex" {
		t.Fatalf("expected alex, got: %s", value)
	}
}

func TestGetSet(t *testing.T) {}
