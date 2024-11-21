package utils

import "sync"

type Once[T any] struct {
	once     *sync.Once
	newer    func() T
	instance T
}

func NewOnce[T any](newer func() T) *Once[T] {
	if newer == nil {
		panic("NewOnce:newer is nil")
	}
	once := &Once[T]{
		once:  &sync.Once{},
		newer: newer,
	}
	return once
}

func (o *Once[T]) Do() T {
	o.once.Do(func() {
		o.instance = o.newer()
	})
	return o.instance
}
