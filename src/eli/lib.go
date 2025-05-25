package eli

type Lib interface {
	Bind(Sym, Value)
	BindType(Type)
	Name() Sym
}

type BaseLib struct {
	name Sym
	bindings map[Sym]*Value
}

func (self *BaseLib) Init(name Sym) {
	self.name = name
}

func (self *BaseLib) Bind(k Sym, v Value) {
	self.bindings[k] = &v
}
