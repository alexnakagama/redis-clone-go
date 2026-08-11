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
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	length := store.Size()
	if length != 1 {
		t.Fatalf("expected length 1, got: %d", length)
	}
}

func TestClear(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	store.Clear()

	_, exists := store.Get("name")
	if exists {
		t.Fatal("expected to get nil")
	}
}

func TestKeys(t *testing.T) {
	store := NewStore()
	
	store.Set("name", "alex")
	store.Set("age", "18")
	store.Set("city", "buenos aires")

	keys := store.Keys()
	if len(keys) != 3 {
		t.Fatalf("expected 3 keys, got %v", keys)
	}

	expected := map[string]bool{
		"name": true,
		"age": true,
		"city": true,
	}

	for _, key := range keys {
		if !expected[key] {
			t.Fatalf("unexpected key: %s", key)
		}
	}
}

func TestRename(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	err = store.Rename("name", "username")
	if err != nil {
		t.Fatal(err)
	}

	value, exists := store.Get("username")
	if !exists {
		t.Fatal(err)
	}

	if value != "alex" {
		t.Fatalf("expected alex, got: %s", value)
	}

	_, exists = store.Get("name")
	if exists {
		t.Fatal("expected name to not exist")
	}
}

func TestTouch(t *testing.T) {
	store := NewStore()

	store.Set("name", "alex")
	store.Set("age", "18")
	store.Set("city", "buenos aires")

	keys := []string{"name", "age", "city"}

	length := store.Touch(keys)
	if length != 3 {
		t.Fatalf("expected 3, got: %d", length)
	}
}

func TestCopy(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	copied := store.Copy("name", "username")
	if !copied {
		t.Fatal("expected key to be copied")
	}

	value, err := store.GetDel("username")
	if err != nil {
		t.Fatalf("expected username to exist")
	}

	if value != "alex" {
		t.Fatalf("expected alex, got: %s", value)
	}

	_, exists := store.Get("name")
	if !exists {
		t.Fatal("expected original key to still exists")
	}
}

func TestType(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	theTypesiculusfilocus := store.Type("name") 
	if theTypesiculusfilocus != "string" {
		t.Fatalf("expected type to be string, got: %s", theTypesiculusfilocus)
	}
}

func TestExpire(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	ok := store.Expire("name", 30)
	if !ok {
		t.Fatalf("expected expire to return true")
	}
}

func TestTTL(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	ttl, err := store.TTL("name")
	if err != nil {
		t.Fatal(err)
	}

	if ttl != -1 {
		t.Fatalf("expected ttl to be -1, got: %d", ttl)
	}
}

func TestTTLWithExpire(t *testing.T) {
	store := NewStore()

	err := store.Set("name", "alex")
	if err != nil {
		t.Fatal(err)
	}

	ok := store.Expire("name", 10)
	if !ok {
		t.Fatalf("expected expire to return true")
	}

	ttl, err := store.TTL("name")
	if err != nil {
		t.Fatal(err)
	}

	if ttl < 0 || ttl > 10 {
		t.Fatalf("expected ttl to be between 0 and 10, got: %d", ttl)
	}
}

func TestPersist(t *testing.T) {}
