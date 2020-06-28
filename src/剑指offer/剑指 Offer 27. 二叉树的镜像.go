package main


func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Right, root.Left = root.Left, root.Right
	if root.Right != nil {
		mirrorTree(root.Right)
	}

	if root.Left != nil {
		mirrorTree(root.Left)
	}

	return root
}


