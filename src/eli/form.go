package eli

import (
	"bufio"
)

type Form interface {
	SLoc() SLoc
	Emit(in *Deque[Form], vm *VM) error
	Quote(vm *VM) Value
	Dump(out *bufio.Writer) error
}

type BaseForm struct {
	sloc SLoc
}

func (self *BaseForm) Init(sloc SLoc) {
	self.sloc = sloc
}

func (self BaseForm) SLoc() SLoc {
	return self.sloc
}
