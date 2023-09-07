package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, p int, q int, m **TreeNode) int {
	exist := 0
	left, right := root.Left, root.Right
	if left != nil {
		exist += dfs(left, p, q, m)
	}
	if right != nil {
		exist += dfs(right, p, q, m)
	}

	if root.Val == p || root.Val == q {
		exist += 1
	}
	if exist == 2 {
		*m = root
		return 3
	}

	return exist
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var m *TreeNode
	dfs(root, p.Val, q.Val, &m)
	return m
}

func main() {
	root := &TreeNode{3, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, nil, nil}},
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}}}

	fmt.Println(lowestCommonAncestor(root, &TreeNode{6, nil, nil}, &TreeNode{5, nil, nil}).Val)
}
