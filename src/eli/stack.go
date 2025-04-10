package eli

type Stack struct {
	Deque[Value]
}

func (s Stack) Peek() *Value {
	return s.PeekLast()
}

func (s *Stack) Pop() Value {
	return s.PopLast()
}

func (s *Stack) Push(it Value) {
	s.PushLast(it)
}

