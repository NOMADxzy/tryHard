package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func canSorted(lists []*ListNode) []bool {
	nextDownNode := func(head *ListNode) *ListNode {
		p := head
		for p.Next != nil && p.Next.Val > p.Val {
			p = p.Next
		}
		return p
	}
	m := len(lists)
	ans := make([]bool, m)

	for idx, head := range lists {
		node := nextDownNode(head)
		if node.Next != nil {
			end := nextDownNode(node.Next)
			if end.Next != nil || end.Val >= head.Val {
				continue
			}
		}
		ans[idx] = true
	}
	return ans
}

func main() {
	lists := []*ListNode{
		{1, &ListNode{2, &ListNode{3, nil}}},
		{2, &ListNode{3, &ListNode{1, nil}}},
		{3, &ListNode{2, &ListNode{1, nil}}},
	}
	fmt.Println(canSorted(lists))
}
