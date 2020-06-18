package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int
func inorderTraversal(root *TreeNode) []int {
	res = []int{}
	return dfs(root)

}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	dfs(root.Left)
	res = append(res, root.Val)
	dfs(root.Right)

	return res
}

func inorderTraversal_1(root *TreeNode) []int{
	if root == nil {
		return []int{}
	}
	res := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}