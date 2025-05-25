package eli

import (
)

type PC = int
type Register = int

type VM struct {
	Debug bool
	Registers Stack[Value]
	Stack Stack[Value]
	
	ops Stack[Op]
	opEvals Stack[OpEval]
}

func (vm *VM) Init() {
	vm.Debug = true
}

func (vm *VM) Alloc(n int) Register {
	result := vm.Registers.Len()

	for i := 0; i < n; i++ {
		vm.Registers.Push(Value{})
	}

	return result
}

func (vm *VM) Compile(from PC) {
	for pc := from; pc < vm.ops.Len(); pc++ {
		vm.opEvals.Push(vm.ops.Items[pc].Compile(vm, pc))
	}
}

func (vm *VM) Emit(op Op) int {
	result := vm.ops.Len()
	vm.ops.Push(op)
	return result
}

func (vm *VM) EmitPC() PC {
	return vm.ops.Len()
}

func (vm *VM) Eval(from, to PC) error {
	if to == -1 {
		to = vm.ops.Len()
	}

	if vm.opEvals.Len() < to {
		vm.Compile(vm.opEvals.Len())
	}

	var err error;
	
	for pc := from;
	err == nil && pc < to;
	pc, err = vm.opEvals.Items[pc]() {
		//Do nothing
	}

	return err
}
