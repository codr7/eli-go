package eli

type Stack[T any] struct {
	Deque[T]
}

func (self Stack[T]) Peek() *T {
	return self.PeekLast()
}

func (self *Stack[T]) Pop() T {
	return self.PopLast()
}

func (self *Stack[T]) Push(it T) {
	self.PushLast(it)
}

