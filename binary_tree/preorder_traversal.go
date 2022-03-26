package binary_tree

func preorderTraversalWithRecursion(root *TreeNode) []int {
	res := make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return res
}

func preorderTraversalWithIterate(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	record := make([]*TreeNode, 0)
	record = append(record, root)
	for len(record) != 0 {
		cur := record[len(record) - 1]
		record = record[0:len(record) - 1]
		res = append(res, cur.Val)
		if cur.Right != nil {
			record = append(record, cur.Right)
		}
		if cur.Left != nil {
			record = append(record, cur.Left)
		}
	}
	return res
}
