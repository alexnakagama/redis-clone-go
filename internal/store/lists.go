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
		return 0, errors.New("type mismtach")
	}

	list, ok := value.Data.([]string)
	if !ok {
		return 0, errors.New("type mismtach")
	}

	slices.Reverse(values)

	list = append(values, list...)

	value.Data = list
	s.data[key] = value

	return len(list), nil
}
