package store

import (
	"errors"
	"strconv"
	"time"
)

func (s *Store) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]

	if !exists {
		return "", false
	}

	if value.Type != "string" {
		return "", false
	}

	strValue, ok := value.Data.(string)
	if !ok {
		return "", false
	}

	return strValue, true
}

func (s *Store) Set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = Value{
		Type: "string",
		Data: value,
	}

	delete(s.exp, key)

	return nil
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
	delete(s.exp, key)

	return true
}

func (s *Store) Append(key string, text string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return 0, errors.New("key not found")
	}

	if value.Type != "string" {
		return 0, errors.New("type mismatch")
	}

	strValue, ok := value.Data.(string)
	if !ok {
		return 0, errors.New("type mismatch")
	}

	strValue += text

	value.Data = strValue

	s.data[key] = value

	return len(strValue), nil
}

func (s *Store) MGet(keys []string) []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := make([]string, 0, len(keys))

	for _, key := range keys {
		s.removeExpired(key)

		value, exists := s.data[key]
		if !exists {
			values = append(values, "key doesnt exists")
			continue
		}

		if value.Type != "string" {
			values = append(values, "type mismatch")
			continue
		}

		strValue, ok := value.Data.(string)
		if !ok {
			values = append(values, "type mismatch")
			continue
		}

		values = append(values, strValue)
	}	

	return values
}

func (s *Store) StrLen(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return 0, errors.New("key not found")
	}

	if value.Type != "string" {
		return 0, errors.New("type mismatch")
	}

	strValue, ok := value.Data.(string)
	if !ok {
		return 0, errors.New("type mismatch")
	}

	return len(strValue), nil
}

func (s *Store) MSet(pairs []string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]

		value := Value{
			Type: "string",
			Data: pairs[i+1],
		}

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

	if value.Type != "string" {
		return 0, errors.New("type mismatch")
	}

	strValue, ok := value.Data.(string)
	if !ok {
		return 0, errors.New("type mismatch")
	}

	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, err
	}

	intValue += num

	strValue = strconv.Itoa(intValue)

	value.Data = strValue
	s.data[key] = value

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

	if value.Type != "string" {
		return 0, errors.New("type mismatch")
	}

	strValue, ok := value.Data.(string)
	if !ok {
		return 0, errors.New("type mismatch")
	}

	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, err
	}

	intValue -= num

	strValue = strconv.Itoa(intValue)

	value.Data = strValue
	s.data[key] = value

	return intValue, nil
}

func (s *Store) Incr(key string) (int, error) {
	return s.IncrBy(key, 1)
}

func (s *Store) Decr(key string) (int, error) {
	return s.DecrBy(key, 1)
}

func (s *Store) GetSet(key string, newValue string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	oldValue := "(nil)"

	value, exists := s.data[key]
	if exists {
		if value.Type != "string" {
			return "", errors.New("type mismatch")
		}

		strValue, ok := value.Data.(string)
		if !ok {
			return "", errors.New("type mismatch")
		}

		oldValue = strValue
	}

	s.data[key] = Value{
		Type: "string",
		Data: newValue,
	}

	delete(s.exp, key)

	return oldValue, nil
}

func (s *Store) GetDel(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	oldValue, exists := s.data[key]
	if !exists {
		return "", errors.New("key not found")
	}

	if oldValue.Type != "string" {
		return "", errors.New("type mismatch")
	}

	strValue, ok := oldValue.Data.(string)
	if !ok {
		return "", errors.New("type mismatch")
	}

	delete(s.data, key)
	delete(s.exp, key)

	return strValue, nil
}

func (s *Store) GetEx(key string, option string, seconds int) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	value, exists := s.data[key]
	if !exists {
		return "", errors.New("key not found")
	}

	if value.Type != "string" {
		return "", errors.New("type mismatch")
	}

	strValue, ok := value.Data.(string)
	if !ok {
		return "", errors.New("type mismatch")
	}

	switch option {
	case "EX":
		s.exp[key] = time.Now().Add(time.Duration(seconds) * time.Second)

	case "PERSIST":
		delete(s.exp, key)

	default:
		return "", errors.New("unknown option")
	}

	return strValue, nil
}

func (s *Store) SetNX(key, inputValue string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(key)

	_, exists := s.data[key]
	if exists {
		return false
	}

	s.data[key] = Value{
		Type: "string",
		Data: inputValue,
	}

	delete(s.exp, key)

	return true
}
