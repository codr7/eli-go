package ops

import "eli/src/eli"

type OSet struct {
	target eli.Register
	value  eli.Value
}

func Set(target eli.Register, value eli.Value) *OSet {
	return &OSet{target: target, value: value}
}

func (o *OSet) Compile(vm *eli.VM, pc eli.PC) eli.Opc {
	next := func() error { return nil }
	pc++
	
	if vm.EmitPC() > pc {
		next = vm.Op(pc).Compile(vm, pc)
	}

	return func () error {
		vm.Registers.Set(o.target, o.value.Dup())
		return next()
	}
}
