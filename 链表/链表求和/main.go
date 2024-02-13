package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var add func(l1, l2 *ListNode, c int) *ListNode
	add = func(l1, l2 *ListNode, c int) *ListNode {
		var l1next, l2next *ListNode
		val := c
		if l1 == nil && l2 == nil && c == 0 {
			return nil
		}
		if l1 != nil {
			l1next = l1.Next
			val += l1.Val
		}
		if l2 != nil {
			l2next = l2.Next
			val += l2.Val
		}
		nextNode := add(l1next, l2next, val/10)
		curNode := &ListNode{val % 10, nextNode}
		return curNode
	}
	return add(l1, l2, 0)
}

func main() {
	l1 := &ListNode{7, nil}
	l2 := &ListNode{5, &ListNode{9, &ListNode{2, nil}}}
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3)
}
