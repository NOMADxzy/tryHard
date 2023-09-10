package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func midOrder(root *TreeNode, init *bool, pre *int, min *int, done *bool) {
	if *done {
		return
	}
	if root.Left != nil {
		midOrder(root.Left, init, pre, min, done)
	}

	if !*init {
		*pre = root.Val
		*init = true
	} else if *min < 0 {
		*min = root.Val - *pre
	} else if root.Val-*pre < *min {
		*min = root.Val - *pre
		if *min == 0 {
			*done = true
		}
	}
	*pre = root.Val

	if root.Right != nil {
		midOrder(root.Right, init, pre, min, done)
	}
}

func getMinimumDifference(root *TreeNode) int {
	done := false
	pre := 0
	init := false
	min := -1

	midOrder(root, &init, &pre, &min, &done)
	return min
}

func main() {
	root := &TreeNode{1, &TreeNode{-3, nil, nil}, &TreeNode{48, &TreeNode{12, nil, nil}, &TreeNode{49, nil, nil}}}
	fmt.Println(getMinimumDifference(root))
}
