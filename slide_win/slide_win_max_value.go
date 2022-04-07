package slide_win

import "container/heap"

//https://leetcode.com/problems/sliding-window-maximum

var tmpNums []int

type priorityQueue []int

func (pq *priorityQueue) Len() int {
	return len(*pq)
}

func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *priorityQueue) Less(i, j int) bool {
	return tmpNums[(*pq)[i]] > tmpNums[(*pq)[j]]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *priorityQueue) Pop() interface{} {
	res := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return res
}

func maxSlidingWindowWithPriorityQueue(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	tmpNums = nums
	pq := &priorityQueue{}
	for i := 0; i < k; i++ {
		heap.Push(pq, i)
	}
	ans := make([]int, 0)
	ans = append(ans, tmpNums[(*pq)[0]])
	for i := k; i < len(nums); i++ {
		heap.Push(pq, i)
		for (*pq)[0] <= i - k {
			heap.Pop(pq)
		}
		ans = append(ans, tmpNums[(*pq)[0]])
	}
	return ans
}
