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

func TestReverseList(t *testing.T) {
	t.Log(getList(reverseList(testHead)))
}

func TestReverseListBetween(t *testing.T) {
	t.Log(getList(reverseListBetween(testHead, 2, 4)))
}

func TestGetKthFromEnd(t *testing.T) {
	t.Log(getList(getKthFromEnd(testHead, 2)))
}

func TestReverseKGroup(t *testing.T) {
	t.Log(getList(reverseKGroup(testHead, 2)))
}

func TestReverseBetween(t *testing.T) {
	t.Log(getList(reverseBetween2(testHead, 2, 4)))
}

func TestReorderList(t *testing.T) {
	reorderList(testHead)
	t.Log(getList(testHead))
}
