package double_pointer

import "sort"

//https://leetcode.com/problems/3sum

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	res := make([][]int, 0)
	sort.Ints(nums)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := nums[first] * -1
		third := n - 1
		for second := first + 1; second < third; second++ {
			if second > first + 1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second] + nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second] + nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}
