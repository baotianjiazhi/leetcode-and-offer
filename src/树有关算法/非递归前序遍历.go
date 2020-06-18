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
func preorderTraversal(root *TreeNode) []int {
	res = []int{}
	dfs(root)
	return res
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res = append(res, root.Val)
	dfs(root.Left)
	dfs(root.Right)
	return res
}

/**
	非递归方法进行前序遍历
 */

var stack []*TreeNode
func preorderTraversal_1(root *TreeNode) []int{
	res := []int{}
	if root == nil {
		return nil
	}

	for root != nil || len(stack) != 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}

	return res
}