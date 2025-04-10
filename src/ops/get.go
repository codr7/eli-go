package ops

import (
	"fmt"
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
		fmt.Printf("GET %v %v\n", vm.StopPC, pc)

		if vm.StopPC == pc {
			return nil
		}
		
		vm.Stack.Push(vm.Registers.Get(op.source))
		return next()
	}
}
