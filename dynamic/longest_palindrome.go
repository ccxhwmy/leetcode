package dynamic

//https://leetcode.com/problems/longest-palindromic-substring

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, 0)
	for i := 0; i < n; i++ {
		tmp := make([]bool, n)
		dp = append(dp, tmp)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	maxLen := 1
	begin := 0
	for L := 2; L <= n; L++ {
		for i := 0; i < n; i++ {
			j := L + i - 1
			if j >= n {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func longestPalindrome2(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	begin, maxLen := 0, 1
	for i := 0; i < n; {
		if n - i < maxLen / 2 {
			break
		}
		l, r := i, i
		for r < n - 1 && s[r] == s[r+1] {
			r++
		}
		i = r + 1
		for r < n - 1 && l > 0 && s[l-1] == s[r+1] {
			l--
			r++
		}
		if r - l + 1 > maxLen {
			begin = l
			maxLen = r - l + 1
		}
	}
	return s[begin:begin+maxLen]
}
