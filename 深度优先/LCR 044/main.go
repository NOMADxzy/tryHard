package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var maxVals []int
	update := func(val, layer int) {
		if layer >= len(maxVals) {
			maxVals = append(maxVals, val)
		} else if maxVals[layer] < val {
			maxVals[layer] = val
		}
	}
	var dfs func(root *TreeNode, layer int)
	dfs = func(root *TreeNode, layer int) {
		if root == nil {
			return
		}
		update(root.Val, layer)
		dfs(root.Left, layer+1)
		dfs(root.Right, layer+1)
	}
	dfs(root, 0)
	return maxVals
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil}}
	fmt.Println(largestValues(root))
}
