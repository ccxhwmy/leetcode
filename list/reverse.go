package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseListBetween(head *ListNode, left, right int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	preListTail := dummyHead
	count := 1
	for count < left && preListTail != nil {
		preListTail = preListTail.Next
		count++
	}
	cur := preListTail.Next
	reverseHead := cur
	var pre *ListNode = nil
	for count <= right && cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		count++
	}
	preListTail.Next = pre
	reverseHead.Next = cur
	retNode := dummyHead.Next
	dummyHead.Next = nil
	return retNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	tail, res := dummyHead, dummyHead
	for k > 0 {
		tail = tail.Next
		k--
	}
	for tail != nil {
		tail = tail.Next
		res = res.Next
	}
	return res
}
