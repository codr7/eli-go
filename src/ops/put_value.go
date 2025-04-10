package ops

import "eli/src/eli"

type TPutValue struct {
	target eli.Register
	value  eli.Value
}

func PutValue(target eli.Register, value eli.Value) *TPutValue {
	return &TPutValue{target: target, value: value}
}

func (op *TPutValue) Compile(vm *eli.VM, pc eli.PC) eli.Opc {
	next := vm.Opc(pc+1)

	return func () error {
		vm.Registers.Put(op.target, op.value.Dup())
		return next()
	}
}
