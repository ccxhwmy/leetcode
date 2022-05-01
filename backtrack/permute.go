package backtrack

import "sort"

//https://leetcode.com/problems/permutations

func permuteWithRecursion(nums []int) [][]int {
	res := make([][]int, 0)
	var callAllPermute func(index int)
	callAllPermute = func(index int) {
		if index == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
		}
		for j := index; j < len(nums); j++ {
			nums[j], nums[index] = nums[index], nums[j]
			callAllPermute(index + 1)
			nums[j], nums[index] = nums[index], nums[j]
		}
	}

	callAllPermute(0)

	return res
}

func permute(nums []int) [][]int {
	n := len(nums)
	if n <= 0 {
		return nil
	}
	res := make([][]int, 0)
	sort.Ints(nums)
	tmp := make([]int, n)
	copy(tmp, nums)
	ok := true
	for ok {
		res = append(res, tmp)
		tmp, ok = calcAllPermute(nums, n)
	}
	return res
}

func calcAllPermute(nums []int, n int) ([]int, bool) {
	i := n - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	if i < 0 {
		return nil, false
	}
	j := n - 1
	for ; j > i && nums[j] <= nums[i]; j-- {
	}
	nums[i], nums[j] = nums[j], nums[i]
	l, r := i+1, n-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	ret := make([]int, n)
	copy(ret, nums)
	return ret, true
}
