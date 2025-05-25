package core

import "eli/src/eli"

type IntType struct {
	eli.BaseType[int]
}

var Int IntType

func init() {
	Int.Init("Int")
}
