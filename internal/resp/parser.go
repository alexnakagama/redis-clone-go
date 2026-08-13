package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Parse(r *bufio.Reader) ([]string, error)
