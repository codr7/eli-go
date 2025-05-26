package eli

import (
	"fmt"
)

type Error struct {
	message string
}

func (self *Error) Init(spec string, args...any) {
	self.message = fmt.Sprintf(spec, args...)
}

func (self Error) Error() string {
	return self.message
}

type EmitError struct {
	Error
}

func NewEmitError(sloc Sloc, spec string, args...any) EmitError {
	var e EmitError
	e.Init("Emit Error in %v: " + spec, append([]any{sloc}, args...)...)
	return e
}
