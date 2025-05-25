package core

import "eli/src/eli"

type Lib struct {
	eli.BasicLib
}

func (self *Lib) Init(name string) {
	self.BasicLib.Init(name)
	self.BindType(&Int)
	self.BindType(&Meta)
	self.BindType(&Nil)
	self.BindType(&Sym)
}

func (self *Lib) BindType(t eli.Type) {
	self.Bind(t.Name(), eli.V(&Meta, t))
}
