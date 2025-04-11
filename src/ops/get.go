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

func (op *TGet) Compile(vm *eli.VM, pc eli.PC) eli.Opc {
	next := vm.Opc(pc+1)
	
	return func () error {
		if pc == vm.Stop {
			return nil
		}
		
		vm.Stack.Push(vm.Registers.Get(op.source))
		return next()
	}
}
