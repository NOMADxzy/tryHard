package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeBack(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	nextNode := removeBack(head.Next)
	if head.Val >= nextNode.Val {
		head.Next = nextNode
		return head
	} else {
		return nextNode
	}
}

func removeNodes(head *ListNode) *ListNode {
	return removeBack(head)
}

func main() {
	head := &ListNode{5, &ListNode{2, &ListNode{13, &ListNode{5, nil}}}}
	head = removeNodes(head)
	fmt.Println(head.Val)
}
