package core

import "eli/src/eli"

type MetaType struct {
	eli.BaseType[eli.Type]
}

var Meta MetaType

func init() {
	Meta.Init("Meta")
}
