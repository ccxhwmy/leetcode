package binary_tree

//https://leetcode.com/problems/binary-tree-postorder-traversal/

func postorderTraversalWithRecursion(root *TreeNode) []int {
	res := make([]int, 0)
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return res
}

func postorderTraversalWithIterate(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	for l, r := 0, len(res) - 1; l < r; {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
