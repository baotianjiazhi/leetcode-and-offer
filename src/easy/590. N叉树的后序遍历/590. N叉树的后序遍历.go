package main

type Node struct {
	Val int
	Children []*Node
}

var res []int
func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	res = []int{}
	dfs(root)

	return res
}

func dfs(root *Node) {
	if root != nil {
		for _, v := range root.Children {
			dfs(v)
		}
	}

	res = append(res, root.Val)
}
