package ops

import "eli/src/eli"

type TPushValue struct {
	value  eli.Value
}

func PushValue(value eli.Value) *TPushValue {
	return &TPushValue{value: value}
}

func (op *TPushValue) Compile(vm *eli.VM, pc eli.PC) eli.Opc {
	next := func() error { return nil }
	pc++
	
	if vm.EmitPC() > pc {
		next = vm.Op(pc).Compile(vm, pc)
	}

	return func () error {
		vm.Stack.Push(op.value.Dup())
		return next()
	}
}
