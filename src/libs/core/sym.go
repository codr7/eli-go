package core

import "eli/src/eli"

type TSym struct {
	eli.BaseType[eli.Sym]
}

var Sym TSym

func init() {
	Sym.Init(eli.S("Sym"))
}
