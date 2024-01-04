package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	var dfs func(root *TreeNode) (int, int)
	ans := 0
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		lsum, lnsum := dfs(root.Left)
		rsum, rnsum := dfs(root.Right)
		if root.Val%2 == 0 {
			ans += lnsum + rnsum
		}
		return root.Val, lsum + rsum
	}
	dfs(root)
	return ans
}

func main() {
	root := &TreeNode{2,
		&TreeNode{2, &TreeNode{20, nil, nil}, nil},
		&TreeNode{2, nil, &TreeNode{20, nil, nil}}}
	fmt.Println(sumEvenGrandparent(root))
}
