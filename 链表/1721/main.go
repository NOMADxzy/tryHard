package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	var leftK, rightK *ListNode
	pre, cur := head, head
	for i := 1; i < k; i++ {
		cur = cur.Next
	}
	leftK = cur
	for cur.Next != nil {
		cur = cur.Next
		pre = pre.Next
	}
	rightK = pre
	leftK.Val, rightK.Val = rightK.Val, leftK.Val
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head = swapNodes(head, 1)
	fmt.Println(head)
}
