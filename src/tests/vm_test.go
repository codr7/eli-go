package tests

import (
	"testing"

	"eli/src/eli"
	"eli/src/libs/core"
	"eli/src/ops"
)

func TestGetValue(t *testing.T) {
	var vm eli.VM
	vm.Init()

	r := vm.Alloc(1)
	pc := vm.EmitPC()
	vm.Emit(ops.Get(r))
	vm.Registers.Put(r, eli.V(&core.Int, 42))
	vm.Eval(pc)

	if v := vm.Stack.Pop().Data(); v != 42 {
		t.Fatalf("Expected 42, actual %v", v)
	}
}

func TestPushValue(t *testing.T) {
	var vm eli.VM
	vm.Init()

	pc := vm.EmitPC()
	vm.Emit(ops.PushValue(eli.V(&core.Int, 42))) 
	vm.Eval(pc)

	if v := vm.Stack.PopLast().Data(); v != 42 {
		t.Fatalf("Expected 42, actual %v", v)
	}
}

func TestPutValue(t *testing.T) {
	var vm eli.VM
	vm.Init()

	r := vm.Alloc(1)
	pc := vm.EmitPC()
	vm.Emit(ops.PutValue(r, eli.V(&core.Int, 42))) 
	vm.Eval(pc)

	if v := vm.Registers.Get(r).Data(); v != 42 {
		t.Fatalf("Expected 42, actual %v", v)
	}
}
