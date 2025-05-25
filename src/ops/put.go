package ops

import "eli/src/eli"

type TPut struct {
	target eli.Register
	value  eli.Value
}

func Put(target eli.Register, value eli.Value) *TPut {
	return &TPut{target: target, value: value}
}

func (self *TPut) Compile(vm *eli.VM, pc eli.PC) eli.OpEval {
	return func () (eli.PC, error) {
		vm.Registers.Items[self.target] = self.value.Dup(vm)
		return pc+1, nil
	}
}
