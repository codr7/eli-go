package forms

import (
	"bufio"
	"eli/src/eli"
	"eli/src/libs/core"
)

type Scope struct {
	eli.BaseForm
	forms eli.Deque[eli.Form]
}

func NewScope(sloc eli.Sloc, forms eli.Deque[eli.Form]) *Scope {
	return new(Scope).Init(sloc, forms)
}

func (self *Scope) Init(sloc eli.Sloc, forms eli.Deque[eli.Form]) *Scope {
	self.BaseForm.Init(sloc)
	self.forms = forms
	return self
}

func (self *Scope) Emit(in *eli.Deque[eli.Form], vm *eli.VM) error {
	return EmitAll(self.forms, vm)
}

func (self Scope) Quote(vm *eli.VM) eli.Value {
	//TODO emit list
	return core.NIL
}

func (self Scope) Dump(out *bufio.Writer) error {
	if _, err := out.WriteRune('('); err != nil {
		return err
	}

	for i, f := range self.forms.Items {
		if i > 0 {
			if _, err := out.WriteRune(' '); err != nil {
				return err
			}
		}

		if err := f.Dump(out); err != nil {
			return err
		}
	}
	
	if _, err := out.WriteRune(')'); err != nil {
		return err
	}

	return nil
}
