package eli

import (
	"unique"
)

func Intern(name string) Sym {
	return unique.Make(name)
}
