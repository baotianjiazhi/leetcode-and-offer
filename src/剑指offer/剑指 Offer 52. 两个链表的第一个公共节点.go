package main


func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := map [*ListNode]int{}

	for headA != nil || headB != nil {
		if headA != nil {
			m[headA]++
			if m[headA] == 2 {
				return headA
			}
			headA = headA.Next

		}
		if headB != nil {
			m[headB]++
			if m[headB] == 2 {
				return headB
			}
			headB = headB.Next

		}

	}


	return nil
}
