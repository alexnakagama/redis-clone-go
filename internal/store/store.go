package store

import ()

type Store struct {
	data map[string]string
}

func (s *Store) Set(key, value string) error {
	return nil
}

func (s *Store) Get(key string) (string, bool) {
}

func (s *Store) Delete(key string) error {
	return nil
}
