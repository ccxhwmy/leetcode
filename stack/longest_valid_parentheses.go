package stack

import "leetcode/base"

//https://leetcode.com/problems/longest-valid-parentheses

func longestValidParentheses(s string) int {
	record := make([]int, 0)
	record = append(record, -1)
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			record = append(record, i)
		} else {
			record = record[:len(record) - 1]
			if len(record) == 0 {
				record = append(record, i)
			} else {
				res = base.Max(res, i - record[len(record) - 1])
			}
		}
	}
	return res
}
