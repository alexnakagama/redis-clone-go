package store

import (
	"errors"
	"slices"
)

func (s *Store) LPush(key string, values []string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		list := make([]string, len(values))

		for i := range values {
			list[len(values)-1-i] = values[i]
		}

		s.data[key] = Value{
			Type: "list",
			Data: list,
		}

		return len(list), nil
	}

	if value.Type != "list" {
		return 0, errors.New("type mismatch")
	}

	list, ok := value.Data.([]string)
	if !ok {
		return 0, errors.New("type mismatch")
	}

	slices.Reverse(values)

	list = append(values, list...)

	value.Data = list
	s.data[key] = value

	return len(list), nil
}

func (s *Store) RPush(key string, values []string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		s.data[key] = Value{
			Type: "list",
			Data: values,
		}
		
		return len(values), nil
	}

	if value.Type != "list" {
		return 0, errors.New("type mismatch")
	}

	list, ok := value.Data.([]string)
	if !ok {
		return 0, errors.New("type mismatch")
	}

	list = append(list, values...)

	value.Data = list
	s.data[key] = value

	return len(list), nil
}

func (s *Store) LLen(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return 0, nil
	}

	if value.Type != "list" {
		return 0, errors.New("type mismatch")
	}

	list, ok := value.Data.([]string)
	if !ok {
		return 0, errors.New("type mismatch")
	}

	return len(list), nil
}

func (s *Store) LPop(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return "", errors.New("key doesnt exists")
	}

	if value.Type != "list" {
		return "", errors.New("type mismatch")
	}

	list, ok := value.Data.([]string)
	if !ok {
		return "", errors.New("type mismatch")
	}

	if len(list) == 0 {
		return "(nil)", nil
	}

	first := list[0]
	list = list[1:]

	value.Data = list
	s.data[key] = value

	return first, nil
}

func (s *Store) RPop(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return "", errors.New("key doesnt exists")
	}

	if value.Type != "list" {
		return "", errors.New("type mismatch")
	}

	list, ok := value.Data.([]string)
	if !ok {
		return "", errors.New("type mismatch")
	}

	last := list[len(list)-1]
	list = list[:len(list)-1]

	value.Data = list
	s.data[key] = value

	return last, nil
}

func (s *Store) LRange(key string, start, stop int) ([]string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	if start < 0 || stop < 0 {
		return nil, errors.New("invalid index")
	}

	value, exists := s.data[key]
	if !exists {
		return nil, errors.New("key doesnt exists")
	}

	if value.Type != "list" {
		return nil, errors.New("type mismatch")
	}

	list, ok := value.Data.([]string)
	if !ok {
		return nil, errors.New("type mismatch")
	}

	if start >= len(list) || start > stop {
		return []string{}, nil
	}

	if stop >= len(list) {
		stop = len(list) - 1
	}

	list = list[start:stop+1]

	return list, nil
} 

func (s *Store) LIndex(key string, index int) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return "", errors.New("key not found")
	}

	if value.Type != "list" {
		return "", errors.New("type mismatch")
	}

	list, ok := value.Data.([]string)
	if !ok {
		return "", errors.New("type mismatch")
	}

	if index < 0 || index >= len(list) {
		return "", errors.New("index out of range")
	}

	return list[index], nil
}

func (s *Store) LSet(key string, index int, newValue string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return errors.New("key doesnt exists")
	}

	if value.Type != "list" {
		return errors.New("type mismatch")
	}

	list, ok := value.Data.([]string)
	if !ok {
		return errors.New("type mismatch")
	}

	if index < 0 || index >= len(list) {
		return errors.New("index out of range")
	}

	list[index] = newValue
	
	value.Data = list
	s.data[key] = value

	return nil
}

func (s *Store) LTrim(key string, start, stop int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return errors.New("key doesnt exists")
	}

	if value.Type != "list" {
		return errors.New("type mismatch")
	}

	list, ok := value.Data.([]string)
	if !ok {
		return errors.New("type mismatch")
	}

	if start < 0 || start >= len(list) {
		return errors.New("index out of range")
	}

	if stop < 0 {
		return errors.New("index out of range")
	}

	if stop >= len(list) {
		stop = len(list) - 1
	}

	list = list[start:stop+1]

	value.Data = list
	s.data[key] = value

	return nil
}

func (s *Store) LPos(key, value string) (int, error) {
}
