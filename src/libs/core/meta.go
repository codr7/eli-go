package core

import "eli/src/eli"

type MetaType struct {
	eli.BasicType[eli.Type]
}

var Meta MetaType

func init() {
	Meta.Init("Meta")
}
