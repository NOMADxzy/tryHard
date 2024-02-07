package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	var doubleNext func(head *ListNode) bool
	doubleNext = func(head *ListNode) bool {
		if head == nil {
			return false
		}
		full := doubleNext(head.Next)
		newVal := head.Val * 2
		if full {
			newVal++
		}
		head.Val = newVal % 10
		return newVal > 9
	}
	full := doubleNext(head)
	if full { // 易错， 需要考虑数字位数增加1的情况
		n := &ListNode{1, head}
		head = n
	}
	return head
}

func main() {
	head := &ListNode{9, &ListNode{9, &ListNode{9, nil}}}
	head = doubleIt(head)
	fmt.Println(head)
}
