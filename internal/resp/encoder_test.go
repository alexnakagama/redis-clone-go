package resp

import (
	"bytes"
	"testing"
)

func TestEncodeSimpleString(t *testing.T) {
	var buf bytes.Buffer

	err := NewEncoder(&buf)
	if err != nil {
		t.Fatal(err)
	}

	expected := "OK\r\n"

	if buf.String() != expected {
		t.Errorf("expected %q, got: %q", expected, buf.String())
	}
}
