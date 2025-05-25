package eli

type Lib interface {
	Bind(string, Value)
	BindType(Type)
	Name() string
}

type BaseLib struct {
	name string
	bindings map[string]*Value
}

func (self *BaseLib) Init(name string) {
	self.name = name
}

func (self *BaseLib) Bind(k string, v Value) {
	self.bindings[k] = &v
}
