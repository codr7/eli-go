package ops

import (
	"eli/src/eli"
)

type TPush struct {
	value  eli.Value
}

func Push(value eli.Value) *TPush {
	return &TPush{value: value}
}

func (self *TPush) Compile(vm *eli.VM, pc eli.PC) eli.OpEval {
	return func () (eli.PC, error) {
		vm.Stack.Push(self.value.Dup(vm))
		return pc+1, nil
	}
}
