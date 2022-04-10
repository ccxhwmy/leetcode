package array_permutation

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	if firstMissingPositive([]int{1,2,0}) != 3 {
		t.Fatal("failed")
	}
	if firstMissingPositive([]int{3,4,-1,1}) != 2 {
		t.Fatal("failed")
	}
	if firstMissingPositive([]int{7,8,9,11,12}) != 1 {
		t.Fatal("failed")
	}
}
