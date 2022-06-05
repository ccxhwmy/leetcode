package binary_tree

import (
	"math"
)

//https://leetcode.com/problems/validate-binary-search-tree

func isValidBST(root *TreeNode) bool {
	var helper func(node *TreeNode, min, max int) bool
	helper = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)
	}
	return helper(root, math.MinInt64, math.MaxInt64)
}

func isValidBST2(root *TreeNode) bool {
	record := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		record = append(record, root.Val)
		root = root.Right
	}
	for i := 0; i < len(record) - 1; i++ {
		if record[i] > record[i+1] {
			return false
		}
	}
	return true
}

