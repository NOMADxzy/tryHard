package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func midOrder(root *TreeNode, k int, m *int, i *int, done *bool) {
	if *done {
		return
	}
	if root.Left != nil {
		midOrder(root.Left, k, m, i, done)
	}
	*i++
	if *i == k {
		*m = root.Val
		*done = true
	}
	if root.Right != nil {
		midOrder(root.Right, k, m, i, done)
	}
}

func kthSmallest(root *TreeNode, k int) int {
	m := 0
	i := 0
	done := false
	midOrder(root, k, &m, &i, &done)
	return m
}

func main() {
	root := &TreeNode{3, &TreeNode{1, &TreeNode{5, nil, nil}, &TreeNode{2, nil, nil}},
		&TreeNode{4, nil, nil}}

	fmt.Println(kthSmallest(root, 3))
}
