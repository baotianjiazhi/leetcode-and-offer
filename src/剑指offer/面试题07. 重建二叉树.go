package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /**
 典型只会思想，写不出来代码
  */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	for k := range inorder {
		if preorder[0] == inorder[k] {
			return &TreeNode{
				Val: preorder[0],
				Left: buildTree(preorder[1:k+1], inorder[0:k]),
				Right:buildTree(preorder[k+1:], inorder[k+1:]),
			}
		}
	}

	return nil
}
