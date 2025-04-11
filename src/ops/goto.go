package ops

import "eli/src/eli"

type TGoto struct {
	pc eli.Register
}

func Goto(pc eli.PC) *TGoto {
	return &TGoto{pc: pc}
}

func (op *TGoto) Compile(vm *eli.VM, pc eli.PC) eli.Opc {
	next := vm.Opc(op.pc)
	
	return func() error {
		if pc == vm.Stop {
			return nil
		}
		
		return next()
	}
}
