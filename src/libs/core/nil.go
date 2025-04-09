package core

import "eli/src/eli"

type NilType struct {
	eli.BasicType[any]
}

var Nil NilType

func init() {
	Nil.Init("Nil")
}
