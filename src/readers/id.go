package readers

import (
	"bufio"
	//"fmt"
	"eli/src/eli"
	"eli/src/forms"
	"io"
	"strings"
	"unicode"
)

type TId struct {
}

var Id TId

func (self TId) Read(vm *eli.VM, in *bufio.Reader, out *eli.Deque[eli.Form], sloc *eli.Sloc) (bool, error) {
	formSloc := *sloc
	var buf strings.Builder

	for {
		c, _, err := in.ReadRune()
		
		if err != nil {
			if err == io.EOF {
				err = nil
			}

			return false, err
		}

		if unicode.IsSpace(c) ||
			c == '(' || c == ')' {
			in.UnreadRune()
			break
		}

		buf.WriteRune(c)
		sloc.Step(c)
	}

	if buf.Len() == 0 {
		return false, nil
	}
	
	out.PushBack(forms.NewId(formSloc, buf.String()))
	return true, nil
}
