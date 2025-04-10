package eli

import (
)

type PC = int
type Register = int

type VM struct {
	Debug bool
	Registers Deque[Value]
	Stack Stack
	StopPC PC
	
	ops Deque[Op]
	opcs Deque[Opc]
}

func (vm *VM) Init() {
	vm.Debug = true
	vm.StopPC = -1
}

func (vm *VM) Alloc(n int) Register {
	result := vm.Registers.Len()

	for i := 0; i < n; i++ {
		vm.Registers.PushLast(Value{})
	}

	return result
}

func (vm *VM) Compile(pc PC) {
	for i := pc; i < vm.ops.Len(); i++ {
		vm.opcs.PushLast(vm.ops.Get(i).Compile(vm, i))
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

func (vm *VM) Eval(fromPC, toPC PC) error {
	if fromPC == vm.ops.Len() {
		return nil
	}

	if vm.opcs.Len() < vm.ops.Len() {
		vm.Compile(vm.opcs.Len())
	}

	prevStopPC := vm.StopPC
	defer func() { vm.StopPC = prevStopPC }()
	vm.StopPC = toPC

	return vm.opcs.Get(fromPC)()
}

func (vm *VM) Opc(pc PC) Opc {
	result := func() error { return nil }

	if pc < vm.EmitPC() {
		result = vm.ops.Get(pc).Compile(vm, pc)
	}

	return result
}
