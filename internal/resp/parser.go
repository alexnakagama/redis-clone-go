package resp

import (
	"bufio"
	"errors"
	"strings"
)

func Parse(r *bufio.Reader) ([]string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return []string{}, err
	}

	line = strings.TrimSpace(line)

	if line[0] != '*' {
		return []string{}, errors.New("invalid RESP")
	}

	countStr := line[1:]
}
