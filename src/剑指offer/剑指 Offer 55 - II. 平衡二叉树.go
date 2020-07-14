package main

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left - right > 1 || left - right < -1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}
