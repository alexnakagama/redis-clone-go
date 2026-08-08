package store

import (
	"errors"
	"math/rand/v2"
	"sync"
	"time"
)

type Value struct {
	Type string
	Data any
}

type Store struct {
	mu   sync.RWMutex
	data map[string]Value
	exp  map[string]time.Time
}

func NewStore() *Store {
	s := &Store{
		data: make(map[string]Value),
		exp:  make(map[string]time.Time),
	}

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			now := time.Now()

			s.mu.Lock()

			for key, expireTime := range s.exp {
				if now.After(expireTime) {
					delete(s.data, key)
					delete(s.exp, key)
				}
			}

			s.mu.Unlock()
		}
	}()

	return s
}

func (s *Store) removeExpired(key string) {
	expireTime, exists := s.exp[key]
	if !exists {
		return
	}

	if time.Now().After(expireTime) {
		delete(s.data, key)
		delete(s.exp, key)
	}
}

func (s *Store) Exists(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	_, exists := s.data[key]

	return exists
}

func (s *Store) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	for key := range s.data {
		s.removeExpired(key)
	}

	length := len(s.data)

	return length
}

func (s *Store) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = make(map[string]Value)
	s.exp = make(map[string]time.Time)
}

func (s *Store) Keys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	sKeys := make([]string, 0, len(s.data))

	for key := range s.data {
		s.removeExpired(key)

		_, exists := s.data[key]
		if !exists {
			continue
		}

		sKeys = append(sKeys, key)
	}

	return sKeys
}

func (s *Store) Rename(oldKey string, newKey string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(oldKey)

	value, exists := s.data[oldKey]
	if !exists {
		return errors.New("key not found")
	}

	if oldKey == newKey {
		return nil
	}

	delete(s.exp, newKey)

	s.data[newKey] = value

	expireTime, exists := s.exp[oldKey]
	if exists {
		s.exp[newKey] = expireTime
		delete(s.exp, oldKey)
	}

	delete(s.data, oldKey)

	return nil
}

func (s *Store) Expire(key string, seconds int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	_, exists := s.data[key]
	if !exists {
		return false
	}

	s.exp[key] = time.Now().Add(time.Duration(seconds) * time.Second)

	return true
}

func (s *Store) TTL(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	expireTime, exists := s.exp[key]
	if !exists {
		return -1, nil
	}

	seconds := int(time.Until(expireTime).Seconds())

	if seconds < 0 {
		return -2, nil
	}

	return seconds, nil
}

func (s *Store) Persist(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	_, dataExists := s.data[key]
	if !dataExists {
		return false
	}

	_, expExists := s.exp[key]
	if !expExists {
		return false
	}

	delete(s.exp, key)

	return true
}

func (s *Store) RandomKey() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	randomKeys := []string{}

	for key := range s.data {
		s.removeExpired(key)

		_, exists := s.data[key]
		if !exists {
			continue
		}

		randomKeys = append(randomKeys, key)
	}

	if len(randomKeys) == 0 {
		return ""
	}

	index := rand.IntN(len(randomKeys))

	return randomKeys[index]
}

func (s *Store) Touch(keys []string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	counter := 0

	for _, key := range keys {
		s.removeExpired(key)

		_, exists := s.data[key]
		if exists {
			counter += 1
		}
	}

	return counter
}

func (s *Store) Copy(src string, dest string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(src)

	value, exists := s.data[src]
	if !exists {
		return false
	}

	s.data[dest] = value
	delete(s.exp, dest)

	return true
}

func (s *Store) Type(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return "none"
	}

	return value.Type
}
