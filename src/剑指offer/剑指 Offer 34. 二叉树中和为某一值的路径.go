package main


var res [][]int
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	res = [][]int{}
	judge(root, sum, []int{})
	return res
}

func judge(root *TreeNode, sum int, path []int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	judge(root.Left, sum-root.Val, path)
	judge(root.Right, sum-root.Val, path)
}