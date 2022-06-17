package search

//https://leetcode.com/problems/number-of-provinces

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	var dfs func(from int)
	dfs = func(from int) {
		visited[from] = true
		for to, connect := range isConnected[from] {
			if connect == 1 && !visited[to] {
				dfs(to)
			}
		}
	}
	res := 0
	for i, v := range visited {
		if !v {
			res++
			dfs(i)
		}
	}
	return res
}

func findCircleNum2(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	res := 0
	for i, v := range visited {
		if !v {
			res++
			visited[i] = true
			record := []int{i}
			for len(record) != 0 {
				from := record[0]
				record = record[1:]
				visited[from] = true
				for to, connect := range isConnected[from] {
					if connect == 1 && !visited[to] {
						record = append(record, to)
					}
				}
			}
		}
	}
	return res
}
