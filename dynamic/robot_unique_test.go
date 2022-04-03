package dynamic

import "testing"

var (
	obstacleGrid1 = [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
)

func TestRobotUnique(t *testing.T) {
	if uniquePaths(3, 2) != 3 {
		t.Fatal("failed")
	}
	if uniquePaths(3, 7) != 28 {
		t.Fatal("failed")
	}
}

func TestUniquePathsWithObstacles(t *testing.T) {
	if uniquePathsWithObstacles(obstacleGrid1) != 2 {
		t.Fatal("failed")
	}
}
