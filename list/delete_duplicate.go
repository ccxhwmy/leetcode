package list

//https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	curNode := dummyHead
	for curNode.Next != nil && curNode.Next.Next != nil {
		if curNode.Next.Val == curNode.Next.Next.Val {
			p := curNode.Next.Next
			for p.Next != nil && p.Next.Val == p.Val {
				p = p.Next
			}
			curNode.Next = p.Next
			continue
		}
		curNode = curNode.Next
	}
	return dummyHead.Next
}
