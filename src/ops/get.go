package ops

import (
	"eli/src/eli"
)

type TGet struct {
	source eli.Register
}

func Get(source eli.Register) *TGet {
	return &TGet{source: source}
}

func (self *TGet) Compile(vm *eli.VM, pc eli.PC) eli.OpEval {
	return func () (eli.PC, error) {
		vm.Stack.Push(vm.Registers.Get(self.source))
		return pc+1, nil
	}
}
