package forms

import (
	"bufio"
	"eli/src/eli"
	"eli/src/libs/core"
)

type Id struct {
	eli.BaseForm
	name eli.Sym
}

func NewId(sloc eli.Sloc, name string) *Id {
	return new(Id).Init(sloc, eli.S(name))
}

func (self *Id) Init(sloc eli.Sloc, name eli.Sym) *Id {
	self.BaseForm.Init(sloc)
	self.name = name
	return self
}

func (self *Id) Emit(in *eli.Deque[eli.Form], vm *eli.VM) error {
	v := vm.CurrentLib().Find(self.name)

	if v == nil {
		return eli.NewEmitError(self.Sloc(),
			"Unknown identifier: %v",
			self.name)
	}

	vm.Stack.Push(*v)
	return nil
}

func (self Id) Quote(vm *eli.VM) eli.Value {
	return eli.V(&core.Sym, self.name)
}

func (self Id) Dump(out *bufio.Writer) error {
	_, err := out.WriteString(self.name.Value())
	return err
}
