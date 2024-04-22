package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetMid(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func binSearch(head *ListNode, t int) *ListNode {
	if head == nil {
		return nil
	}
	left, right := head, head
	for right.Next != nil {
		right = right.Next
	}
	for left != right {
		mid := GetMid(left)
		if mid.Val >= t {
			right = mid
			right.Next = nil
		} else {
			left = mid.Next
		}
	}
	if right.Val == t {
		return right
	}
	return nil
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{4, &ListNode{5, nil}}}}
	ans := binSearch(head, 4)
	if ans != nil {
		fmt.Println(ans.Val)
	}
}
