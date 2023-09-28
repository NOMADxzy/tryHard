package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func SimpleAdd(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	node.Val = l1.Val + l2.Val
	if l1.Next != nil {
		nextNode := SimpleAdd(l1.Next, l2.Next)
		if nextNode.Val >= 10 {
			nextNode.Val -= 10
			node.Val++
		}
		node.Next = nextNode
	}
	return node
}

func appendTail(p *ListNode, k int, l2 *ListNode) *ListNode {
	if k == 0 {
		return SimpleAdd(p, l2)
	} else {
		node := &ListNode{p.Val, nil}
		nextNode := appendTail(p.Next, k-1, l2)
		if nextNode.Val >= 10 {
			nextNode.Val -= 10
			node.Val += 1
		}
		node.Next = nextNode
		return node
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	w1, w2 := 0, 0
	for p := l1; p != nil; p = p.Next {
		w1++
	}
	for p := l2; p != nil; p = p.Next {
		w2++
	}
	if w2 > w1 {
		l1, l2 = l2, l1
		w1, w2 = w2, w1
	}
	r := appendTail(l1, w1-w2, l2)
	if r.Val >= 10 {
		r.Val -= 10
		r = &ListNode{1, r}
	}
	return r

}

func main() {
	l2 := &ListNode{9, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}
	l1 := &ListNode{7, &ListNode{6, &ListNode{4, nil}}}
	fmt.Println(addTwoNumbers(l1, l2).Val)
}
