package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	} else {
		l, lc := check(root.Left)
		r, rc := check(root.Right)
		max := l
		cur_check := l-r <= 1
		if r > l {
			max = r
			cur_check = r-l <= 1
		}
		return max + 1, lc && rc && cur_check
	}
}

func isBalanced(root *TreeNode) bool {
	_, ch := check(root)
	return ch
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}}
	fmt.Println(isBalanced(root))
}
