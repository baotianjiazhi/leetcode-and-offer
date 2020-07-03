package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return pathFrom(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func judge(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val == sum {
		res++
	}

	res += judge(root.Left, sum-root.Val)
	res += judge(root.Right, sum-root.Val)

	return res
}

func pathFrom(root *TreeNode, sum int) int {
	if root == nil { // 如果树为空
		return 0 // 就返回数量0
	}
	cnt := 0             // 否则初始化路径数量为0
	if root.Val == sum { // 如果当前节点值等于sum
		cnt++ // 说明找到一条满足条件的路径,数量+1
	}
	// 接着无论找没找到满足条件的路径
	// 都向下遍历左右子树继续寻找
	// 目标值更新为sum减去当前节点值
	// 因为即使当前到当前节点已经是一条满足条件的路径，
	// 后面也可能出现正负节点值互相抵消的情况
	// 从而新增一条和为给定值的更长的路径
	cnt += pathFrom(root.Left, sum-root.Val)
	cnt += pathFrom(root.Right, sum-root.Val)
	return cnt // 最后返回路径数量即可
}