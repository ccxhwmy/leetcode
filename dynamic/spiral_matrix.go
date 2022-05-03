package dynamic

//https://leetcode.com/problems/spiral-matrix

func spiralMatrix(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	d := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	res := make([]int, 0)
	x, y := 0, 0
	curd := 0
	for len(res) < n * m {
		if !visited[x][y] {
			visited[x][y] = true
			res = append(res, matrix[x][y])
		}
		nextX, nextY := x + d[curd][0], y + d[curd][1]
		if isValid(m, n, nextX, nextY) && !visited[nextX][nextY] {
			x, y = nextX, nextY
		} else {
			curd = (curd + 1) % 4
		}
	}

	return res
}
