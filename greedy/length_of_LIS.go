package greedy

//https://leetcode.com/problems/longest-increasing-subsequence

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	resLen := 1
	dp := make([]int, n+1)
	dp[resLen] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > dp[resLen] {
			resLen = resLen + 1
			dp[resLen] = nums[i]
		} else {
			l, r := 1, resLen
			pos := 0
			for l <= r {
				mid := l + (r - l) / 2
				if dp[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			dp[pos+1] = nums[i]
		}
	}

	return resLen
}
