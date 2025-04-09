package core

import "eli/src/eli"

type Lib struct {
	eli.BasicLib
}

var nilValue eli.Value

func NIL() eli.Value {
	return nilValue
}

func init() {
	nilValue.Init(&Nil, nil)
}
