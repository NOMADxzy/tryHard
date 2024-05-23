package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	var dfs func(head *ListNode, x int) (*ListNode, *ListNode)
	dfs = func(head *ListNode, x int) (*ListNode, *ListNode) {
		if head.Next == nil {
			return head, head
		}
		start, end := dfs(head.Next, x)
		if head.Val < x {
			head.Next = start
			return head, end
		} else {
			head.Next = end.Next
			end.Next = head
			return start, end
		}
	}
	ans, _ := dfs(head, x)
	return ans
}

func main() {
	head := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{
		2, &ListNode{5, &ListNode{2, nil}}}}}}
	partition(head, 3)
	fmt.Println(head)
}
