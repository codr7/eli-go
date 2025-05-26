package forms

import (
	"eli/src/eli"
)

func Emit(in *eli.Deque[eli.Form], vm *eli.VM) error {
	for in.Len() > 0 {
		if err := in.PopFront().Emit(in, vm); err != nil {
			return err
		}
	}

	return nil
}
