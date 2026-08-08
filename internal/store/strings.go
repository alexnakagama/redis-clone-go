package store

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
