package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	res := true
	if root != nil {
		res = treeistree(root.Left, root.Right)
	}

	return res
}


func treeistree(A, B *TreeNode) bool {

	if A == nil && B == nil {
		return true
	}

	if A == nil || B == nil {
		return false
	}

	if A.Val != B.Val {
		return false
	}

	return treeistree(A.Left, B.Right) && treeistree(A.Right, B.Left)
}