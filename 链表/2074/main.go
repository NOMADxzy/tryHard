package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	var reverseList func(root *ListNode) (*ListNode, *ListNode)
	reverseList = func(root *ListNode) (*ListNode, *ListNode) {
		if root.Next == nil {
			return root, root
		}
		nextHead, nextEnd := reverseList(root.Next)
		nextEnd.Next = root
		root.Next = nil
		return nextHead, root
	}
	var dfs func(head *ListNode, total int)
	dfs = func(head *ListNode, total int) {
		if head.Next == nil {
			return
		}
		cur := head
		nextPos := head.Next
		var i int
		for i = 1; i <= total+1 && nextPos != nil; i++ {
			if i == total+1 && i%2 == 0 {
				nextPos = nextPos.Next
				cur = cur.Next
				cur.Next = nil
			} else {
				cur = cur.Next
				nextPos = nextPos.Next
			}
		}
		if i%2 == 1 {
			start, end := reverseList(head.Next)
			end.Next = nextPos
			head.Next = start
			dfs(end, total+1)
		} else {
			dfs(cur, total+1)
		}
	}
	dfs(head, 1)
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5,
		&ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, nil}}}}}}}}}
	//head := &ListNode{1, &ListNode{1, &ListNode{0, &ListNode{6, &ListNode{5, nil}}}}}
	head = reverseEvenLengthGroups(head)
	fmt.Println(head)
}
