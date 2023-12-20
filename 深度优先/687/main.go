package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findNext(root *TreeNode, ans *int) int {
	l, r := 0, 0
	if root.Left != nil {
		leftChildMaxLen := findNext(root.Left, ans)
		if root.Val == root.Left.Val {
			l = leftChildMaxLen + 1
		}
	}
	if root.Right != nil {
		rightChildMaxLen := findNext(root.Right, ans)
		if root.Val == root.Right.Val {
			r = rightChildMaxLen + 1
		}
	}
	if l+r > *ans {
		*ans = l + r
	}
	return max(l, r)
}

func longestUnivaluePath(root *TreeNode) int {
	var ans int
	if root != nil {
		findNext(root, &ans)
	}
	return ans
}

func main() {
	root := &TreeNode{5,
		&TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}},
		&TreeNode{5, nil, &TreeNode{5, nil, nil}}}
	fmt.Println(longestUnivaluePath(root))
}
