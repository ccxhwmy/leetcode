package search

//Find Minimum in Rotated Sorted Array

func searchRange(nums []int, target int) []int {
	firstIndex := lowerBound(nums, target)
	lastIndex := lowerBound(nums, target + 1) - 1
	if firstIndex == len(nums) || nums[firstIndex] != target {
		return []int{-1, -1}
	}
	return []int{firstIndex, max(firstIndex, lastIndex)}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r - l) / 2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
