package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func find(root *TreeNode, toLeft bool, acc int) int {
	if root == nil {
		return acc
	}

	var inherit int
	var restart int
	if toLeft {
		inherit = find(root.Left, false, acc+1)
		restart = find(root.Right, true, 1)
	} else {
		inherit = find(root.Right, true, acc+1)
		restart = find(root.Left, false, 1)
	}

	return max(inherit, restart)
}

func longestZigZag(root *TreeNode) int {
	l := find(root.Left, false, 1)
	r := find(root.Right, true, 1)
	return max(l, r) - 1
}

func main() {
	root := &TreeNode{1,
		nil,
		&TreeNode{1,
			&TreeNode{1, nil, nil},
			&TreeNode{1,
				&TreeNode{1, nil,
					&TreeNode{1, nil, &TreeNode{1, nil, nil}}},
				&TreeNode{1, nil, nil}}}}
	fmt.Println(longestZigZag(root))
}
