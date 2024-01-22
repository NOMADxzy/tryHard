package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	minDistance, maxDistance := 1<<31, 0
	var dfs func(head *ListNode, pre, i int) (int, int)
	dfs = func(head *ListNode, pre int, i int) (int, int) {
		if head.Next == nil {
			return 1 << 31, -1
		}
		leftidx, rightidx := dfs(head.Next, head.Val, i+1)
		if head.Val < head.Next.Val && head.Val < pre || head.Val > head.Next.Val && head.Val > pre {
			if rightidx == -1 {
				leftidx = i
				rightidx = i
			} else {
				minDistance = min(minDistance, leftidx-i)
				leftidx = i
				maxDistance = max(maxDistance, rightidx-i)
			}
		}
		return leftidx, rightidx
	}
	if head.Next != nil {
		dfs(head.Next, head.Val, 1)
	}
	if maxDistance <= 0 {
		return []int{-1, -1}
	}
	return []int{minDistance, maxDistance}
}

func main() {
	head := &ListNode{5, &ListNode{3, &ListNode{1, &ListNode{2,
		&ListNode{5, &ListNode{1, &ListNode{2, nil}}}}}}}
	fmt.Println(nodesBetweenCriticalPoints(head))
}
