package eli

import (
	"cmp"
)

type BSetCompare[K cmp.Ordered] = func(x, y K) int
type BSetKey[K cmp.Ordered, V any] = func(V) K

type BSet[K cmp.Ordered, V any] struct {
	Deque[V]
	compare BSetCompare[K]
	key BSetKey[K, V]
}

func (self *BSet[K, V]) Init(c BSetCompare[K], k BSetKey[K, V]) {
	self.compare = c
	self.key = k
}
