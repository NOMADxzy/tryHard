package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(root *TreeNode) (int, int)
	best := 0
	dfs = func(root *TreeNode) (int, int) {
		childMin, childMax := root.Val, root.Val
		if root.Left != nil {
			min1, max1 := dfs(root.Left)
			childMin = min(childMin, min1)
			childMax = max(childMax, max1)
		}
		if root.Right != nil {
			min2, max2 := dfs(root.Right)
			childMin = min(childMin, min2)
			childMax = max(childMax, max2)
		}
		d1, d2 := max(root.Val-childMin, childMin-root.Val), max(root.Val-childMax, childMax-root.Val)
		best = max(best, max(d1, d2))
		return childMin, childMax
	}
	dfs(root)
	return best
}

func main() {
	root := &TreeNode{1,
		&TreeNode{4, nil, nil},
		&TreeNode{5, &TreeNode{1, nil, nil}, nil}}
	fmt.Println(maxAncestorDiff(root))
}
