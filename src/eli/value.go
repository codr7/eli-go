package eli

import (
	"fmt"
)

type Value struct {
	Type Type
	Data any
}

func V[T any](t DataType[T], d T) Value {
	return *new(Value).Init(t, d)
}

func (self *Value) Init(t Type, d any) *Value {
	self.Type = t
	self.Data = d
	return self
}

func (self Value) Dup(vm *VM) Value {
	return self.Type.Dup(self, vm)
}

func Cast[T any](v Value, t DataType[T]) (T, error) {
	if v.Type != t {
		var tv T
		return tv, fmt.Errorf("Expected %v: %v", v.Type, t)
	}
	
	return v.Data.(T), nil
}
