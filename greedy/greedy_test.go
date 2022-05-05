package greedy

import "testing"

func TestLengthOfLIS(t *testing.T) {
	if lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}) != 4 {
		t.Fatal("failed")
	}
	if lengthOfLIS([]int{0, 1, 0, 3, 2, 3}) != 4 {
		t.Fatal("failed")
	}
	if lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}) != 1 {
		t.Fatal("failed")
	}
}
