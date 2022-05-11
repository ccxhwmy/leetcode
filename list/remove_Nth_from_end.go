package list

//https://leetcode.com/problems/remove-nth-node-from-end-of-list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	firstNode, secondNode := dummyHead, dummyHead
	for ;n + 1 > 0 && firstNode != nil; n-- {
		firstNode = firstNode.Next
	}
	for firstNode != nil {
		firstNode = firstNode.Next
		secondNode = secondNode.Next
	}
	secondNode.Next = secondNode.Next.Next
	return dummyHead.Next
}
