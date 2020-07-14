package main

var count int
var res int
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	count = k
	res = 0
	visit_tree(root)
	return res

}

func visit_tree(root *TreeNode){
	if root != nil {
		visit_tree(root.Right)
		count--
		if count == 0 {
			res = root.Val
		}
		visit_tree(root.Left)
	}
}
