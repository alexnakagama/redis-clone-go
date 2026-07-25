package store

import (
	"errors"
	"strconv"
	"sync"
)

type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
	return nil
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, exists := s.data[key]

	if !exists {
		return "", false
	}

	return value, true
}

func (s *Store) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.data[key]

	if !exists {
		return false
	}

	delete(s.data, key)
	return true
}

func (s *Store) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, exists := s.data[key]

	return exists
}

func (s *Store) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	length := len(s.data)

	return length
}

func (s *Store) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	newMap := make(map[string]string)

	s.data = newMap
}

func (s *Store) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	sKeys := make([]string, 0, len(s.data))

	for key := range s.data {
		sKeys = append(sKeys, key)
	}

	return sKeys
}

func (s *Store) Incr(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, exists := s.data[key]
	if !exists {
		return 0, errors.New("key not found")
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	intValue += 1

	strValue := strconv.Itoa(intValue)

	s.data[key] = strValue

	return intValue, nil
}

func (s *Store) Decr(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, exists := s.data[key]
	if !exists {
		return 0, errors.New("key not found")
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	intValue -= 1

	strValue := strconv.Itoa(intValue)

	s.data[key] = strValue

	return intValue, nil
}

func (s *Store) Append(text string) (int, error) {

}
