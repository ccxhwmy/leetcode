package list

//https://leetcode.com/problems/reverse-linked-list-ii

func reverseBetween2(head *ListNode, left, right int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	preListTail := dummyHead
	for i := 1; i < left && preListTail != nil; i++ {
		preListTail = preListTail.Next
	}
	if preListTail == nil {
		return nil
	}
	cur := preListTail.Next
	midListTail := cur
	var pre *ListNode = nil
	for i := left; i <= right && cur != nil; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	preListTail.Next = pre
	midListTail.Next = cur
	retNode := dummyHead.Next
	dummyHead.Next = nil

	return retNode
}
