package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		cnt := 1 + dfs(root.Left) + dfs(root.Right)
		return cnt
	}
	cnt2 := dfs(t2)
	var ans bool
	var compare func(root, node *TreeNode) bool
	compare = func(root *TreeNode, node *TreeNode) bool {
		if root == nil {
			return node == nil
		}
		if root.Val != node.Val {
			return false
		}
		if compare(root.Left, node.Left) && compare(root.Right, node.Right) {
			return true
		}
		return false
	}
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		cnt := 1 + dfs(root.Left) + dfs(root.Right)
		if cnt == cnt2 && compare(root, t2) {
			ans = true
		}
		return cnt
	}
	dfs(t1)
	return ans
}

func main() {
	t1 := &TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{2, &TreeNode{1, nil, nil}, nil}}
	t2 := &TreeNode{2, &TreeNode{1, nil, nil}, nil}
	fmt.Println(checkSubTree(t1, t2))
}
