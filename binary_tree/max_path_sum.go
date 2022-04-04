package binary_tree

import "math"

//https://leetcode.com/problems/binary-tree-maximum-path-sum/

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := math.MinInt32
	var maxPath func(node *TreeNode) int
	maxPath = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftPath := max(maxPath(node.Left), 0)
		rightPath := max(maxPath(node.Right), 0)
		res = max(leftPath + rightPath + node.Val, res)
		return node.Val + max(leftPath, rightPath)
	}
	maxPath(root)
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
