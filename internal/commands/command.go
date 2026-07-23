package commands

import (
	"strings"

	"github.com/alexnakagama/redis-clone-go/internal/store"
)

func Process(message string, s *store.Store) (string, error) {
	message = strings.TrimSpace(message)

	switch message {
	case "PING":
		return "PONG", nil
	default:
		return "ERROR: unknown command", nil
	}
}
