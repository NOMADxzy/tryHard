package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insert0(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var dfs func(head *ListNode) *ListNode
	dfs = func(head *ListNode) *ListNode {
		if head.Next == nil {
			return head
		}
		newNode := &ListNode{0, dfs(head.Next)}
		head.Next = newNode
		return head
	}
	return dfs(head)
}
