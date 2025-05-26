package core

import "eli/src/eli"

type TNil struct {
	eli.BaseType[any]
}

var Nil TNil
var NIL eli.Value

func init() {
	Nil.Init(eli.S("Nil"))
	NIL.Init(&Nil, nil)
}
