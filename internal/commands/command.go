package commands

import (
	"strings"

	"github.com/alexnakagama/redis-clone-go/internal/store"
)

func Process(message string, st *store.Store) (string, bool, error) {
	parts := strings.Fields(message)

	if len(parts) == 0 {
		return "ERROR: empty command\n", false, nil
	}

	command := strings.ToUpper(parts[0])

	switch command {

	case "PING":
		return "PONG\n", false, nil

	case "GET":
		if len(parts) < 2 {
			return "ERROR: missing key\n", false, nil
		}

		value, exists := st.Get(parts[1])
		if !exists {
			return "(nil)\n", false, nil
		}

		return value + "\n", false, nil

	case "SET":
		if len(parts) < 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		err := st.Set(parts[1], parts[2])

		if err != nil {
			return "ERROR: set failed\n", false, err
		}

		return "OK\n", false, nil

	case "DEL":
		if len(parts) < 2 {
			return "ERROR: missing key\n", false, nil
		}

		if st.Delete(parts[1]) {
			return "OK\n", false, nil
		}

		return "(nil)\n", false, nil

	case "QUIT":
		return "OK\n", true, nil

	default:
		return "ERROR: unknown command\n", false, nil
	}
}
