package dynamic

import "math"

//https://leetcode.com/problems/integer-break

func integerBreak(n int) int {
	dp := make([]int, n + 1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j * (i - j), j * dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}

func integerBreak2(n int) int {
	if n <= 3 {
		return n - 1
	}
	quotient := n / 3
	remainder := n % 3
	if remainder == 0 {
		return int(math.Pow(3, float64(quotient)))
	} else if remainder == 1 {
		return int(math.Pow(3, float64(quotient - 1))) * 4
	}
	return int(math.Pow(3, float64(quotient))) * 2
}
