package glibs

import (
	"sync"
)

type Pool[T any] struct {
	noCopy   noCopy
	internal *sync.Pool
}

func NewPool[T any]() *Pool[T] {
	return &Pool[T]{
		internal: &sync.Pool{
			New: func() any {
				return new(T)
			},
		},
	}
}

func (p *Pool[T]) Get() T {
	return p.internal.Get().(T)
}

func (p *Pool[T]) Set(t T) {
	p.internal.Put(t)
}
