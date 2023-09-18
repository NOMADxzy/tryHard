package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var cur []*TreeNode
	var next []*TreeNode
	var lastVal int

	cur = append(cur, root)

	for len(cur) > 0 {
		node := cur[0]
		cur = cur[1:]
		lastVal = node.Val

		if node.Right != nil {
			next = append(next, node.Right)
		}
		if node.Left != nil {
			next = append(next, node.Left)
		}

		if len(cur) == 0 {
			cur = next
			next = nil
		}

	}
	return lastVal
}

func main() {
	root := &TreeNode{2,
		&TreeNode{1, nil, nil},
		&TreeNode{3, nil, nil}}
	fmt.Println(findBottomLeftValue(root))
}
