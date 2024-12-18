package utils

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	s := NewSingleton(func() int {
		return 1
	})
	if s.Get() != 1 {
		t.Fatal("TestSingleton failed")
	}
}

func BenchmarkSingleton(b *testing.B) {
	s := NewSingleton(func() int {
		return 1
	})
	s.Get()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Get()
	}
}
