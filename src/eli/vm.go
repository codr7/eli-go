package eli

import (
	"bufio"
)

type PC = int
type Register = int
type Reader = func (*VM, *bufio.Reader, *Deque[Form], *Sloc) (bool, error)
	
type VM struct {
	Debug bool
	
	Reader Reader
	Registers Stack[Value]
	Stack Stack[Value]

	currentLib Lib
	userLib BaseLib

	ops Stack[Op]
	opEvals Stack[OpEval]
}

func (self *VM) Init() {
	self.Debug = true
	self.userLib.Init("user")
	self.currentLib = &self.userLib
}

func (self *VM) Alloc(n int) Register {
	result := self.Registers.Len()

	for i := 0; i < n; i++ {
		self.Registers.Push(Value{})
	}

	return result
}

func (self *VM) Compile(from PC) {
	for pc := from; pc < self.ops.Len(); pc++ {
		self.opEvals.Push(self.ops.Items[pc].Compile(self, pc))
	}
}

func (self *VM) CurrentLib() Lib {
	return self.currentLib
}

func (self *VM) Emit(op Op) int {
	result := self.ops.Len()
	self.ops.Push(op)
	return result
}

func (self VM) EmitPC() PC {
	return self.ops.Len()
}

func (self *VM) Eval(from, to PC) error {
	if to == -1 {
		to = self.ops.Len()
	}

	if self.opEvals.Len() < to {
		self.Compile(self.opEvals.Len())
	}

	var err error;
	
	for pc := from;
	err == nil && pc < to;
	pc, err = self.opEvals.Items[pc]() {
		//Do nothing
	}

	return err
}

func (self *VM) Read(in *bufio.Reader, out *Deque[Form], sloc *Sloc) error {
	for {
		ok, err := self.Reader(self, in, out, sloc)

		if err != nil {
			return err
		}

		if !ok {
			break;
		}
	}

	return nil
}
