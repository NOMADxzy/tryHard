package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func find(root *TreeNode, max int) int {
	cnt := 0
	if max <= root.Val {
		cnt++
		max = root.Val
	}
	if root.Left != nil {
		cnt += find(root.Left, max)
	}
	if root.Right != nil {
		cnt += find(root.Right, max)
	}
	return cnt
}

func goodNodes(root *TreeNode) int {
	return find(root, root.Val)
}

func main() {
	root := &TreeNode{3,
		&TreeNode{1,
			&TreeNode{3, nil, nil}, nil},
		&TreeNode{4,
			&TreeNode{1, nil, nil},
			&TreeNode{5, nil, nil}}}
	fmt.Println(goodNodes(root))
}
