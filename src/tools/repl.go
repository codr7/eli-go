package tools

import (
	"bufio"
	"bytes"
	"eli/src/eli"
	"eli/src/forms"
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
			sloc := eli.NewSloc("repl")
			pc := vm.EmitPC()
			var fs eli.Deque[eli.Form]

			if err := vm.Read(bufio.NewReader(&b), &fs, sloc);
			err != nil {
				fmt.Println(err)
				b.Reset()
				goto NEXT
			}

			b.Reset()

			if err := forms.Emit(&fs, vm); err != nil {
				fmt.Println(err)
				goto NEXT
			}

			if err := vm.Eval(pc, -1); err != nil {
				fmt.Println(err)
				goto NEXT
			}
		NEXT:
			eli.DumpStack(vm.Stack, out, vm)
			out.WriteRune('\n')
		} else {
			fmt.Fprintln(&b, line)
		}
	}
}
