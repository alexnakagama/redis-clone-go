package commands

import (
	"strings"

	"github.com/alexnakagama/redis-clone-go/internal/store"
)

func Process(message string, s *store.Store) (string, error) {
	parts := strings.Fields(message)

	if len(parts) == 0 {
		return "ERROR: empty command\n", nil
	}

	command := parts[0]

	switch command {

	case "PING":
		return "PONG", nil

	case "GET":
		if len(parts) < 2 {
			return "ERROR: missing key\n", nil
		}

		value, exists := s.Get(parts[1])
		if !exists {
			return "ERROR: value not found\n", nil
		}

		return value, nil

	case "SET":
		if len(parts) < 3 {
			return "ERROR: missing key\n", nil
		}

		err := s.Set(parts[1], parts[2])

		if err != nil {
			return "ERROR: set failed\n", err
		}

		return "OK\n", nil

	case "DEL":
		if len(parts) < 2 {
			return "ERROR: missing key\n", nil
		}

		exists := s.Delete(parts[1])
		if !exists {
			return "ERROR: key not found\n", nil
		}

		return "OK\n", nil

	default:
		return "ERROR: unknown command\n", nil
	}
}
