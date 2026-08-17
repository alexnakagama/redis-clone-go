package commands

import (
	"testing"

	"github.com/alexnakagama/redis-clone-go/internal/store"
)

func TestProcessPing(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"PING"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "PONG\n" {
		t.Errorf("expected %q, got: %q", response, "PONG\n")
	}

	if closeConn {
		t.Error("PING should not close the connection")
	}
}

func TestProcessSet(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "OK\n" {
		t.Errorf("expected %q, got: %q", response, "OK\n")
	}

	if closeConn {
		t.Error("SET should not close the connection")
	}
}

func TestProcessGet(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "alex\n" {
		t.Errorf("expected alex, got: %q", response)
	}

	if closeConn {
		t.Error("GET should not close the connection")
	}
}

func TestProcessGetNonExisting(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "(nil)\n" {
		t.Errorf("expected (nil), got: %q", response)
	}

	if closeConn {
		t.Error("GET should not close the connection")
	}
}

func TestProcessGetMissingKey(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"GET"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: missing key\n" {
		t.Errorf("expected ERROR: missing key, got: %q", response)
	}

	if closeConn {
		t.Error("GET should not close the connection")
	}
}

func TestProcessDelete(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"DEL", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "OK\n" {
		t.Errorf("expected OK, got: %q", response)
	}

	if closeConn {
		t.Error("DEL should not close the connection")
	}

	setResponse, _, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if setResponse != "(nil)\n" {
		t.Errorf("expected (nil), got: %q", setResponse)
	}
}

func TestProcessDeleteNonExisting(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"DEL", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "(nil)\n" {
		t.Errorf("expected (nil), got: %q", response)
	}

	if closeConn {
		t.Errorf("DEL should not close the connection")
	}
}

func TestProcessDeleteMissingKey(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"DEL"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: missing key\n" {
		t.Errorf("expected ERROR: missing key, got: %q", response)
	}

	if closeConn {
		t.Errorf("DEL should not close the connection")
	}
}

func TestProcessExists(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"EXISTS", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "1\n" {
		t.Errorf("expected 1, got: %q", response)
	}

	if closeConn {
		t.Errorf("EXISTS should not close the connection")
	}
}

func TestProcessExistsNonExisting(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"EXISTS", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "0\n" {
		t.Errorf("expected 0, got: %q", response)
	}

	if closeConn {
		t.Errorf("EXISTS should not close the connection")
	}
}

func TestProcessExistsMissingKey(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"EXISTS"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: expected 2 arguments\n" {
		t.Errorf("expected ERROR: expected 2 arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("EXISTS should not close the connection")
	}
}

func TestProcessSize(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"SIZE"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "1\n" {
		t.Errorf("expected 1, got: %q", response)
	}

	if closeConn {
		t.Errorf("SIZE should not close the connection")
	}
}

func TestProcessSizeNonExisting(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"SIZE"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "0\n" {
		t.Errorf("expected 0, got: %q", response)
	}

	if closeConn {
		t.Errorf("SIZE should not close the connection")
	}
}

