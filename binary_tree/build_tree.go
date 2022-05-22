package binary_tree

//https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal

func buildTreeWithRecursion(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	index := findIndex(inorder, preorder[0])
	root.Left = buildTreeWithRecursion(preorder[1:index+1], inorder[:index])
	root.Right = buildTreeWithRecursion(preorder[index+1:], inorder[index+1:])
	return root
}

func findIndex(inorder []int, target int) int {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == target {
			return i
		}
	}
	return -1
}
