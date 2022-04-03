package trap_water

func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}
	leftHigh := make([]int, n)
	leftHigh[0] = height[0]
	for i := 1; i < n; i++ {
		leftHigh[i] = max(leftHigh[i-1], height[i])
	}
	rightHigh := make([]int, n)
	rightHigh[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightHigh[i] = max(rightHigh[i+1], height[i])
	}
	res := 0
	for i := 1; i < n-1; i++ {
		res += min(rightHigh[i], leftHigh[i]) - height[i]
	}
	return res
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
