package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findNext(root *TreeNode, limit int, acc int) bool {
	ok := false
	if root.Left == nil && root.Right == nil {
		return acc >= limit
	} else {
		ok = false
		if root.Left != nil && findNext(root.Left, limit, acc+root.Left.Val) {
			ok = true
		} else {
			root.Left = nil
		}
		if root.Right != nil && findNext(root.Right, limit, acc+root.Right.Val) {
			ok = true
		} else {
			root.Right = nil
		}
	}
	return ok
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if !findNext(root, limit, root.Val) {
		return nil
	}
	return root
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}},
			&TreeNode{-99, &TreeNode{-99, nil, nil}, &TreeNode{-99, nil, nil}}},
		&TreeNode{3,
			&TreeNode{-99, &TreeNode{12, nil, nil}, &TreeNode{13, nil, nil}},
			&TreeNode{7, &TreeNode{-99, nil, nil}, &TreeNode{14, nil, nil}}}}
	r := sufficientSubset(root, 1)
	fmt.Println(r.Val)
}
