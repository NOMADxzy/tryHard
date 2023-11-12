package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		root = &TreeNode{val, root, nil}
	} else if depth > 2 {
		addOneRow(root.Left, val, depth-1)
		addOneRow(root.Right, val, depth-1)
	} else {
		l, r := root.Left, root.Right
		root.Left = &TreeNode{val, l, nil}
		root.Right = &TreeNode{val, nil, r}
	}
	return root
}

func main() {
	root := &TreeNode{4,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}},
		&TreeNode{6, &TreeNode{5, nil, nil}, nil}}

	newRoot := addOneRow(root, 1, 1)
	fmt.Println(newRoot.Val)
}
