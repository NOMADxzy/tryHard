package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	root := &TreeNode{preorder[0], nil, nil}
	if len(preorder) == 1 {
		return root
	} else {
		var l int
		for l = 1; postorder[l-1] != preorder[1]; l++ {
		}
		root.Left = constructFromPrePost(preorder[1:l+1], postorder[0:l])
		if l+1 < len(preorder) {
			root.Right = constructFromPrePost(preorder[l+1:], postorder[l:len(postorder)-1])
		}
		return root
	}
}

func main() {
	root := constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1})
	fmt.Println(root.Val)
}
