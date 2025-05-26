package tests

import (
	"testing"

	"eli/src/eli"
	"eli/src/libs/core"
	"eli/src/ops"
	"eli/src/readers"
)

func newVM() *eli.VM {
	return new(eli.VM).Init(&readers.Form)
}

func TestGet(t *testing.T) {
	vm := newVM()

	r := vm.AllocateRegisters(1)
	vm.Registers.Items[r] = eli.V(&core.Int, 42)

	pc := vm.EmitPC()
	vm.Emit(ops.Get(r))
	vm.Eval(pc, -1)

	if v := vm.Stack.Pop().Data; v != 42 {
		t.Fatalf("Expected 42, actual %v", v)
	}
}

func TestPushValue(t *testing.T) {
	vm := newVM()

	pc := vm.EmitPC()
	vm.Emit(ops.Push(eli.V(&core.Int, 42))) 
	vm.Eval(pc, -1)

	if v := vm.Stack.Pop().Data; v != 42 {
		t.Fatalf("Expected 42, actual %v", v)
	}
}

func TestPutValue(t *testing.T) {
	vm := newVM()

	r := vm.AllocateRegisters(1)
	pc := vm.EmitPC()
	vm.Emit(ops.Put(r, eli.V(&core.Int, 42))) 
	vm.Eval(pc, -1)

	if v := vm.Registers.Items[r].Data; v != 42 {
		t.Fatalf("Expected 42, actual %v", v)
	}
}
