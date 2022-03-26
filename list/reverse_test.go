package list

import "testing"

var testHead = &ListNode{
	Val:  1,
	Next: &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{
					Val:  5,
					Next: &ListNode{
						Val:  6,
						Next: nil,
					},
				},
			},
		},
	},
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

func TestReverseList(t *testing.T) {
	t.Log(getList(reverseList(testHead)))
}

func TestReverseListBetween(t *testing.T) {
	t.Log(getList(reverseListBetween(testHead, 2, 4)))
}

func TestGetKthFromEnd(t *testing.T) {
	t.Log(getList(getKthFromEnd(testHead, 2)))
}
