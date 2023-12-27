package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if root.Val > val {
		root.Right = insertIntoMaxTree(root.Right, val)
		return root
	} else {
		return &TreeNode{val, root, nil}
	}
}

func main() {
	root := &TreeNode{4,
		&TreeNode{1, nil, nil},
		&TreeNode{3, &TreeNode{2, nil, nil}, nil}}
	nroot := insertIntoMaxTree(root, 5)
	fmt.Println(nroot.Val)
}
