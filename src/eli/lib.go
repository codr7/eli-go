package eli

type Lib interface {
	Bind(string, Value)
	BindType(Type)
	Name() string
}

type BasicLib struct {
	name string
	bindings map[string]*Value
}

func (self *BasicLib) Init(name string) {
	self.name = name
}

func (self *BasicLib) Bind(k string, v Value) {
	self.bindings[k] = &v
}
