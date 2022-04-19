package top_k

import "container/heap"

//https://leetcode.com/problems/kth-largest-element-in-an-array/

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	pq := new(priorityQueue)
	for _, num := range nums {
		if pq.Len() >= k {
			min := heap.Pop(pq).(int)
			if min < num {
				heap.Push(pq, num)
			} else {
				heap.Push(pq, min)
			}
		} else {
			heap.Push(pq, num)
		}
	}
	return heap.Pop(pq).(int)
}

type priorityQueue []int

func (p *priorityQueue) Len() int {
	return len(*p)
}

func (p *priorityQueue) Less(i, j int) bool {
	return (*p)[i] < (*p)[j]
}

func (p *priorityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *priorityQueue) Push(i interface{}) {
	*p = append(*p, i.(int))
}

func (p *priorityQueue) Pop() interface{} {
	res := (*p)[len(*p) - 1]
	*p = (*p)[0:len(*p) - 1]
	return res
}
