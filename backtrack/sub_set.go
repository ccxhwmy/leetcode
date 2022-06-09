package backtrack

//https://leetcode.com/problems/subsets

func subset(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	if n == 0 {
		return res
	}
	path := make([]int, 0)
	var dfs func(index int)
	dfs = func(index int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := index; i < n; i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path) - 1]
		}
	}
	dfs(0)
	return res
}

func subset2(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := make([]int, 0)
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		res = append(res, append([]int(nil), set...))
	}
	return res
}
