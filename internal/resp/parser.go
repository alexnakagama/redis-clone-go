package resp

import (
	"bufio"
	"errors"
	"strconv"
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

	arrayCount, err := strconv.Atoi(countStr)
	if err != nil {
		return []string{}, err
	}

	result := make([]string, 0, arrayCount)

	for range arrayCount {
		bulkLine, err := r.ReadString('\n')
		if err != nil {
			return []string{}, err
		}

		bulkLine = strings.TrimSpace(bulkLine)

		if bulkLine[0] != '$' {
			return []string{}, errors.New("invalid RESP")
		}

		lengthStr := bulkLine[1:]

		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return []string{}, err
		}

		buffer := make([]byte, length)

		_, err = r.Read(buffer)
		if err != nil {
			return []string{}, err
		}

		result = append(result, string(buffer))

		_, err = r.ReadString('\n')
		if err != nil {
			return []string{}, err
		}
	}

	return result, nil
}
