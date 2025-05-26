package core

import (
	"bufio"
	"eli/src/eli"
)

type TBool struct {
	eli.BaseType[bool]
}

var Bool TBool

func init() {
	Bool.Init(eli.S("Bool"))
}

func (self TBool) Dump(v eli.Value, out *bufio.Writer, vm *eli.VM) error {
	if v.Data.(bool) {
		_, err := out.WriteRune('T')
		return err
	}
	
	_, err := out.WriteRune('F')
	return err
}

func (self TBool) Write(v eli.Value, out *bufio.Writer, vm *eli.VM) error {
	return self.Dump(v, out, vm)
}


