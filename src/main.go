package main

import (
	"eli/src/eli"
	"eli/src/libs"
	"eli/src/readers"
	"eli/src/tools"
	//"fmt"
)

func main() {
	var vm eli.VM
	vm.Init(&readers.Form)
	vm.CurrentLib().Import(&libs.Core)
	tools.Repl(&vm)
}
