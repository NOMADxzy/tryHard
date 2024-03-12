package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodeIdx := map[*Node]int{}
	idxNode := map[int]*Node{}

	ans := &Node{}
	q := ans

	p := head
	idx := 0
	for p != nil {
		nodeIdx[p] = idx
		newNode := &Node{p.Val, nil, nil}
		q.Next = newNode
		q = q.Next
		idxNode[idx] = newNode

		idx++
		p = p.Next
	}

	p = head
	for r := ans.Next; r != nil; r = r.Next {
		if p.Random != nil {
			randomIdx := nodeIdx[p.Random]
			r.Random = idxNode[randomIdx]
		}
		p = p.Next
	}
	return ans.Next
}

func main() {
	fmt.Println()
}
