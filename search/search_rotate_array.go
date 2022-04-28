package search

func searchRotateArray(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + (r - l) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target <= nums[n-1] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
