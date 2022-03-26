package binary_tree

import (
	"testing"
)

//		1
//	2		3
//4   5   6


var testRoot = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	},
}

func TestInorderTraversalWithIterate(t *testing.T) {
	t.Log(inorderTraversalWithIterate(testRoot))
}

func TestPostorderTraversalWithRecursion(t *testing.T) {
	t.Log(postorderTraversalWithRecursion(testRoot))
}

func TestPostorderTraversalWithIterate(t *testing.T) {
	t.Log(postorderTraversalWithIterate(testRoot))
}

func TestLevelOrder(t *testing.T) {
	t.Log(levelOrder(testRoot))
}

func TestZigzagLevelOrder(t *testing.T) {
	t.Log(zigzagLevelOrder2(testRoot))
}
