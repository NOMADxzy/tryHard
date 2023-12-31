package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	hist := map[*ListNode]map[int]bool{}
	var checkNext func(head *ListNode) *ListNode
	checkNext = func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}
		next := checkNext(head.Next)
		if next != nil && hist[next][-head.Val] || head.Val == 0 {
			sum := head.Val
			for sum != 0 {
				sum += next.Val
				next = next.Next
			}
			return next
		} else {
			newMap := map[int]bool{head.Val: true}
			for v, _ := range hist[next] {
				newMap[v+head.Val] = true
			}
			hist[head] = newMap
			head.Next = next
			return head
		}
	}
	return checkNext(head)
}

func main() {
	head := &ListNode{0, &ListNode{3, &ListNode{3, &ListNode{-3, &ListNode{-3, nil}}}}}
	h := removeZeroSumSublists(head)
	fmt.Println(h)
}
