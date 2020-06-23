package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var t int = 0
	l3 := &ListNode{}
	deadline := l3
	for l1 != nil || l2 != nil || t != 0 {
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}
		deadline.Next = &ListNode{
			Val: t % 10,
		}
		deadline = deadline.Next
		t = t / 10
	}

	return l3.Next
}
