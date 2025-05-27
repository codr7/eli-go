package readers

import (
	"bufio"
	//"fmt"
	"eli/src/eli"
	"eli/src/forms"
	"io"
)

type TScope struct {
}

var Scope TScope

func (self TScope) Read(vm *eli.VM, in *bufio.Reader, out *eli.Deque[eli.Form], sloc *eli.Sloc) (bool, error) {
	formSloc := *sloc
	
	if c, _, err := in.ReadRune(); err != nil {
		if err == io.EOF {
			err = nil
		}

		return false, err
	} else if c != '(' {
		in.UnreadRune()
		return false, nil
	} else {
		sloc.Step(c)
	}

	var buf eli.Deque[eli.Form]
	
	for {
		if _, err := Whitespace.Read(vm, in, &buf, sloc); err != nil {
			return false, err
		}
		
		if c, _, err := in.ReadRune(); err != nil {
			if err == io.EOF {
				err = eli.NewReadError(*sloc, "Invalid syntax")
			}
			
			return false, err
		} else if c == ')' {
			sloc.Step(c)
			break
		}

		in.UnreadRune()

		if ok, err := Form.Read(vm, in, &buf, sloc); err != nil {
			return false, err
		} else if !ok {
			err = eli.NewReadError(*sloc, "Invalid syntax")
		}
	}

	out.PushBack(forms.NewScope(formSloc, buf))
	return true, nil
}
