package commands

import (
	"strings"

	"github.com/alexnakagama/redis-clone-go/internal/store"
)

func Process(message string, s *store.Store) (string, error) {
	parts := strings.Fields(message)

	if len(parts) == 0 {
		return "ERROR: empty command", nil
	}

	command := parts[0]

	switch command {

	case "PING":
		return "PONG", nil

	case "GET":
		if len(parts) < 2 {
			return "ERROR: missing key", nil
		}

		value, exists := s.Get(parts[1])
		if !exists {
			return "(nil)", nil
		}

		return value, nil

	case "SET":
		if len(parts) < 3 {
			return "ERROR: missing key", nil
		}

		err := s.Set(parts[1], parts[2])

		if err != nil {
			return "ERROR: set failed", err
		}

		return "OK", nil

	case "DEL":
		if len(parts) < 2 {
			return "ERROR: missing key", nil
		}

		exists := s.Delete(parts[1])
		if !exists {
			return "ERROR: key not found", nil
		}

		return "OK", nil

	default:
		return "ERROR: unknown command", nil
	}
}
