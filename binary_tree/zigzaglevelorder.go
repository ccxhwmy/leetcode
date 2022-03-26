package binary_tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	judge := true
	record := make([]*TreeNode, 0)
	record = append(record, root)
	for len(record) > 0 {
		curLevel := make([]int, 0)
		curLen := len(record)
		nextRecord := make([]*TreeNode, 0)
		for i := curLen - 1; i >= 0; i-- {
			curNode := record[i]
			curLevel = append(curLevel, curNode.Val)
			if judge {
				if curNode.Left != nil {
					nextRecord = append(nextRecord, curNode.Left)
				}
				if curNode.Right != nil {
					nextRecord = append(nextRecord, curNode.Right)
				}
			} else {
				if curNode.Right != nil {
					nextRecord = append(nextRecord, curNode.Right)
				}
				if curNode.Left != nil {
					nextRecord = append(nextRecord, curNode.Left)
				}
			}
		}
		judge = !judge
		res = append(res, curLevel)
		record = nextRecord
	}
	return res
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	record := make([]*TreeNode, 0)
	record = append(record, root)
	judge := true
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
		if !judge {
			for l, r := 0, len(curLevel) - 1; l < r; {
				curLevel[l], curLevel[r] = curLevel[r], curLevel[l]
				l++
				r--
			}
		}
		judge = !judge
		res = append(res, curLevel)
	}
	return res
}
