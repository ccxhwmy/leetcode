package dynamic

//https://leetcode.com/problems/coin-change

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := 1; i < amount+ 1; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if i >= v {
				dp[i] = min(dp[i], dp[i-v] + 1)
			}
		}
	}
	if dp[amount] == amount + 1 {
		return -1
	}
	return dp[amount]
}
