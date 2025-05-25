package core

import "eli/src/eli"

type NilType struct {
	eli.BaseType[any]
}

var Nil NilType
var NIL eli.Value

func init() {
	Nil.Init("Nil")
	NIL.Init(&Nil, nil)
}
