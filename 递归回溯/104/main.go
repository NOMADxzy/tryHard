package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		l1 := maxDepth(root.Left)
		l2 := maxDepth(root.Right)
		if l2 > l1 {
			l1 = l2
		}
		return l1 + 1
	}
}

func main() {
	root := &TreeNode{2, &TreeNode{2, nil, nil}, &TreeNode{2, &TreeNode{2, nil, nil}, nil}}
	fmt.Println(maxDepth(root))
}
