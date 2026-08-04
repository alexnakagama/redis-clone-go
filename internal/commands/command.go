package commands

import (
	"strconv"
	"strings"

	"github.com/alexnakagama/redis-clone-go/internal/store"
)

// This function process all the functions defined in the store.go file
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

	case "EXISTS":
		if len(parts) < 2 {
			return "ERROR: missing arguments\n", false, nil
		}

		exists := st.Exists(parts[1])

		if exists {
			return "1\n", false, nil
		}

		return "0\n", false, nil

	case "SIZE":
		length := st.Size()

		if length == 0 {
			return "0\n", false, nil
		}

		response := strconv.Itoa(length)

		return response + "\n", false, nil

	case "CLEAR":
		if len(parts) > 1 {
			return "ERROR: too many arguments\n", false, nil
		}

		st.Clear()
		return "OK\n", false, nil

	case "KEYS":
		if len(parts) > 1 {
			return "ERROR: too many arguments\n", false, nil
		}

		sKeys := st.Keys()

		if len(sKeys) == 0 {
			return "0\n", false, nil
		}

		return strings.Join(sKeys, "\n"), false, nil

	case "INCR":
		if len(parts) < 2 {
			return "ERROR: missing arguments\n", false, nil
		}

		intValue, err := st.Incr(parts[1])
		if err != nil {
			return "ERROR: " + err.Error() + "\n", false, nil
		}

		return strconv.Itoa(intValue) + "\n", false, nil

	case "DECR":
		if len(parts) < 2 {
			return "ERROR: missing arguments\n", false, nil
		}

		intValue, err := st.Decr(parts[1])
		if err != nil {
			return "ERROR: " + err.Error() + "\n", false, nil
		}

		return strconv.Itoa(intValue) + "\n", false, nil

	case "APPEND":
		if len(parts) < 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		length, err := st.Append(parts[1], parts[2])
		if err != nil {
			return "ERROR: " + err.Error() + "\n", false, nil
		}

		return strconv.Itoa(length) + "\n", false, nil

	case "RENAME":
		if len(parts) < 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		err := st.Rename(parts[1], parts[2])
		if err != nil {
			return "ERROR: " + err.Error() + "\n", false, nil
		}

		return "OK\n", false, nil

	case "MGET":
		if len(parts) == 1 {
			return "ERROR: missing arguments\n", false, nil
		}

		keys := parts[1:]

		values := st.MGet(keys)

		response := strings.Join(values, "\n")

		return response + "\n", false, nil

	case "STRLEN":
		if len(parts) < 2 {
			return "ERROR: missing arguments\n", false, nil
		}

		length, err := st.StrLen(parts[1])
		if err != nil {
			return "ERROR: " + err.Error() + "\n", false, nil
		}

		return strconv.Itoa(length) + "\n", false, nil

	case "EXPIRE":
		if len(parts) != 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		seconds, err := strconv.Atoi(parts[2])
		if err != nil {
			return "ERROR: invalid time\n", false, nil
		}

		ok := st.Expire(parts[1], seconds)

		if ok {
			return "1\n", false, nil
		}

		return "0\n", false, nil

	case "TTL":
		if len(parts) < 2 {
			return "ERROR: missing arguments\n", false, nil
		}

		seconds, err := st.TTL(parts[1])
		if err != nil {
			return "ERROR: " + err.Error() + "\n", false, nil
		}

		return strconv.Itoa(seconds) + "\n", false, nil

	case "MSET":
		if len(parts) < 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		if (len(parts)-1)%2 != 0 {
			return "ERROR: arguments must be key value pairs\n", false, nil
		}

		st.MSet(parts[1:])

		return "OK\n", false, nil

	case "INCRBY":
		if len(parts) != 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		num, err := strconv.Atoi(parts[2])
		if err != nil {
			return "ERROR: value is not an integer\n", false, nil
		}

		value, err := st.IncrBy(parts[1], num)
		if err != nil {
			return "ERROR: " + err.Error() + "\n", false, nil
		}

		return strconv.Itoa(value) + "\n", false, nil

	case "QUIT":
		return "OK\n", true, nil

	default:
		return "ERROR: unknown command\n", false, nil
	}
}
