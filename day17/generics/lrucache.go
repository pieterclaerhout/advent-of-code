package generics

import (
	"container/heap"
)

type LruCache[K comparable, V any] struct {
	byAge    *Heap[*lruEntry[K, V]]
	byKey    map[K]*lruEntry[K, V]
	data     []lruEntry[K, V]
	capacity int
	epoch    lruEpoch
}

func NewLruCache[K comparable, V any](cap int) *LruCache[K, V] {
	return &LruCache[K, V]{
		byKey: map[K]*lruEntry[K, V]{},
		byAge: NewHeap(func(x, y *lruEntry[K, V]) bool {
			return x.epoch < y.epoch
		}),
		data:     make([]lruEntry[K, V], cap),
		capacity: cap,
		epoch:    1,
	}
}

func (lru LruCache[K, V]) Len() int {
	return len(lru.byKey)
}

func (lru *LruCache[K, V]) Get(key K) (v V, ok bool) {
	entry, ok := lru.byKey[key]
	if !ok {
		return v, false
	}
	lru.epoch++
	entry.epoch = lru.epoch
	return entry.data, true
}

func (lru LruCache[K, V]) Set(k K, v V) {
	lru.epoch++
	entry, ok := lru.byKey[k]
	if ok {
		entry.epoch = lru.epoch
		return
	}
	var ptr *lruEntry[K, V]
	if lru.Len() >= lru.capacity {
		ptr = heap.Pop(lru.byAge).(*lruEntry[K, V])
		delete(lru.byKey, ptr.key)
	} else {
		ptr = &lru.data[lru.Len()]
	}

	ptr.epoch = lru.epoch
	ptr.key = k
	ptr.data = v

	lru.byKey[k] = ptr
	heap.Push(lru.byAge, ptr)
}

type lruEpoch uint64

type lruEntry[K, V any] struct {
	epoch lruEpoch
	key   K
	data  V
}
