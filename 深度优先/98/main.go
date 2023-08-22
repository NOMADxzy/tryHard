package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(root *TreeNode, upper int, lower int) bool {
	if root == nil {
		return true
	} else if root.Val < upper && root.Val > lower {
		leftOk := check(root.Left, root.Val, lower)
		rightOk := check(root.Right, upper, root.Val)

		if leftOk && rightOk {
			return true
		}
	}
	return false
}

func isValidBST(root *TreeNode) bool {
	return check(root, 1<<31, -(1<<31 + 1)) // 根节点 正无穷和负无穷
}
func main() {
	tree := &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}}
	fmt.Println(isValidBST(tree))
}
