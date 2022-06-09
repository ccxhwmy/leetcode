package dynamic

import "fmt"

func minThreeInt(x, y, z int) int {
	ret := x
	if ret > y {
		ret = y
	}
	if ret > z {
		ret = z
	}
	return ret
}

func display2DimensionArray(arr [][]int) {
	m := len(arr)
	if m == 0 {
		return
	}
	n := len(arr[0])
	if n == 0 {
		return
	}
	fmt.Println("===============================")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%3d", arr[i][j])
		}
		fmt.Println("")
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
