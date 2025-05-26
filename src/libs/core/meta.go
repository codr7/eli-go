package core

import "eli/src/eli"

type TMeta struct {
	eli.BaseType[eli.Type]
}

var Meta TMeta

func init() {
	Meta.Init(eli.S("Meta"))
}
