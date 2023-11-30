package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumNext(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	val := root.Val
	leftVal := SumNext(root.Left, ans)
	rightVal := SumNext(root.Right, ans)
	cur_root_best_path := val
	if leftVal > 0 {
		cur_root_best_path += leftVal
	}
	if rightVal > 0 {
		cur_root_best_path += rightVal
	}
	if cur_root_best_path > *ans {
		*ans = cur_root_best_path
	}
	if max(leftVal, rightVal) > 0 {
		val += max(leftVal, rightVal)
	}
	return val
}

func maxPathSum(root *TreeNode) int {
	ans := root.Val
	SumNext(root, &ans)
	return ans
}

func main() {
	root := &TreeNode{-10,
		&TreeNode{9, nil, nil},
		&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(maxPathSum(root))
}
