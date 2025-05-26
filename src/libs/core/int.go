package core

import "eli/src/eli"

type TInt struct {
	eli.BaseType[int]
}

var Int TInt

func init() {
	Int.Init(eli.S("Int"))
}
