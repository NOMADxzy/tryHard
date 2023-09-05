package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var queue []*TreeNode
	var layer []int
	var scene []int

	if root == nil {
		return scene
	}
	queue = append(queue, root)
	layer = append(layer, 1)

	for len(queue) > 0 {
		node := queue[0]
		l := layer[0]
		layer = layer[1:]
		queue = queue[1:]
		val, left, right := node.Val, node.Left, node.Right
		if l > len(scene) {
			scene = append(scene, val)
		}
		if right != nil {
			queue = append(queue, right)
			layer = append(layer, l+1)
		}
		if left != nil {
			queue = append(queue, left)
			layer = append(layer, l+1)
		}

	}
	return scene
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}
	fmt.Println(rightSideView(root))
}
