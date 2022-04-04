package list

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	pq := new(PriorityQueue)
	for _, list := range lists {
		if list != nil {
			heap.Push(pq, list)
		}
	}
	dummyHead := new(ListNode)
	curNode := dummyHead
	for pq.Len() != 0 {
		node := heap.Pop(pq).(*ListNode)
		curNode.Next = node
		curNode = curNode.Next
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
	}
	return dummyHead.Next
}

type PriorityQueue []*ListNode

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].Val < (*pq)[j].Val
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	res := (*pq)[len(*pq) - 1]
	*pq = (*pq)[:len(*pq) - 1]
	return res
}
