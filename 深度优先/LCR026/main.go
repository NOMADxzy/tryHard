package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList1(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

func reorderList(head *ListNode) {
	pre := map[*ListNode]*ListNode{}
	var p *ListNode
	for p = head; p.Next != nil; p = p.Next {
		pre[p.Next] = p
	}
	var dfs func(l, r *ListNode) *ListNode
	dfs = func(l, r *ListNode) *ListNode {
		if l == r || l.Next == r {
			r.Next = nil
			return l
		}
		nextL := l.Next
		l.Next = r
		nextNode := dfs(nextL, pre[r])
		r.Next = nextNode
		return l
	}
	dfs(head, p)
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	reorderList(head)
	fmt.Println(head)
}
