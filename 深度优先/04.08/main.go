package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) int
	var ans *TreeNode
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		cnt := 0
		if root == p || root == q {
			cnt++
		}
		leftCnt := dfs(root.Left)
		rightCnt := dfs(root.Right)
		cnt += leftCnt + rightCnt
		if cnt == 2 && leftCnt <= 1 && rightCnt <= 1 {
			ans = root
		}
		return cnt
	}
	dfs(root)
	return ans
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right).Val)
}
