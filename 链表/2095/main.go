package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	var dfs func(head *ListNode, pre *ListNode, preCnt int) int
	dfs = func(head *ListNode, pre *ListNode, preCnt int) int {
		if head == nil {
			return 0
		}
		nextCnt := dfs(head.Next, head, preCnt+1)
		if nextCnt < 0 {
			return nextCnt
		}
		if preCnt == nextCnt || preCnt == nextCnt+1 {
			pre.Next = head.Next
			return -1
		}
		return nextCnt + 1
	}
	if head.Next == nil {
		return nil
	}
	dfs(head.Next, head, 1)
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	head = deleteMiddle(head)
	fmt.Println(head)
}
