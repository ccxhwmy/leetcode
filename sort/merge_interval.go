package sort

import "sort"

//https://leetcode.com/problems/merge-intervals

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] <= intervals[j][1]
		} else {
			return intervals[i][0] <= intervals[j][0]
		}
	})
	for i := 0; i < len(intervals) - 1; {
		if intervals[i][1] >= intervals[i+1][0] {
			if intervals[i][1] < intervals[i+1][1] {
				intervals[i][1] = intervals[i+1][1]
			}
			intervals = append(intervals[:i+1], intervals[i+2:]...)
		} else {
			i++
		}
	}
	return intervals
}
