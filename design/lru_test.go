package design

import "testing"

func TestLru(t *testing.T) {
	lru := Constructor(10)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 3)
	lru.Put(5, 3)
	lru.Put(6, 3)
	lru.Put(7, 3)
	lru.Put(8, 3)
	lru.Put(9, 3)
	lru.Put(10, 3)
	lru.Put(11, 3)
	lru.Put(12, 3)
	if lru.Get(5) != 3 {
		t.Fatal("failed")
	}
	if lru.Get(1) != -1 {
		t.Fatal("failed")
	}
}
