package main

import (
	"fmt"
)

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

func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	} else {
		left := maxDepth(root.Left) + 1
		right := maxDepth(root.Right) + 1
		return max(left, right)
	}
}

func max(a int, b int) int {
	if a > b{
		return a
	} else if b > a {
		return b
	}

	return a
}

func main() {
	tree2 := &TreeNode{}
	tree := &TreeNode{
		Val: 2,
		Left: tree2,
	}

	fmt.Println(maxDepth(tree))
}