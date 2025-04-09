package core

type Type interface {
	Name() string
}

type DataType[T any] interface {
	Type
}

type BasicType struct {
	name string
}

func (t *BasicType) Name() string {
	return t.name
}
