package tools

import (
	"bufio"
	"bytes"
	"eli/src/eli"
	"fmt"
	"log"
	"os"
)

func Repl(vm *eli.VM) {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var b bytes.Buffer

	for {
		out.WriteString("  ")
		out.Flush()
		
		if !in.Scan() {
			if err := in.Err(); err != nil {
				log.Fatal(err)
			}

			break
		}

		line := in.Text()

		if line == "" {
			sloc := eli.NewSLoc("repl")
			pc := vm.EmitPC()
			var forms eli.Deque[eli.Form]

			if err := vm.Read(bufio.NewReader(&b), &forms, sloc);
			err != nil {
				fmt.Println(err)
				b.Reset()
				goto NEXT
			}

			b.Reset()

			/*if err := forms.Emit(vm, env); err != nil {
				fmt.Println(err)
				goto NEXT
			}*/

			if err := vm.Eval(pc, -1); err != nil {
				fmt.Println(err)
				goto NEXT
			}
		NEXT:
			vm.Stack.Dump(out)
			out.WriteRune('\n')
		} else {
			fmt.Fprintln(&b, line)
		}
	}
}
