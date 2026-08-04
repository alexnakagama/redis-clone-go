package store

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

type Store struct {
	mu   sync.RWMutex
	data map[string]string
	exp  map[string]time.Time
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
		exp:  make(map[string]time.Time),
	}
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

func (s *Store) Set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
	return nil
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]

	if !exists {
		return "", false
	}

	return value, true
}

func (s *Store) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	_, exists := s.data[key]

	if !exists {
		return false
	}

	delete(s.data, key)
	return true
}

func (s *Store) Exists(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

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

func (s *Store) Append(key string, text string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return 0, errors.New("key not found")
	}

	s.data[key] = value + text

	return len(s.data[key]), nil
}

func (s *Store) Rename(oldKey string, newKey string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, exists := s.data[oldKey]
	if !exists {
		return errors.New("key not found")
	}

	if oldKey == newKey {
		return nil
	}

	s.data[newKey] = value

	delete(s.data, oldKey)

	return nil
}

func (s *Store) MGet(keys []string) []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := make([]string, 0, len(keys))

	for _, key := range keys {
		s.removeExpired(key)

		value, exists := s.data[key]

		if !exists {
			values = append(values, "(nil)")
			continue
		}

		values = append(values, value)
	}

	return values
}

func (s *Store) StrLen(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return 0, errors.New("key noy found")
	}

	length := len(value)

	return length, nil
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

func (s *Store) MSet(pairs []string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		value := pairs[i+1]

		s.data[key] = value

		delete(s.exp, key)
	}
}

func (s *Store) IncrBy(key string, num int) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return 0, errors.New("key not found")
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	intValue += num

	strValue := strconv.Itoa(intValue)

	s.data[key] = strValue

	return intValue, nil
}

func (s *Store) DecrBy(key string, num int) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return 0, errors.New("key not found")
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	intValue -= num

	strValue := strconv.Itoa(intValue)

	s.data[key] = strValue

	return intValue, nil
}
