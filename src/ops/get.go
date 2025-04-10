package ops

import "eli/src/eli"

type TGet struct {
	target eli.Register
	value  eli.Value
}

func Get(target eli.Register) *TGet {
	return &TGet{target: target}
}

func (op *TGet) Compile(vm *eli.VM, pc eli.PC) eli.Opc {
	next := func() error { return nil }
	pc++
	
	if vm.EmitPC() > pc {
		next = vm.Op(pc).Compile(vm, pc)
	}

	return func () error {
		vm.Stack.Push(vm.Registers.Get(op.target))
		return next()
	}
}
