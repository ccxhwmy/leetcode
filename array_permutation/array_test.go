package array_permutation

import (
	"reflect"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	if firstMissingPositive([]int{1, 2, 0}) != 3 {
		t.Fatal("failed")
	}
	if firstMissingPositive([]int{3, 4, -1, 1}) != 2 {
		t.Fatal("failed")
	}
	if firstMissingPositive([]int{7, 8, 9, 11, 12}) != 1 {
		t.Fatal("failed")
	}
}

func TestNextPermutation(t *testing.T) {
	arr := []int{1, 2, 3}
	res := []int{1, 3, 2}
	nextPermutation(arr)
	if !reflect.DeepEqual(arr, res) {
		t.Fatal("failed")
	}
	arr2 := []int{3, 2, 1}
	res2 := []int{1, 2, 3}
	nextPermutation(arr2)
	if !reflect.DeepEqual(arr2, res2) {
		t.Fatal("failed")
	}
	arr3 := []int{1, 1, 5}
	res3 := []int{1, 5, 1}
	nextPermutation(arr3)
	if !reflect.DeepEqual(arr3, res3) {
		t.Fatal("failed")
	}
}
