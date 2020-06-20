package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func levelOrder(root *TreeNode) []int {
	if root == nil {
		return  nil
	}
	queue := []*TreeNode{}
	res := []int{}
	queue = append(queue, root)
	for len(queue) != 0 {
		if root.Left != nil {
			queue = append(queue, root.Left)
		}

		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		res = append(res, root.Val)
		queue = queue[1:]
		if len(queue) != 0 {
			root = queue[0]
		}
	}

	return res
}
