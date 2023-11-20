package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goNext(root *TreeNode, mark [10]bool) int {
	mark[root.Val] = !mark[root.Val]
	if root.Left == nil && root.Right == nil {
		cnt := 0
		for i := 1; i <= 9; i++ {
			if mark[i] {
				cnt++
			}
		}
		if cnt > 1 {
			return 0
		} else {
			return 1
		}
	}
	res := 0
	if root.Left != nil {
		res += goNext(root.Left, mark)
	}
	if root.Right != nil {
		res += goNext(root.Right, mark)
	}
	return res
}

func pseudoPalindromicPaths(root *TreeNode) int {
	return goNext(root, [10]bool{})
}

func main() {
	root := &TreeNode{2,
		&TreeNode{3, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}},
		&TreeNode{1, nil, &TreeNode{1, nil, nil}}}
	fmt.Println(pseudoPalindromicPaths(root))
}
