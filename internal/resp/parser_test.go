package resp

import (
	"testing"
	"bufio"
	"reflect"
	"strings"
)

func TestParse(t *testing.T) {
	input := "*2\r\n" +
		"$3\r\n" +
		"GET\r\n" +
		"$4\r\n" +
		"name\r\n"

	reader := bufio.NewReader(strings.NewReader(input))

	result, err := Parse(reader)

	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"GET", "name"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
