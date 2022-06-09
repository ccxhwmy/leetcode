package search

//https://leetcode.com/problems/search-a-2d-matrix-ii

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	x, y := m - 1, 0
	for x >= 0 && y < n {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] < target {
			y++
		} else {
			x--
		}
	}
	return false
}
