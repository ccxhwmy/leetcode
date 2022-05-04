package dynamic

//https://leetcode.com/problems/longest-increasing-subsequence

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	res := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := i-1; j >= 0; j-- {
			if nums[i] <= nums[j] {
				continue
			}
			if dp[j] + 1 > dp[i] {
				dp[i] = dp[j] + 1
				if dp[i] > res {
					res = dp[i]
					break
				}
			}
		}
	}
	return res
}
