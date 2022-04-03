package binary_tree

//https://leetcode.com/problems/binary-tree-level-order-traversal/

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	record := make([]*TreeNode, 0)
	record = append(record, root)
	for len(record) > 0 {
		curLevel := make([]int, 0)
		curLen := len(record)
		for i := 0; i < curLen; i++ {
			curNode := record[0]
			record = record[1:]
			curLevel = append(curLevel, curNode.Val)
			if curNode.Left != nil {
				record = append(record, curNode.Left)
			}
			if curNode.Right != nil {
				record = append(record, curNode.Right)
			}
		}
		res = append(res, curLevel)
	}
	return res
}
