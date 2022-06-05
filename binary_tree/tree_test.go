package binary_tree

import (
	"testing"
)

//		1
//	2		3
//4   5   6
//
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

func TestIsValidBST(t *testing.T) {
	t.Log(isValidBST2(testRoot))
}

func TestInorderTraversalWithRecursion(t *testing.T) {
	t.Log(inorderTraversalWithRecursion(testRoot))
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

var testRoot2 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	},
	Right: &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	},
}

var testRoot3 = &TreeNode{
	Val: -10,
	Left: &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	},
	Right: &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val:   15,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	},
}

func TestMaxPathSum(t *testing.T) {
	if maxPathSum(testRoot2) != 6 {
		t.Fatal("failed")
	}
	if maxPathSum(testRoot3) != 42 {
		t.Fatal("failed")
	}
}

func TestSerializeAndDeserialize(t *testing.T) {
	s := serializeWithBfs(testRoot)
	root := deserializeWithBfs(s)
	t.Log("root: ", root)
}

func TestRightSideVeiw(t *testing.T) {
	res := rightSideVies(testRoot)
	t.Log(res)
}

func TestBuildTree(t *testing.T) {
	root := buildTreeWithRecursion([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	t.Log(root)
}

var testSumNumbers1 = &TreeNode{
	Val:   1,
	Left:  &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	},
	Right: &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	},
}

var testSumNumbers2 = &TreeNode{
	Val:   4,
	Left:  &TreeNode{
		Val:   9,
		Left:  &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	},
}

func TestSumNumbers(t *testing.T) {
	if sumNumbers(testSumNumbers1) != 25 {
		t.Fatal("failed")
	}
	if sumNumbers(testSumNumbers2) != 1026 {
		t.Fatal("failed")
	}
}
