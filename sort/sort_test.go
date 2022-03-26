package sort

import "testing"

func TestQuickSort(t *testing.T)  {
	nums := []int{5,2,3,1}
	t.Log(quickSort(nums))
}
