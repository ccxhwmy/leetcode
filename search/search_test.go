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

func TestSearchTotateArray(t *testing.T) {
	if searchRotateArray([]int{1,3,5}, 1) != 0 {
		t.Fatal("failed")
	}
	if searchRotateArray([]int{4,5,6,7,0,1,2}, 0) != 4 {
		t.Fatal("failed")
	}
	if searchRotateArray([]int{4,5,6,7,0,1,2}, 3) != -1 {
		t.Fatal("failed")
	}
	if searchRotateArray([]int{1}, 0) != -1 {
		t.Fatal("failed")
	}
}
