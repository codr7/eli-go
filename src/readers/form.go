package readers

import (
	"bufio"
	"eli/src/eli"
	"io"
)

type TForm struct {
}

var Form TForm

func (self TForm) Read(vm *eli.VM, in *bufio.Reader, out *eli.Deque[eli.Form], sloc *eli.Sloc) (bool, error) {
	if _, err := Whitespace.Read(vm, in, out, sloc); err != nil {
		return false, err
	}

	c, _, err := in.ReadRune()

	if err != nil {
		if err == io.EOF {
			err = nil
		}

		return false, err
	}

	in.UnreadRune()
	
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return false, eli.NewReadError(*sloc, "Not implemented")
	case '(':
		return false, eli.NewReadError(*sloc, "Not implemented")
	default:
		return Id.Read(vm, in, out, sloc)
	}

	return false, eli.NewReadError(*sloc, "Invalid syntax: %v", c)
}
