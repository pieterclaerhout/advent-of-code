package generics

import (
	"container/heap"
)

type Comparator[T any] func(T, T) bool

type Heap[T any] struct {
	data []T
	comp func(x, y T) bool
}

func NewHeap[T any](best Comparator[T]) *Heap[T] {
	h := &Heap[T]{
		data: []T{},
		comp: best,
	}
	heap.Init(h)
	return h
}

func (h Heap[T]) Len() int {
	return len(h.data)
}

func (h Heap[T]) Less(i, j int) bool {
	return h.comp(h.data[i], h.data[j])
}

func (h Heap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) Push(x any) {
	h.data = append(h.data, x.(T))
}

func (h *Heap[T]) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[:n-1]
	return x
}
