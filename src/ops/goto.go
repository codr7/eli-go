package ops

import "eli/src/eli"

type TGoto struct {
	pc eli.Register
}

func Goto(pc eli.PC) *TGoto {
	return &TGoto{pc: pc}
}

func (self *TGoto) Compile(vm *eli.VM, pc eli.PC) eli.OpEval {
	return func() (eli.PC, error) {
		return self.pc, nil
	}
}
