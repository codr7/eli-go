package forms

import (
	"io"
)

type Id struct {
	eli.BaseForm
	name eli.Sym
}

func NewId(sloc eli.Sloc, name string) *Id {
	return new(Id).Init(pos, eli.S(name))
}

func (self *Id) Init(sloc eli.Sloc, name string) *Id {
	self.BaseForm.Init(sloc)
	self.name = name
	return self
}

func (self *Id) Emit(in *Deque[Form], vm *VM) error {
	v := vm.CurrentLib.Find(self.name)

	if v == nil {
		return eli.NewEmitError(self.Sloc(),
			"Unknown identifier: %v",
			self.name)
	}

	vm.Stack.Push(v)
	return nil
}

func (self Id) Quote(vm *VM) Value {
	return eli.V(&core.Sym, name)
}

func (self Id) Dump(out *bufio.Writer) error {
	out.WriteString(self.name)
}
