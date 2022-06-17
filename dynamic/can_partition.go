package dynamic

func canPartition(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	sum, maxNum := 0, 0
	for _, v := range nums {
		sum += v
		if v > maxNum {
			maxNum = v
		}
	}
	if sum % 2 != 0 {
		return false
	}
	target := sum / 2
	if target < maxNum {
		return false
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target + 1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 0; j <= target; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}
