package binary_tree

//https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == root || q == root {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Left, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
