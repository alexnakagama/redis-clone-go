package store

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Set(key, value string) error {
	s.data[key] = value
	return nil
}

func (s *Store) Get(key string) (string, bool) {
	value, exists := s.data[key]

	if !exists {
		return "", false
	}

	return value, true
}

func (s *Store) Delete(key string) bool {
	_, exists := s.data[key]

	if !exists {
		return false
	}

	delete(s.data, key)
	return true
}
