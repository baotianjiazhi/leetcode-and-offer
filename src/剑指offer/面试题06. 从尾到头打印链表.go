package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i := 0;i < len(res)/2; i++ {
		temp := res[i]
		res[i] = res[len(res)-1-i]
		res[len(res)-1-i] = temp
	}


	return res
}


func reversePrint_recursion(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	return appendData(head)
}

func appendData(head *ListNode) []int {
	if head.Next != nil {
		list := appendData(head.Next)
		list = append(list, head.Val)
		return list
	}
	return []int{head.Val}
}

