package binary_tree

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(node *TreeNode, left int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if node.Left == nil && node.Right == nil && left == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, target)
	return res
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}
