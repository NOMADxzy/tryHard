package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reCombine(head *ListNode, leftLast **ListNode) (*ListNode, *ListNode) {
	if head != nil && head.Next != nil {
		left, right := reCombine(head.Next.Next, leftLast)
		if left != nil && left.Next == nil {
			*leftLast = left
		}
		head_ := head.Next
		head.Next = left
		head_.Next = right
		return head, head_
	} else {
		return head, nil
	}
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	leftLast := head
	left, right := reCombine(head, &leftLast)
	if leftLast != nil {
		(*leftLast).Next = right
	}
	return left
}

func main() {
	head := &ListNode{2, &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{6,
		&ListNode{4, &ListNode{7, nil}}}}}}}
	fmt.Println(oddEvenList(head).Val)
}
