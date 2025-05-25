package eli

import (
	"fmt"
)

type SLoc struct {
	source       string
	line, column int
}

func NewSLoc(source string) *SLoc {
	return new(SLoc).Init(source, 1, 0)
}

func (self *SLoc) Init(source string, line, column int) *SLoc {
	self.source = source
	self.line = line
	self.column = column
	return self
}

func (self SLoc) String() string {
	return fmt.Sprintf("'%v' at line %v, column %v",
		self.source, self.line, self.column)
}
