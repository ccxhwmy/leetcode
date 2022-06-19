package backtrack

//https://leetcode.com/problems/combination-sum

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	n := len(candidates)
	if n == 0 {
		return res
	}
	path := make([]int, 0)
	sum := 0
	var generateCombinationSum func(index int)
	generateCombinationSum = func(index int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		if sum > target {
			return
		}
		for i := index; i < n; i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			generateCombinationSum(i)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	generateCombinationSum(0)
	return res
}
