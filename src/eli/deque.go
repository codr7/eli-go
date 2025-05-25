package eli

type Deque[T any] struct {
	Items []T
}

func NewDeque[T any](items...T) *Deque[T] {
	return new(Deque[T]).Init(items...)
}

func (self *Deque[T]) Init(items...T) *Deque[T] {
	self.Items = items
	return self
}

func (self Deque[T]) Get(i int) T {
	return self.Items[i]
}

func (self Deque[T]) Len() int {
	return len(self.Items)
}

func (self Deque[T]) PeekFirst() T {
	return self.Items[0]
}

func (self Deque[T]) PeekLast() *T {
	return &self.Items[len(self.Items)-1]
}

func (self *Deque[T]) PopFirst() T {
	it := self.Items[0]
	self.Items = self.Items[1:]
	return it
}

func (self *Deque[T]) PopLast() T {
	i := len(self.Items) - 1
	it := self.Items[i]
	self.Items = self.Items[:i]
	return it
}

func (self *Deque[T]) PushFirst(it T) {
	self.Items = append(self.Items, it)
	copy(self.Items[:len(self.Items)-1], self.Items[1:])
}

func (self *Deque[T]) PushLast(it T) {
	self.Items = append(self.Items, it)
}

func (self *Deque[T]) Put(i int, it T) {
	self.Items[i] = it
}
