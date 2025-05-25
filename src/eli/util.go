package eli

import (
	"bufio"
	"unique"
)

func (self Stack[Value]) Dump(out *bufio.Writer) {
	out.WriteRune('[')
	out.WriteRune(']')	
}

type Sym = unique.Handle[string]

func S(name string) Sym {
	return unique.Make(name)
}
