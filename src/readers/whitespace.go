package readers

import (
	"bufio"
	"eli/src/eli"
	"io"
	"unicode"
)

type TWhitespace struct {
}

var Whitespace TWhitespace

func (self TWhitespace) Read(vm *eli.VM, in *bufio.Reader, out *eli.Deque[eli.Form], sloc *eli.Sloc) (bool, error) {
	found := false
	
	for {
		c, _, err := in.ReadRune()
		
		if err != nil {
			if err == io.EOF {
				err = nil
			}

			return false, err
		}

		if unicode.IsSpace(c) {
			sloc.Step(c)
			found = true
		} else {
			in.UnreadRune()
			break
		}
	}

	return found, nil
}
