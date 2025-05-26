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
	case '0':
		break
	case '(':
		break
	default:
		return Id.Read(vm, in, out, sloc)
	}

	return false, eli.NewReadError(*sloc, "Invalid syntax: %v", c)
}
