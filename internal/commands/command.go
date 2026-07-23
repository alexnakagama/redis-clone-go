package commands

import (
	"strings"

	"github.com/alexnakagama/redis-clone-go/internal/store"
)

func Process(message string, st *store.Store) (string, error) {
	parts := strings.Fields(message)

	if len(parts) == 0 {
		return "ERROR: empty command\n", nil
	}

	command := strings.ToUpper(parts[0])

	switch command {

	case "PING":
		return "PONG\n", nil

	case "GET":
		if len(parts) < 2 {
			return "ERROR: missing key\n", nil
		}

		value, exists := st.Get(parts[1])
		if !exists {
			return "(nil)\n", nil
		}

		return value + "\n", nil

	case "SET":
		if len(parts) < 3 {
			return "ERROR: missing arguments\n", nil
		}

		err := st.Set(parts[1], parts[2])

		if err != nil {
			return "ERROR: set failed\n", err
		}

		return "OK\n", nil

	case "DEL":
		if len(parts) < 2 {
			return "ERROR: missing key\n", nil
		}

		if st.Delete(parts[1]) {
			return "OK\n", nil
		}

		return "(nil)\n", nil

	default:
		return "ERROR: unknown command\n", nil
	}
}
