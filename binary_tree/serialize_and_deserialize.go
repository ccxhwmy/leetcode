package binary_tree

import (
	"strconv"
	"strings"
)

//https://leetcode.com/problems/serialize-and-deserialize-binary-tree

func serializeWithBfs(root *TreeNode) string {
	q := make([]*TreeNode, 0)
	q = append(q, root)
	res := make([]string, 0)
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			res = append(res, "null")
		} else {
			res = append(res, strconv.Itoa(node.Val))
			q = append(q, node.Left)
			q = append(q, node.Right)
		}
	}
	return strings.Join(res, ",")
}

func deserializeWithBfs(data string) *TreeNode {
	if data == "null" {
		return nil
	}
	record := strings.Split(data, ",")
	q := make([]*TreeNode, 0)
	val, _ := strconv.Atoi(record[0])
	root := &TreeNode{Val: val}
	q = append(q, root)
	cursor := 1
	for cursor < len(record) {
		node := q[0]
		q = q[1:]
		leftVal := record[cursor]
		rightVal := record[cursor + 1]
		if leftVal != "null" {
			val, _ := strconv.Atoi(leftVal)
			leftNode := &TreeNode{Val: val}
			node.Left = leftNode
			q = append(q, leftNode)
		}
		if rightVal != "null" {
			val, _ := strconv.Atoi(rightVal)
			rightNode := &TreeNode{Val: val}
			node.Right = rightNode
			q = append(q, rightNode)
		}
		cursor += 2
	}
	return root
}
