package resp

import (
	"bytes"
	"testing"
)

func TestEncodeSimpleString(t *testing.T) {
	var buf bytes.Buffer

	encoder := NewEncoder(&buf)

	err := encoder.EncodeSimpleString("OK")
	if err != nil {
		t.Fatal(err)
	}

	expected := "+OK\r\n"

	if buf.String() != expected {
		t.Errorf("expected %q, got: %q", expected, buf.String())
	}
}

func TestEncodeError(t *testing.T) {
	var buf bytes.Buffer

	encoder := NewEncoder(&buf)

	err := encoder.EncodeError("-ERR unknown command")
	if err != nil {
		t.Fatal(err)
	}

	expected := "--ERR unknown command\r\n"

	if buf.String() != expected {
		t.Errorf("expected %q, got: %q", expected, buf.String())
	}
}

func TestEncodeInteger(t *testing.T) {
	var buf bytes.Buffer

	encoder := NewEncoder(&buf)

	err := encoder.EncodeInteger(42)
	if err != nil {
		t.Fatal(err)
	}

	expected := ":42\r\n"

	if buf.String() != expected {
		t.Errorf("expected %q, got: %q", expected, buf.String())
	}
}

func TestEncodeBulkString(t *testing.T) {
	var buf bytes.Buffer

	encoder := NewEncoder(&buf)

	err := encoder.EncodeBulkString("alex")
	if err != nil {
		t.Fatal(err)
	}

	expected := "$4\r\nalex\r\n" 

	if buf.String() != expected {
		t.Errorf("expected %q, got: %q", expected, buf.String())
	}
}

func TestEncodeNull(t *testing.T) {
	var buf bytes.Buffer

	encoder := NewEncoder(&buf)

	err := encoder.EncodeNull()
	if err != nil {
		t.Fatal(err)
	}

	expected := "$-1\r\n"

	if buf.String() != expected {
		t.Errorf("expected %q, got: %q", expected, buf.String())
	}
}
