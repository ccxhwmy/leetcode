package other

import "math/rand"

//https://leetcode.com/problems/implement-rand10-using-rand7

func rand10() int {
	for {
		row := rand.Intn(7)
		col := rand.Intn(7)
		idx := (row - 1) * 7 + col
		if idx <= 40 {
			return idx % 10
		}
	}
}
