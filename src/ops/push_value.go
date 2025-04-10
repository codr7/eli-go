package ops

import "eli/src/eli"

type TPushValue struct {
	value  eli.Value
}

func PushValue(value eli.Value) *TPushValue {
	return &TPushValue{value: value}
}

func (op *TPushValue) Compile(vm *eli.VM, pc eli.PC) eli.Opc {
	next := vm.Opc(pc+1)

	return func () error {
		if vm.StopPC == pc {
			return nil
		}
		
		vm.Stack.Push(op.value.Dup())
		return next()
	}
}
