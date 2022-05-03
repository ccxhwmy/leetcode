package list

//https://leetcode.com/problems/linked-list-cycle

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fastNode, lowNode := head, head
	for fastNode != nil && fastNode.Next != nil {
		lowNode = lowNode.Next
		fastNode = fastNode.Next.Next
		if lowNode == fastNode {
			return true
		}
	}
	return false
}

//https://leetcode.com/problems/linked-list-cycle-ii

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	intersect := getIntersect(head)
	if intersect == nil {
		return nil
	}
	ptr1 := head
	ptr2 := intersect
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return ptr1
}

func getIntersect(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fastNode, lowNode := head, head
	for fastNode != nil && fastNode.Next != nil {
		lowNode = lowNode.Next
		fastNode = fastNode.Next.Next
		if lowNode == fastNode {
			return fastNode
		}
	}
	return nil
}
