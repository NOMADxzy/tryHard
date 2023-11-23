package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseSimple(node *ListNode) (*ListNode, *ListNode) {
	if node == nil {
		return nil, nil
	}
	start, preNode := reverseSimple(node.Next)
	node.Next = nil
	if preNode != nil {
		preNode.Next = node
	}
	if start == nil {
		start = node
	}
	return start, node
}

func reverseNext(node *ListNode, k, pos int, pre, preLast, totalPre *ListNode) {
	if node == nil {
		start, _ := reverseSimple(pre)
		totalPre.Next = start
		return
	}
	nextNode := node.Next
	node.Next = pre
	if k == pos {
		totalPre.Next = node
		reverseNext(nextNode, k, 1, nil, nil, preLast)
	} else {
		if preLast == nil {
			preLast = node
		}
		reverseNext(nextNode, k, pos+1, node, preLast, totalPre)
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	totalPre := &ListNode{}
	reverseNext(head, k, 1, nil, nil, totalPre)
	return totalPre.Next
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(reverseKGroup(head, 3))
}
