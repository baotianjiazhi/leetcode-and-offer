package main


type Node struct {
	Val int
	Next *Node
	Random *Node
}
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	clonenextnode(head)
	clonerandomnode(head)
	return seperatenode(head)
}


func clonenextnode(head *Node) {
	pnode := head
	for pnode != nil {
		cnode := new(Node)
		cnode.Val = pnode.Val
		cnode.Next = pnode.Next
		cnode.Random = nil

		pnode.Next = cnode
		pnode = cnode.Next
	}
}

func clonerandomnode(head *Node) {
	pnode := head
	for pnode != nil {
		cnode := pnode.Next
		if (pnode.Random != nil) {
			cnode.Random = pnode.Random.Next
		}

		pnode = cnode.Next
	}
}

func seperatenode(head *Node) *Node {
	pnode := head
	chead := new(Node)
	cnode := new(Node)

	if pnode != nil {
		chead = pnode.Next
		cnode = pnode.Next

		pnode.Next = cnode.Next
		pnode = pnode.Next
	}

	for pnode != nil {
		cnode.Next = pnode.Next
		cnode = pnode.Next
		pnode.Next = cnode.Next
		pnode = cnode.Next
	}

	return chead
}