package eli

type Lib interface {
	Bind(Sym, Value)
	BindType(Type)
	Find(k Sym) *Value
	Name() Sym
}

type BaseLib struct {
	name Sym
	bindings map[Sym]*Value
}

func (self *BaseLib) Init(name Sym) {
	self.name = name
	self.bindings = make(map[Sym]*Value)
}

func (self *BaseLib) Bind(k Sym, v Value) {
	self.bindings[k] = &v
}

func (self *BaseLib) Find(k Sym) *Value {
	return self.bindings[k]
}

func (self *BaseLib) Name() Sym {
	return self.name
}
