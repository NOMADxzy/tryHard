package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var dfs func(head *ListNode) *ListNode
	dfs = func(head *ListNode) *ListNode {
		if head.Next == nil {
			return head
		}
		nexts := dfs(head.Next)
		if head.Val <= nexts.Val {
			head.Next = nexts
			return head
		} else {
			pre := nexts
			for pre.Next != nil && pre.Next.Val < head.Val {
				pre = pre.Next
			}
			head.Next = pre.Next
			pre.Next = head
			return nexts
		}
	}
	return dfs(head)
}

func main() {
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	newHead := sortList(head)
	fmt.Println(newHead.Val)
}
