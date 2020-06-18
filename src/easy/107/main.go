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

func levelOrderBottom(root *TreeNode) [][]int {
	curTree := make([]*TreeNode, 0)
	curTree = append(curTree, root)
	queue := make([][]int, 0)
	if root == nil {
		return queue
	}
	for {
		if len(curTree) == 0 {
			break
		}
		curVal := make([]int, 0)
		nextTree := make([]*TreeNode, 0)
		for _, tree := range curTree {
			curVal = append(curVal, tree.Val)
			if tree.Left != nil {
				nextTree = append(nextTree, tree.Left)
			}
			if tree.Right != nil{
				nextTree = append(nextTree, tree.Right)
			}
			fmt.Println(nextTree)
		}
		if len(curVal) != 0 {
			queue = append(queue, curVal)
		}
		curTree = nextTree
	}

	for i, j := 0, len(queue)-1; i < j; i, j = i+1, j-1 {
		queue[i], queue[j] = queue[j], queue[i]
	}
	return queue
}


func main() {
	tree2 := &TreeNode{
		Val: 0,
	}
	tree3 := &TreeNode{
		Val: 1,
	}
	tree := &TreeNode{
		Val: 2,
		Left: tree2,
		Right: tree3,
	}

	fmt.Println(levelOrderBottom(tree))
}