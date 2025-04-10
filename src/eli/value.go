package eli

type Value struct {
	etype Type
	data any
}

func V[D any](etype DataType[D], data D) Value {
	return new(Value).Init(etype, data)
}

func (v *Value) Init(etype Type, data any) Value {
	return Value{etype: etype, data: data}
}

func (v Value) Data() any {
	return v.data
}

func (v Value) Dup() Value {
	return v
}
