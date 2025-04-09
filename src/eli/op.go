package eli

type Opc func () error

type Op interface {
	Compile(vm *VM, pc PC) Opc
}
