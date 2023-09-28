package main

import "fmt"

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatChild(root *Node) (*Node, *Node) {
	p := root
	l := root
	for p != nil {
		if p.Child != nil {
			first, last := flatChild(p.Child)
			pNext := p.Next
			p.Next = first
			first.Prev = p
			last.Next = pNext
			if pNext != nil {
				pNext.Prev = last
			}
			p.Child = nil

			l = last
			p = pNext
		} else {
			l = p
			p = p.Next
		}
	}
	return root, l
}

func flatten(root *Node) *Node {
	r, _ := flatChild(root)
	return r
}

func main() {
	//root := &Node{1, nil, &Node{2, nil, &Node{3, nil, &Node{6, nil, nil, nil}, nil}, nil},
	//	&Node{4, nil, &Node{5, nil, nil, nil}, nil}}
	root := &Node{1, nil, nil, &Node{2, nil, nil, &Node{3, nil, nil, nil}}}
	fmt.Println(flatten(root))
}
