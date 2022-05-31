package binary_tree

//https://leetcode.com/problems/sum-root-to-leaf-numbers

func sumNumbers(root *TreeNode) int {
	res := 0
	var dfs func(path int, node *TreeNode)
	dfs = func(path int, node *TreeNode) {
		path = path * 10 + node.Val
		if node.Left == nil && node.Right == nil {
			res += path
		}
		if node.Left != nil {
			dfs(path, node.Left)
		}
		if node.Right != nil {
			dfs(path, node.Right)
		}
	}
	dfs(0, root)
	return res
}
