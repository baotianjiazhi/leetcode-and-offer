package main


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