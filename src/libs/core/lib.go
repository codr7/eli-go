package core

import "eli/src/eli"

type Lib struct {
	eli.BaseLib
}

func (self *Lib) Init(name eli.Sym) {
	self.BaseLib.Init(name)
	self.BindType(&Int)
	self.BindType(&Meta)
	self.BindType(&Nil)
	self.BindType(&Sym)
}

func (self *Lib) BindType(t eli.Type) {
	self.Bind(t.Name(), eli.V(&Meta, t))
}
