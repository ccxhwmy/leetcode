package binary_tree

//https://leetcode.com/problems/binary-tree-right-side-view

func rightSideVies(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	record := make([]*TreeNode, 0)
	record = append(record, root)
	for len(record) != 0 {
		curLen := len(record)
		var lastNode *TreeNode
		for i := 0; i < curLen; i++ {
			lastNode = record[0]
			record = record[1:]
			if lastNode.Left != nil {
				record = append(record, lastNode.Left)
			}
			if lastNode.Right != nil {
				record = append(record, lastNode.Right)
			}
		}
		res = append(res, lastNode.Val)
	}
	return res
}
