package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	tmp := head
	for i := 0; i < k-1; i++ {
		head = head.Next
	}

	for head.Next != nil {
		head = head.Next
		tmp = tmp.Next
	}

	return tmp
}