package utils

import (
	"fmt"
	"sync"
)

// 单例模式 Singleton

type Singleton[T any] struct {
	ins   T
	newer func() T
	once  *sync.Once
}

func NewSingleton[T any](newer func() T) *Singleton[T] {
	if newer == nil {
		panic(fmt.Errorf("newer func is nil"))
	}
	return &Singleton[T]{
		newer: newer,
		once:  &sync.Once{},
	}
}

func (s *Singleton[T]) Get() T {
	s.once.Do(func() {
		s.ins = s.newer()
	})
	return s.ins
}
