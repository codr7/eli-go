package main

import "fmt"

import "eli/src/eli"
import "eli/src/ops"
import "eli/src/libs/core"

func main() {
     var vm eli.VM
     vm.Init()

     r := vm.Alloc(1)
     pc := vm.EmitPC()
     vm.Emit(ops.PutValue(r, eli.NewValue(&core.Int, 42))) 
     vm.Eval(pc)

     fmt.Printf("%v\n", vm.Registers.Get(r).Data())
}