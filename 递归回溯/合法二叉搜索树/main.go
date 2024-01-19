package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var check func(root *TreeNode) (bool, int, int)
	check = func(root *TreeNode) (bool, int, int) {
		ok := true
		minVal, maxVal := root.Val, root.Val
		if root.Left != nil {
			lok, lmin, lmax := check(root.Left)
			if !lok || lmax >= root.Val {
				ok = false
			} else {
				minVal = lmin
			}
		}
		if root.Right != nil {
			rok, rmin, rmax := check(root.Right)
			if !rok || rmin <= root.Val {
				ok = false
			} else {
				maxVal = rmax
			}
		}
		return ok, minVal, maxVal
	}
	if root == nil {
		return true
	}
	ans, _, _ := check(root)
	return ans
}

func main() {
	root := &TreeNode{5,
		&TreeNode{1, nil, nil},
		&TreeNode{9, &TreeNode{6, nil, nil}, &TreeNode{10, nil, nil}}}
	fmt.Println(isValidBST(root))
}
