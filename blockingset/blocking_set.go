package blockingset

import (
	"github.com/ngc4736/collections/set"
	"sync"
)

type BlockingSet[T comparable] struct {
	keys *set.Set[T]
	lock sync.RWMutex
}

func New[T comparable](keys ...T) *BlockingSet[T] {
	return &BlockingSet[T]{
		keys: set.New[T](keys...),
		lock: sync.RWMutex{},
	}
}

func (b *BlockingSet[T]) Size() int {
	b.lock.RLock()
	size := b.keys.Size()
	b.lock.RUnlock()
	return size
}

func (b *BlockingSet[T]) Put(key T) bool {
	b.lock.Lock()
	ok := b.keys.Put(key)
	b.lock.Unlock()
	return ok
}

func (b *BlockingSet[T]) Remove(key T) bool {
	b.lock.Lock()
	ok := b.keys.Remove(key)
	b.lock.Unlock()
	return ok
}

func (b *BlockingSet[T]) Contains(key T) bool {
	b.lock.RLock()
	ok := b.keys.Contains(key)
	b.lock.RUnlock()
	return ok
}
