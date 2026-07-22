package commands

import (
	"strings"
)

func Process(message string) (string, error) {
	message = strings.TrimSpace(message)

	switch message {
	case "PING":
		return "PONG", nil
	default:
		return "ERROR: unknown command", nil
	}
}
