package eli

import (
	"bufio"
	"fmt"
)

type Type interface {
	Dump(Value, *bufio.Writer, *VM)
	Dup(Value, *VM) Value
	Name() Sym
}

type DataType[T any] interface {
	Type
	Zero() T
}

type BaseType[T any] struct {
	name Sym
}

func (t *BaseType[T]) Init(name Sym) {
	t.name = name
}

func (t *BaseType[T]) Dump(v Value, out *bufio.Writer, vm *VM) {
	fmt.Fprintf(out, "%v", v.Data)
}

func (t *BaseType[T]) Dup(v Value, vm *VM) Value {
	return v
}

func (t *BaseType[T]) Name() Sym {
	return t.name
}

func (t *BaseType[T]) Zero() T {
	var v T
	return v
}
