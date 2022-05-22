package backtrack

import (
	"strconv"
	"strings"
)

//https://leetcode.com/problems/restore-ip-addresses

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	var dfs func([]string, string)
	dfs = func(putStr []string, curStr string) {
		if len(putStr) == 4 && len(curStr) == 0 {
			res = append(res, strings.Join(putStr, "."))
			return
		}
		if len(putStr) >= 4 {
			return
		}
		for i := 1; i <= 3; i++ {
			if len(curStr) < i {
				return
			}
			tmpStr := curStr[:i]
			if !isValidIpAddr(tmpStr) {
				return
			}
			putStr = append(putStr, tmpStr)
			dfs(putStr, curStr[i:])
			putStr = putStr[:len(putStr)-1]
		}
	}
	dfs(make([]string, 0), s)
	return res
}

func isValidIpAddr(s string) bool {
	if s == "0" {
		return true
	}
	if s[0] == '0' {
		return false
	}
	n, _ := strconv.Atoi(s)
	if n > 255 || n < 0 {
		return false
	}
	return true
}
