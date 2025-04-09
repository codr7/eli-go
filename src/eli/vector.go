package eli

type Vector[T any] struct {
	items []T
}

func (v *Vector[T]) Get(i int) T {
	return v.items[i]
}

func (v *Vector[T]) Items() []T {
	return v.items
}

func (v *Vector[T]) Len() int {
	return len(v.items)
}

func (v *Vector[T]) Push(it T) {
	v.items = append(v.items, it)
}

func (v *Vector[T]) Set(i int, it T) {
	v.items[i] = it
}
