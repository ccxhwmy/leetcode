package dynamic

//https://leetcode.com/problems/longest-consecutive-sequence

func longestConsecutive(nums []int) int {
	record := make(map[int]int, 0)
	for _, v := range nums {
		record[v] = 1
	}
	var forward func(num int) int
	forward = func(num int) int {
		cnt, ok := record[num]
		if !ok {
			return 0
		}
		if cnt > 1 {
			return cnt
		}
		cnt = forward(num + 1) + 1
		record[num] = cnt
		return cnt
	}
	for _, v := range nums {
		forward(v)
	}
	res := 0
	for _, v := range record {
		if res < v {
			res = v
		}
	}
	return res
}
