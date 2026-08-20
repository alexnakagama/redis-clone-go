package resp

import (
	"fmt"
	"io"
)

type Encoder struct {
	writer io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		writer: w,
	}
}

func (e *Encoder) EncodeSimpleString(value string) error {
	_, err := fmt.Fprintf(e.writer, "+%s\r\n", value)
	return err	
}

func (e *Encoder) EncodeError(value string) error {
}
