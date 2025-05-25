package eli

import (
	"unique"
)

func S(name string) Sym {
	return unique.Make(name)
}
