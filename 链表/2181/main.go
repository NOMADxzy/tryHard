package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	var p, q *ListNode
	q = &ListNode{0, nil}
	newHead := &ListNode{-1, nil}
	pre := newHead
	p = head
	for p != nil {
		if p.Val == 0 {
			p = p.Next
		} else {
			for p.Val != 0 {
				q.Val += p.Val
				p = p.Next
			}
			pre.Next = q
			pre = q
			q = &ListNode{0, nil}
		}
	}
	return newHead.Next
}

func main() {
	head := &ListNode{0, &ListNode{3, &ListNode{1, &ListNode{0, nil}}}}
	fmt.Println(mergeNodes(head).Val)
}
