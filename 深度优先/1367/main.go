package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	var dfs func(head *ListNode, root *TreeNode, isHead bool) bool
	dfs = func(head *ListNode, root *TreeNode, isHead bool) bool {
		if head == nil {
			return true
		} else if root == nil {
			return false
		} else if head.Val == root.Val && (dfs(head.Next, root.Left, false) || dfs(head.Next, root.Right, false)) {
			return true
		}
		if isHead && (dfs(head, root.Left, true) || dfs(head, root.Right, true)) {
			return true
		}
		return false
	}
	return dfs(head, root, true)
}

func main() {
	head := &ListNode{4, &ListNode{2, &ListNode{8, nil}}}
	root := &TreeNode{1, nil,
		&TreeNode{4,
			&TreeNode{2, &TreeNode{6, nil, nil},
				&TreeNode{8, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}}},
			nil}}
	fmt.Println(isSubPath(head, root))
}
