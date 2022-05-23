package array_permutation

//https://leetcode.com/problems/next-permutation

func nextPermutation(nums []int)  {
	n := len(nums)
	if n <= 1 {
		return
	}
	i := n - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {}
	if i < 0 {
		reverse(nums, 0, n-1)
		return
	}
	k := n - 1
	for ; k > i && nums[i] >= nums[k]; k-- {}
	nums[k], nums[i] = nums[i], nums[k]
	reverse(nums, i + 1, n - 1)
	return
}

func reverse(arr []int, l, r int) {
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
