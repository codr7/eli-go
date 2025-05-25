package core

import "eli/src/eli"

type SymType struct {
	eli.BaseType[eli.Sym]
}

var Sym SymType

func init() {
	Sym.Init("Sym")
}
