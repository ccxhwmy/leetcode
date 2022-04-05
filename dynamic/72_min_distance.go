package dynamic

//https://leetcode.com/problems/edit-distance

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, 0)
	for i := 0; i < m + 1; i++ {
		tmp := make([]int, n + 1)
		dp = append(dp, tmp)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minThreeInt(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	display2DimensionArray(dp)
	return dp[m][n]
}
