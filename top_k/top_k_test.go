package top_k

import "testing"

func TestFindKthLargest(t *testing.T) {
	if findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2) != 5 {
		t.Fatal("failed")
	}
	if findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4) != 4 {
		t.Fatal("failed")
	}
}
