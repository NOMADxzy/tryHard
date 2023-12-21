package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkNext(root *TreeNode) bool {
	if root == nil {
		return false
	}
	ok := false
	if root.Val == 1 {
		ok = true
	}
	if checkNext(root.Left) {
		ok = true
	} else {
		root.Left = nil
	}
	if checkNext(root.Right) {
		ok = true
	} else {
		root.Right = nil
	}
	return ok
}

func pruneTree(root *TreeNode) *TreeNode {
	if !checkNext(root) {
		return nil
	}
	return root
}

func main() {
	root := &TreeNode{1, nil,
		&TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}}}
	pruneTree(root)
	fmt.Println(root.Val)
}