func TestProcessClear(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	getResponse, _, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if getResponse != "alex\n" {
		t.Errorf("expected alex, got: %q", getResponse)
	}

	clearResponse, closeConn, err := Process([]string{"CLEAR"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if clearResponse != "OK\n" {
		t.Errorf("expected OK, got: %q", clearResponse)
	}

	if closeConn {
		t.Errorf("CLEAR should not close the connection")
	}

	response, _, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "(nil)\n" {
		t.Errorf("expected (nil), got: %q", response)
	}
}

func TestProcessClearTooManyArgs(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"CLEAR", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: too many arguments\n" {
		t.Errorf("expected ERROR: too many arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("CLEAR should not close the connection")
	}
}

func TestProcessKeys(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"KEYS"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "name" {
		t.Errorf("expected name, got: %q", response)
	}

	if closeConn {
		t.Errorf("KEYS should not close the connection")
	}
}

func TestProcessKeysNonExisting(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"KEYS"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "0\n" {
		t.Errorf("expected 0, got: %q", response)
	}

	if closeConn {
		t.Errorf("KEYS should not close the connection")
	}
}

func TestProcessKeysTooManyArgs(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"KEYS", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: too many arguments\n" {
		t.Errorf("expected ERROR: too many arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("KEYS should not close the connection")
	}
}

func TestProcessIncr(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "age", "18"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"INCR", "age"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "19\n" {
		t.Errorf("expected 19, got: %q", response)
	}

	if closeConn {
		t.Errorf("INCR should not close the connection")
	}
}

func TestProcessIncrMissingKey(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"INCR"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: missing arguments\n" {
		t.Errorf("expected ERROR: missing arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("INCR should not close the connection")
	}
}

func TestProcessDecr(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "age", "18"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"DECR", "age"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "17\n" {
		t.Errorf("expected 17, got: %q", response)
	}

	if closeConn {
		t.Errorf("DECR should not close the connection")
	}
}

func TestProcessDecrMissingKey(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"DECR"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: missing arguments\n" {
		t.Errorf("expected ERROR: missing arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("DECR should not close the connection")
	}
}

func TestProcessIncrBy(t *testing.T) {
	st := store.NewStore()
	
	_, _, err := Process([]string{"SET", "age", "18"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"INCRBY", "age", "2"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "20\n" {
		t.Errorf("expected 20, got: %q", response)
	}

	if closeConn {
		t.Errorf("INCRBY should not close the connection")
	}
}

func TestProcessIncrByMissingKey(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"INCRBY"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: expected 3 arguments\n" {
		t.Errorf("expected ERROR: expected 3 arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("INCRBY should not close the connection")
	}
}

func TestProcessDecrBy(t *testing.T) {
	st := store.NewStore()
	
	_, _, err := Process([]string{"SET", "age", "18"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"DECRBY", "age", "2"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "16\n" {
		t.Errorf("expected 16, got: %q", response)
	}

	if closeConn {
		t.Errorf("DECRBY should not close the connection")
	}
}

func TestProcessDecrByMissingKey(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"DECRBY"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: missing arguments\n" {
		t.Errorf("expected ERROR: missing arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("DECRBY should not close the connection")
	}
}

func TestProcessAppend(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"APPEND", "name", "a"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "5\n" {
		t.Errorf("expected 5, got: %q", response)
	}

	getResponse, _, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if getResponse != "alexa\n" {
		t.Errorf("expected alexa, got: %q", getResponse)
	}

	if closeConn {
		t.Errorf("APPEND should not close the connection")
	}
}

func TestProcessAppendMissingArgs(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"APPEND"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: expected 3 arguments\n" {
		t.Errorf("expected ERROR: expected 3 arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("APPEND should not close the connection")
	}
}

func TestProcessRename(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"RENAME", "name", "username"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "OK\n" {
		t.Errorf("expected OK, got: %q", response)
	}

	getResponse, _, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if getResponse != "(nil)\n" {
		t.Errorf("expected (nil), got: %q", getResponse)
	}

	realValue, _, err := Process([]string{"GET", "username"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if realValue != "alex\n" {
		t.Errorf("expected alex, got: %q", realValue)
	}

	if closeConn {
		t.Errorf("RENAME should not close the connection")
	}
}

func TestProcessRenameMissingArgs(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"RENAME"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: expected 3 arguments\n" {
		t.Errorf("expected ERROR: expected 3 arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("RENAME should not close the connection")
	}
}

func TestProcessMGet(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = Process([]string{"SET", "age", "20"}, st)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = Process([]string{"SET", "city", "buenos aires"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"MGET", "name", "age", "city"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "alex\n20\nbuenos aires\n" {
		t.Errorf("expected alex 20 buenos aires, got: %q", response)
	}

	if closeConn {
		t.Errorf("MGET should not close the connection")
	}
}

func TestProcessMGetMissingArgs(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"MGET"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: missing arguments\n" {
		t.Errorf("expected ERROR: missing arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("MGET should not close the connection")
	}
}

func TestProcessStrLen(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"STRLEN", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "4\n" {
		t.Errorf("expected 4, got: %q", response)
	}

	if closeConn {
		t.Errorf("STRLEN should not close the connection")
	}
}

func TestProcessStrLenMissingArgs(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"STRLEN"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: expected 2 arguments\n" {
		t.Errorf("expected ERROR: expected 2 arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("STRLEN should not close the connection")
	}
}

func TestProcessExpire(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"EXPIRE", "name", "30"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "1\n" {
		t.Errorf("expected 1, got: %q", response)
	}

	if closeConn {
		t.Errorf("EXPIRE should not close the connection")
	}
}

func TestProcessExpireMissingArgs(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"EXPIRE"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: expected 3 arguments\n" {
		t.Errorf("expected ERROR: expected 3 arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("EXPIRE should not close the connection")
	}
}

func TestProcessTTL(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = Process([]string{"EXPIRE", "name", "30"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"TTL", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "29\n" {
		t.Errorf("expected 30, got: %q", response)
	}

	if closeConn {
		t.Errorf("TTL should not close the connection")
	}
}

func TestProcessTTLMissingArgs(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"TTL"} ,st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: missing arguments\n" {
		t.Errorf("expected ERROR: missing arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("TTL should not close the connection")
	}
}

func TestProcessTTLNonExisting(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"TTL", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "-1\n" {
		t.Errorf("expected -1, got: %q", response)
	}

	if closeConn {
		t.Errorf("TTL should not close the connection")
	}
}

func TestProcessMSet(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"MSET", "name", "alex", "age", "20"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "OK\n" {
		t.Errorf("expected OK, got: %q", response)
	}

	mGetResponse, _, err := Process([]string{"MGET", "name", "age"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if mGetResponse != "alex\n20\n" {
		t.Errorf("expected alex 20, got: %q", response)
	}

	if closeConn {
		t.Errorf("MSET should not close the connection")
	}
}

func TestProcessMSetInvalid(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"MSET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: arguments must be key value pairs\n" {
		t.Errorf("expected ERROR: arguments must be key value pairs, got: %q", response)
	}

	if closeConn {
		t.Errorf("MSET should not close the connection")
	}
}

func TestProcessGetSet(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"GETSET", "name", "fran"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "alex\n" {
		t.Errorf("expected alex, got: %q", response)
	}

	getResponse, _, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if getResponse != "fran\n" {
		t.Errorf("expected fran, got: %q", getResponse)
	}

	if closeConn {
		t.Errorf("GETSET should not close the connection")
	}
}

func TestProcessGetSetMissingArgs(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"GETSET"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: missing arguments\n" {
		t.Errorf("expected ERROR: missing arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("GETSET should not close the connection")
	}
}

func TestProcessGetDel(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"GETDEL", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "alex\n" {
		t.Errorf("expected alex, got: %q", response)
	}

	getResponse, _, err := Process([]string{"GET", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if getResponse != "(nil)\n" {
		t.Errorf("expected (nil), got: %q", getResponse)
	}

	if closeConn {
		t.Errorf("GETDEL should not close the connection")
	}
}

func TestProcessGetDelMissingArgs(t *testing.T) {
	st := store.NewStore()

	response, closeConn, err := Process([]string{"GETDEL"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "ERROR: expected 2 arguments\n" {
		t.Errorf("expected ERROR: expected 2 arguments, got: %q", response)
	}

	if closeConn {
		t.Errorf("GETDEL should not close the connection")
	}
}

func TestProcessPersist(t *testing.T) {
	st := store.NewStore()

	_, _, err := Process([]string{"SET", "name", "alex"}, st)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = Process([]string{"EXPIRE", "name", "30"}, st)
	if err != nil {
		t.Fatal(err)
	}

	response, closeConn, err := Process([]string{"PERSIST", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if response != "1\n" {
		t.Errorf("expected 1, got: %q", response)
	}

	ttlResponse, _, err := Process([]string{"TTL", "name"}, st)
	if err != nil {
		t.Fatal(err)
	}

	if ttlResponse != "-1\n" {
		t.Fatalf("expected -1, got: %q", ttlResponse)
	}

	if closeConn {
		t.Errorf("PERSIST should not close the connection")
	}
}

func TestProcessPersistMissingArgs(t *testing.T) {}
