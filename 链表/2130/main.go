package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow, fast := head, head
	var stack []int
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		stack = append(stack, slow.Val)
		slow = slow.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	ans := 0
	for slow != nil {
		ans = max(ans, slow.Val+stack[len(stack)-1])
		stack = stack[:len(stack)-1]
		slow = slow.Next
	}
	return ans
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Println(pairSum(head))
}
