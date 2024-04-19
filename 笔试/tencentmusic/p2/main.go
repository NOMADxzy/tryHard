package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func create(n int) *TreeNode {
	sum := 1 << (n - 1)
	var build func(isRight bool, layer, curLayerCount int) *TreeNode

	build = func(isRight bool, layer, curLayerCount int) *TreeNode {
		if layer == n {
			return &TreeNode{1, nil, nil}
		}
		node := &TreeNode{1, nil, nil}
		node.Left = build(false, layer+1, curLayerCount*2)
		if isRight {
			node.Val = sum - curLayerCount + 1
			node.Right = build(true, layer+1, curLayerCount*2)
		} else {
			node.Right = build(false, layer+1, curLayerCount*2)
		}
		return node
	}

	return build(true, 1, 1)
}

func main() {
	root := create(3)
	fmt.Println(root.Val)
}
