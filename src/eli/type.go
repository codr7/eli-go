package eli

type Type interface {
	Name() string
}

type DataType[T any] interface {
	Type
	Zero() T
}

type BasicType[T any] struct {
	name string
}

func (t *BasicType[T]) Init(name string) {
	t.name = name
}

func (t *BasicType[T]) Name() string {
	return t.name
}

func (t *BasicType[T]) Zero() T {
	var v T
	return v
}
