package utils

import "sync"

type Once struct {
	once *sync.Once
	fn   func()
}

func NewOnce(fn func()) *Once {
	if fn == nil {
		panic("NewOnce:fn is nil")
	}
	once := &Once{
		once: &sync.Once{},
		fn:   fn,
	}
	return once
}

func (o *Once) Do() {
	o.once.Do(func() {
		o.fn()
	})
}
