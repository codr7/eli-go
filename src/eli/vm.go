package eli

type PC = int
type Register = int

type VM struct {
	Debug bool
	Registers Vector[Value]
	Stack Stack
	
	ops Vector[Op]
	opcs Vector[Opc]
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

func (vm *VM) Compile(pc PC) {
	for i := pc; i < vm.ops.Len(); i++ {
		vm.opcs.Push(vm.ops.Get(i).Compile(vm, i))
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

func (vm *VM) Eval(pc PC) error {
	if vm.opcs.Len() < vm.ops.Len() {
		vm.Compile(vm.opcs.Len())
	}
	
	if pc >= vm.opcs.Len() {
		return nil
	}
	
	return vm.opcs.Get(pc)()
}

func (vm *VM) Op(pc PC) Op {
	return vm.ops.Get(pc)
}
