package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func midOrder(root *TreeNode, arr *[]int) {
	if root.Left != nil {
		midOrder(root.Left, arr)
	}
	*arr = append(*arr, root.Val)
	if root.Right != nil {
		midOrder(root.Right, arr)
	}
}

func constructBalanceTree(arr []int, start, end int) *TreeNode {
	mid := (end + start) / 2
	root := &TreeNode{arr[mid], nil, nil}
	if mid > start {
		root.Left = constructBalanceTree(arr, start, mid-1)
	}
	if mid < end {
		root.Right = constructBalanceTree(arr, mid+1, end)
	}
	return root
}

func balanceBST(root *TreeNode) *TreeNode {
	var arr []int
	midOrder(root, &arr)
	return constructBalanceTree(arr, 0, len(arr)-1)
}

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}}
	root = balanceBST(root)
	fmt.Println(root.Val)
}
