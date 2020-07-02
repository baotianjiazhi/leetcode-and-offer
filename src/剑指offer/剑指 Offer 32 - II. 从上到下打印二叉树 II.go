package main

func levelOrder_2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	floor := 0
	for len(queue) != 0 {
		res = append(res, []int{})
		tmp := []*TreeNode{}
		for _, v := range queue {
			res[floor] = append(res[floor], v.Val)
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}

			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}
		floor++
		queue = tmp
	}
	return res
}
