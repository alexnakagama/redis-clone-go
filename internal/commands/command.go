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

	case "DECRBY":
		if len(parts) != 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		num, err := strconv.Atoi(parts[2])
		if err != nil {
			return "ERROR: value is not an integer\n", false, nil
		}

		value, err := st.DecrBy(parts[1], num)
		if err != nil {
			return "ERROR: " + err.Error() + "\n", false, nil
		}

		return strconv.Itoa(value) + "\n", false, nil

	case "GETSET":
		if len(parts) != 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		oldValue, err := st.GetSet(parts[1], parts[2])
		if err != nil {
			return "ERROR: invalid arguments\n", false, nil
		}

		return oldValue + "\n", false, nil

	case "GETDEL":
		if len(parts) != 2 {
			return "ERROR: missing arguments\n", false, nil
		}

		value, err := st.GetDel(parts[1])
		if err != nil {
			return "(nil)\n", false, nil
		}

		return value + "\n", false, nil

	case "PERSIST":
		if len(parts) != 2 {
			return "ERROR: missing arguments\n", false, nil
		}

		ok := st.Persist(parts[1])

		if !ok {
			return "0\n", false, nil
		}

		return "1\n", false, nil

	case "RANDOMKEY":
		if len(parts) != 1 {
			return "ERROR: too many arguments\n", false, nil
		}

		return st.RandomKey() + "\n", false, nil

	case "TOUCH":
		if len(parts) == 1 {
			return "ERROR: missing arguments\n", false, nil
		}

		counter := st.Touch(parts[1:])

		return strconv.Itoa(counter) + "\n", false, nil

	case "COPY":
		if len(parts) != 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		if st.Copy(parts[1], parts[2]) {
			return "1\n", false, nil
		}

		return "0\n", false, nil

	case "GETEX":
		if len(parts) < 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		key := parts[1]
		option := strings.ToUpper(parts[2])

		if option == "EX" {
			if len(parts) != 4 {
				return "ERROR: missing seconds\n", false, nil
			}

			seconds, err := strconv.Atoi(parts[3])
			if err != nil {
				return "ERROR: invalid seconds\n", false, nil
			}

			value, err := st.GetEx(key, option, seconds)
			if err != nil {
				return "ERROR: " + err.Error() + "\n", false, nil
			}

			return value + "\n", false, nil
		}

		if option == "PERSIST" {
			if len(parts) != 3 {
				return "ERROR: too many arguments\n", false, nil
			}

			value, err := st.GetEx(key, option, 0)
			if err != nil {
				return "ERROR: " + err.Error() + "\n", false, nil
			}

			return value + "\n", false, nil
		}

		return "ERROR: invalid option\n", false, nil

	case "SETNX":
		if len(parts) < 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		if len(parts) > 3 {
			return "ERROR: too many arguments\n", false, nil
		}

		c := st.SetNX(parts[1], parts[2])
		if c {
			return "1\n", false, nil
		}

		return "0\n", false, nil

	case "TYPE":
		if len(parts) < 2 {
			return "ERROR: missing arguments\n", false, nil
		}

		if len(parts) > 2 {
			return "ERROR: too many arguments\n", false, nil
		}

		return st.Type(parts[1]) + "\n", false, nil

	case "LPUSH":
		if len(parts) < 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		key := parts[1]
		values := parts[2:]

		length, err := st.LPush(key, values)
		if err != nil {
			return "ERROR: " + err.Error() + "\n", false, nil
		}

		strLen := strconv.Itoa(length)

		return strLen + "\n", false, nil

	case "RPUSH":
		if len(parts) < 3 {
			return "ERROR: missing arguments\n", false, nil
		}

		key := parts[1]
		values := parts[2:]

		length, err := st.RPush(key, values)
		if err != nil {
			return "ERROR: " + err.Error() + "\n", false, nil
		}

		strLen := strconv.Itoa(length)

		return strLen + "\n", false, nil

	case "LLEN":
		if len(parts) != 2 {
			return "ERROR: missing arguments\n", false, nil
		}

		key := parts[1]

		length, err := st.LLen(key)
		if err != nil {
			return "ERROR: " + err.Error() + "\n", false, nil
		}

		strLen := strconv.Itoa(length)

		return strLen + "\n", false, nil
		
	case "QUIT":
		return "OK\n", true, nil

	default:
		return "ERROR: unknown command\n", false, nil
	}
}
