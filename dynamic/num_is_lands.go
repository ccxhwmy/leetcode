package dynamic

//https://leetcode.com/problems/number-of-islands

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	d := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nextX, nextY := x + d[i][0], y + d[i][1]
			if isValid(m, n, nextX, nextY) && grid[nextX][nextY] == '1' && !visited[nextX][nextY] {
				dfs(nextX, nextY)
			}
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}

func isValid(m, n, x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}
