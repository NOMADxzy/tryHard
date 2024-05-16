package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var ans int
	curMaxLayer := -1
	var dfs func(root *TreeNode, layer int)
	dfs = func(root *TreeNode, layer int) {
		if root == nil {
			return
		}
		if layer > curMaxLayer {
			ans = root.Val
			curMaxLayer = layer
		}
		dfs(root.Left, layer+1)
		dfs(root.Right, layer+1)
	}
	dfs(root, 0)
	return ans
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	fmt.Println(findBottomLeftValue(root))
}
