package eli

import (
)

type PC = int
type Register = int

type VM struct {
	Debug bool
	Registers Deque[Value]
	Stack Stack
	Stop PC
	
	ops Deque[Op]
	opcs Deque[Opc]
}

func (vm *VM) Init() {
	vm.Debug = true
	vm.Stop = -1
}

func (vm *VM) Alloc(n int) Register {
	result := vm.Registers.Len()

	for i := 0; i < n; i++ {
		vm.Registers.PushLast(Value{})
	}

	return result
}

func (vm *VM) Compile(from PC) {
	for pc := from; pc < vm.ops.Len(); pc++ {
		vm.opcs.PushLast(vm.Opc(pc))
	}
}

func (vm *VM) Emit(op Op) int {
	result := vm.ops.Len()
	vm.ops.PushLast(op)
	return result
}

func (vm *VM) EmitPC() PC {
	return vm.ops.Len()
}

func (vm *VM) Eval(from, to PC) error {
	if from == vm.ops.Len() {
		return nil
	}

	if vm.opcs.Len() < vm.ops.Len() {
		vm.Compile(vm.opcs.Len())
	}

	prevStop := vm.Stop
	defer func() { vm.Stop = prevStop }()
	vm.Stop = to

	return vm.opcs.Get(from)()
}

func (vm *VM) Opc(pc PC) Opc {
	result := func() error { return nil }

	if pc < vm.EmitPC() {
		result = vm.ops.Get(pc).Compile(vm, pc)
	}

	return result
}
