package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if low > root.Val {
		return trimBST(root.Right, low, high)
	} else if high < root.Val {
		return trimBST(root.Left, low, high)
	} else {
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
		return root
	}
}

func main() {
	root := &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, nil}}
	root = trimBST(root, 10, 20)
	fmt.Println(root.Val)
}
