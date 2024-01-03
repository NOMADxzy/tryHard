package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxSumBST(root *TreeNode) int {
	ans := -1 << 31
	var checkNext func(node *TreeNode) (bool, int, int, int)
	checkNext = func(node *TreeNode) (bool, int, int, int) {
		ok := true
		sum := node.Val
		minVal, maxVal := node.Val, node.Val
		if node.Left != nil {
			leftOk, lmin, lmax, lsum := checkNext(node.Left)
			sum += lsum
			if !leftOk || lmax >= node.Val {
				ok = false
			} else {
				minVal = lmin
			}
		}
		if node.Right != nil {
			rightOk, rmin, rmax, rsum := checkNext(node.Right)
			sum += rsum
			if !rightOk || rmin <= node.Val {
				ok = false
			} else {
				maxVal = rmax
			}
		}
		if ok {
			ans = max(ans, sum)
			return true, minVal, maxVal, sum
		}
		return false, 0, 0, sum
	}
	checkNext(root)
	return max(ans, 0)
}

func main() {
	root := &TreeNode{1,
		&TreeNode{4,
			&TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{5, nil, nil}}}
	fmt.Println(maxSumBST(root))
}
