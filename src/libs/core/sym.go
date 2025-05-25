package core

import "eli/src/eli"

type SymType struct {
	eli.BasicType[eli.Sym]
}

var Sym SymType

func init() {
	Sym.Init("Sym")
}
