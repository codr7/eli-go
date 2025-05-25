package eli

type Type interface {
	Dup(Value, *VM) Value
	Name() string
}

type DataType[T any] interface {
	Type
	Zero() T
}

type BaseType[T any] struct {
	name string
}

func (t *BaseType[T]) Init(name string) {
	t.name = name
}

func (t *BaseType[T]) Dup(v Value, vm *VM) Value {
	return v
}

func (t *BaseType[T]) Name() string {
	return t.name
}

func (t *BaseType[T]) Zero() T {
	var v T
	return v
}
