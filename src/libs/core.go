package libs

import (
	"eli/src/eli"
	"eli/src/libs/core"
)

type TCore struct {
	eli.BaseLib
}

var Core TCore

func init() {
	Core.Init("core")
}

func (self *TCore) Init(name string) {
	self.BaseLib.Init(name, nil)
	self.BindType(&core.Bool)
	self.BindType(&core.Int)
	self.BindType(&core.Meta)
	self.BindType(&core.Nil)
	self.BindType(&core.Sym)

	self.Bind(eli.S("T"), eli.V(&core.Bool, true))
	self.Bind(eli.S("F"), eli.V(&core.Bool, false))
}

func (self *TCore) BindType(t eli.Type) {
	self.Bind(t.Name(), eli.V(&core.Meta, t))
}
