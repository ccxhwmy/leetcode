package search

import "testing"

func TestBinarySearch(t *testing.T) {
	if binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9) != 4 {
		t.Fatal("failed")
	}
	if binarySearch([]int{-1, 0, 3, 5, 9, 12}, 2) != -1 {
		t.Fatal("failed")
	}
}
