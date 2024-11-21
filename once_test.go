package utils

import (
	"math/rand"
	"testing"
)

func TestOnce(t *testing.T) {
	once := NewOnce(func() int {
		r := rand.Int()
		t.Log("rand:", r)
		return r
	})
	t.Log(once.Do())
	t.Log(once.Do())
	t.Log(once.Do())
}
