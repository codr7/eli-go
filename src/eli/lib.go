package eli

type Lib interface {
	Name() string
}

type BasicLib struct {
	name string
}

func (l *BasicLib) Init(name string) {
	l.name = name
}
