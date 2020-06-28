package main

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}

	if A == nil || B == nil {
		return false
	}

	res := false
	if A.Val == B.Val {
		res = HaveTree(A, B)
	}

	if !res {
		res = isSubStructure(A.Left, B)
	}
	if !res {
		res = isSubStructure(A.Right, B)
	}


	return res
}

func HaveTree(A, B *TreeNode) bool{
	if B == nil {
		return true
	}

	if A == nil {
		return false
	}

	if A.Val != B.Val {
		return false
	}

	return HaveTree(A.Left, B.Left) && HaveTree(A.Right, B.Right)
}
