package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var midOrder func(root *TreeNode)
	var ans *TreeNode
	var find bool
	midOrder = func(root *TreeNode) {
		if root == nil || ans != nil {
			return
		}
		midOrder(root.Left)

		if find && ans == nil {
			ans = root
		}
		if root == p || root.Val == p.Val {
			find = true
		}
		midOrder(root.Right)
	}
	midOrder(root)
	return ans
}

func main() {
	root := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(inorderSuccessor(root, &TreeNode{3, nil, nil}))
}
