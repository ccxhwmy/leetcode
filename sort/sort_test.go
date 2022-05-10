package sort

import "testing"

func TestQuickSort(t *testing.T)  {
	nums := []int{5,2,3,1}
	t.Log(quickSort(nums))
}

func TestMerge(t *testing.T) {
	//t.Log(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	t.Log(merge([][]int{{1, 4}, {0, 2}, {3, 5}}))
}
