package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func downCheck(arr []*ListNode, i int) {
	left := 2*i + 1
	right := 2*i + 2
	smallest := i
	if left < len(arr) && arr[left].Val < arr[smallest].Val {
		smallest = left
	}
	if right < len(arr) && arr[right].Val < arr[smallest].Val {
		smallest = right
	}
	if smallest != i {
		arr[smallest], arr[i] = arr[i], arr[smallest]
		downCheck(arr, smallest)
	}
}

func upCheck(arr []*ListNode, i int) {
	f := (i - 1) / 2
	if f >= 0 && arr[f].Val > arr[i].Val {
		arr[f], arr[i] = arr[i], arr[f]
		upCheck(arr, f)
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	m := len(lists)
	var ptrs []*ListNode
	for i := 0; i < m; i++ {
		if lists[i] != nil {
			ptrs = append(ptrs, lists[i])
			upCheck(ptrs, len(ptrs)-1)
		}
	}
	ansHead := &ListNode{}
	cur := ansHead
	for len(ptrs) > 0 {
		cur.Next = ptrs[0]
		cur = cur.Next

		ptrs[0] = ptrs[0].Next
		if ptrs[0] == nil {
			ptrs[0] = ptrs[len(ptrs)-1]
			ptrs = ptrs[:len(ptrs)-1]
		}
		downCheck(ptrs, 0)
	}
	return ansHead.Next
}

func main() {
	l1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l3 := &ListNode{2, &ListNode{6, nil}}
	fmt.Println(mergeKLists([]*ListNode{l1, l2, l3, nil}).Val)
}
