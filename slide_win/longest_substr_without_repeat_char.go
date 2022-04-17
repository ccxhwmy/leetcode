package slide_win

//https://leetcode.com/problems/longest-substring-without-repeating-characters/

func longestSubStrWithoutRepeatChar(s string) int {
	record := make([]int, 256)
	l, r := 0, -1
	res := 0
	for l < len(s) {
		if r+1 < len(s) && record[s[r+1]] == 0 {
			r++
			record[s[r]]++
			if res < r-l+1 {
				res = r - l + 1
			}
		} else {
			record[s[l]]--
			l++
		}
	}
	return res
}
