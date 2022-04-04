package list

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

func makeList(items []int) *ListNode {
	if len(items) == 0 {
		return nil
	}
	dummyHead := &ListNode{}
	cur := dummyHead
	for _, item := range items {
		cur.Next = &ListNode{
			Val:  item,
			Next: nil,
		}
		cur = cur.Next
	}
	return dummyHead.Next
}

func makeLists(items [][]int) []*ListNode {
	res := make([]*ListNode, 0)
	for _, item := range items {
		tmp := makeList(item)
		if tmp != nil {
			res = append(res, tmp)
		}
	}
	return res
}

func compareArrayOfInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
