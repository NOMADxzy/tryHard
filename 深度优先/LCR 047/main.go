package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	var check func(root *TreeNode) bool
	check = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		leftAll0 := check(root.Left)
		rightAll0 := check(root.Right)
		if root.Val == 0 && leftAll0 && rightAll0 {
			return true
		} else {
			if leftAll0 {
				root.Left = nil
			}
			if rightAll0 {
				root.Right = nil
			}
		}
		return false
	}
	// TODO 考虑根节点删除情况
	if check(root) {
		return nil
	}
	return root
}

func main() {
	root := &TreeNode{0, nil,
		&TreeNode{0,
			&TreeNode{0, nil, nil},
			&TreeNode{0, nil, nil}}}
	newRoot := pruneTree(root)
	fmt.Println(newRoot)
}
