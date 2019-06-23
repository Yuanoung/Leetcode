package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s, carry := 0, 0
	p1, p2, head := l1, l2, &ListNode{}
	p3 := head
	for p1 != nil && p2 != nil {
		s = p1.Val + p2.Val + carry
		carry = s / 10
		s = s % 10

		p3.Next = &ListNode{Val: s}
		p3 = p3.Next
		p1 = p1.Next
		p2 = p2.Next
	}
	if p1 == nil {
		p1 = p2
	}
	for p1 != nil {
		s = p1.Val + carry
		carry = s / 10
		s = s % 10

		p3.Next = &ListNode{Val: s}
		p3 = p3.Next
		p1 = p1.Next
	}
	if carry > 0 {
		p3.Next = &ListNode{Val: carry}
	}
	return head.Next
}
