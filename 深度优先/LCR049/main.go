package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode, cur int)
	dfs = func(node *TreeNode, cur int) {
		if node.Left == nil && node.Right == nil {
			ans += cur + node.Val
			return
		}
		if node.Left != nil {
			dfs(node.Left, (cur+node.Val)*10)
		}
		if node.Right != nil {
			dfs(node.Right, (cur+node.Val)*10)
		}
	}
	dfs(root, 0)
	return ans
}

func main() {
	fmt.Println(sumNumbers(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}))
}
