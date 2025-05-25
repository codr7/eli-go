package main

import (
	"eli/src/eli"
	"eli/src/tools"
	//"fmt"
)

func main() {
	var vm eli.VM
	vm.Init()
	tools.Repl(&vm)
}
