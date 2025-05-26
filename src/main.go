package main

import (
	"eli/src/eli"
	"eli/src/readers"
	"eli/src/tools"
	//"fmt"
)

func main() {
	var vm eli.VM
	vm.Init(&readers.Form)
	tools.Repl(&vm)
}
