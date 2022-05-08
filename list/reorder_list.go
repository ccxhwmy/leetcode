package list

//https://leetcode.com/problems/reorder-list

func reorderList(head *ListNode) {
	midPreNode := findMidPreNode(head)
	midNode := midPreNode.Next
	midPreNode.Next = nil
	reversedNode := reverseList2(midNode)
	mergeList(head, reversedNode)
}

func findMidPreNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fastNode, slowNode := head, head
	for fastNode.Next != nil && fastNode.Next.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}
	return slowNode
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return  nil
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

func mergeList(l1, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for l1 != nil || l2 != nil {
		if l1 != nil {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}
		if l2 != nil {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}
	return dummyHead.Next
}
