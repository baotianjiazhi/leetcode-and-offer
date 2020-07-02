package main

func levelOrder_3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}

	stack_1 := []*TreeNode{}
	stack_2 := []*TreeNode{}
	stack_2 = append(stack_2, root)
	floor := 0
	for len(stack_1) != 0 || len(stack_2) != 0 {
		res = append(res, []int{})
		if floor % 2 != 0 {
			for i:= len(stack_1)-1; i >= 0; i--{
				res[floor] = append(res[floor], stack_1[i].Val)

				if stack_1[i].Right != nil {
					stack_2 = append(stack_2, stack_1[i].Right)
				}

				if stack_1[i].Left != nil {
					stack_2 = append(stack_2, stack_1[i].Left)
				}


			}
			stack_1 = stack_1[:0]
		} else {
			for i:= len(stack_2)-1; i >= 0; i--{
				res[floor] = append(res[floor], stack_2[i].Val)
				if stack_2[i].Left != nil {
					stack_1 = append(stack_1, stack_2[i].Left)
				}

				if stack_2[i].Right != nil {
					stack_1 = append(stack_1, stack_2[i].Right)
				}
			}
			stack_2 = stack_2[:0]
		}

		floor++
	}

	return res
}
