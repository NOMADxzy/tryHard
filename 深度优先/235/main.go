package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func find(root *TreeNode, ptr **TreeNode, val1 int, val2 int, done *bool) int {
	if root == nil || *done {
		return 0
	} else {
		l := find(root.Left, ptr, val1, val2, done)
		r := find(root.Right, ptr, val1, val2, done)
		cnt := l + r
		if root.Val == val1 || root.Val == val2 {
			cnt++
		}
		if cnt == 2 && !*done {
			*ptr = root
			*done = true
		}
		return cnt
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ptr *TreeNode
	done := false
	find(root, &ptr, p.Val, q.Val, &done)
	return ptr
}

func main() {
	root := &TreeNode{6,
		&TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{8, nil, nil}}
	node := lowestCommonAncestor(root, &TreeNode{0, nil, nil}, &TreeNode{4, nil, nil})
	fmt.Println(node.Val)
}
