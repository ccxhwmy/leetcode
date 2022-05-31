package backtrack

//https://leetcode.com/problems/generate-parentheses

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var dfs func(path string, l, r int)
	dfs = func(path string, l, r int) {
		if l == n && r == n {
			res = append(res, path)
			return
		}
		if l < n {
			dfs(path + "(", l + 1, r)
		}
		if r < l {
			dfs(path + ")", l, r + 1)
		}
	}
	dfs("", 0, 0)
	return res
}
