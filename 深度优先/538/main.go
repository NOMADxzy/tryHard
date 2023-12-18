package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, acc int) int {
	if root == nil {
		return 0
	}
	rightSum := dfs(root.Right, acc)
	leftSum := dfs(root.Left, acc+root.Val+rightSum)
	preVal := root.Val
	root.Val = rightSum + root.Val + acc
	return leftSum + rightSum + preVal
}

func convertBST(root *TreeNode) *TreeNode {
	dfs(root, 0)
	return root
}

func main() {
	root := &TreeNode{4,
		&TreeNode{1,
			&TreeNode{0, nil, nil},
			&TreeNode{2, nil, nil}},
		&TreeNode{6,
			&TreeNode{5, nil, nil},
			&TreeNode{7, nil, nil}}}
	root1 := convertBST(root)
	fmt.Println(root1.Val)
}

//4
//1   6
//0 2 5 7
