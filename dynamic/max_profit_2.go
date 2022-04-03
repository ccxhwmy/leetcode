package dynamic

func maxProfit2WithDynamic(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func maxProfit2WithDynamic2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		//todayDp0 := max(dp0, dp1 + prices[i])
		//todayDp1 := max(dp1, dp0 - prices[i])
		//dp0 = todayDp0
		//dp1 = todayDp1
		//equal to above
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	return dp0
}

func maxProfit2WithGreedy(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	res := 0
	for i := 1; i < n; i++ {
		res += max(0, prices[i] - prices[i-1])
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
