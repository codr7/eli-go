package eli

type Deque[T any] struct {
	items []T
}

func NewDeque[T any](items...T) *Deque[T] {
	return new(Deque[T]).Init(items...)
}

func (d *Deque[T]) Init(items...T) *Deque[T] {
	d.items = items
	return d
}

func (d Deque[T]) Get(i int) T {
	return d.items[i]
}

func (d Deque[T]) Items() []T {
	return d.items
}

func (d Deque[T]) Len() int {
	return len(d.items)
}

func (d Deque[T]) PeekFirst() T {
	return d.items[0]
}

func (d Deque[T]) PeekLast() *T {
	return &d.items[len(d.items)-1]
}

func (d *Deque[T]) PopFirst() T {
	it := d.items[0]
	d.items = d.items[1:]
	return it
}

func (d *Deque[T]) PopLast() T {
	i := len(d.items) - 1
	it := d.items[i]
	d.items = d.items[:i]
	return it
}

func (d *Deque[T]) PushFirst(it T) {
	d.items = append(d.items, it)
	copy(d.items[:len(d.items)-1], d.items[1:])
}

func (d *Deque[T]) PushLast(it T) {
	d.items = append(d.items, it)
}

func (d *Deque[T]) Put(i int, it T) {
	d.items[i] = it
}
