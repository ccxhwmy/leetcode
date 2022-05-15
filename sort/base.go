package sort

type ListNode struct {
	Val  int
	Next *ListNode
}

func getList(head *ListNode) []int {
	res := make([]int, 0)
	if head == nil {
		return res
	}
	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}
