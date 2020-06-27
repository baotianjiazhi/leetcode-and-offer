package main


func reverseList(head *ListNode) *ListNode {

	if head == nil  || head.Next == nil{
		return head
	}
	l := head
	r := head.Next

	for r != nil {
		tmp := r.Next
		r.Next = l
		l = r
		r = tmp
	}
	head.Next = nil
	return l
}