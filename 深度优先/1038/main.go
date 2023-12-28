package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {

	var dfs func(root *TreeNode, preSum int) int
	dfs = func(root *TreeNode, preSum int) int {
		if root == nil {
			return 0 + preSum
		}
		newSum := dfs(root.Right, preSum)
		root.Val += newSum
		newSum = root.Val
		return dfs(root.Left, newSum)
	}
	dfs(root, 0)
	return root
}

func main() {
	root := &TreeNode{1,
		&TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}},
		&TreeNode{1, nil, nil}}
	bstToGst(root)
	fmt.Println(root.Val)
}
