package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(left *TreeNode, right *TreeNode) bool {
	if left == nil {
		return right == nil
	} else if right == nil {
		return false
	} else {
		if left.Val != right.Val {
			return false
		} else {
			return check(left.Left, right.Right) && check(left.Right, right.Left)
		}
	}
}

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}

	fmt.Println(isSymmetric(root))
}
