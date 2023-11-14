package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root2 == root1
	} else if root1.Val != root2.Val {
		return false
	}
	return flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) ||
		flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
}

func main() {
	root1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	root2 := &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}
	fmt.Println(flipEquiv(root1, root2))
}
