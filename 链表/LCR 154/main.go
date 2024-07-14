package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodeMap := map[*Node]*Node{}
	p := head
	for p != nil {
		newNode := &Node{Val: p.Val}
		nodeMap[p] = newNode
		p = p.Next
	}
	p = head
	for p != nil {
		nodeMap[p].Next = nodeMap[p.Next]
		nodeMap[p].Random = nodeMap[p.Random]
		p = p.Next
	}
	return nodeMap[head]
}

func main() {
	fmt.Println(copyRandomList(nil))
}
