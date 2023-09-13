package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	var cur []*TreeNode

	cur = append(cur, root)
	max := root.Val
	layer := 1
	largestLayer := 1

	for len(cur) > 0 {

		sum := 0

		var next []*TreeNode
		for len(cur) > 0 {
			node := cur[0]
			cur = cur[1:]

			sum += node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		if sum > max {
			max = sum
			largestLayer = layer
		}

		cur = next
		layer++

	}
	return largestLayer

}

func main() {
	root := &TreeNode{-100, &TreeNode{-200, &TreeNode{-20, nil, nil}, &TreeNode{-5, nil, nil}}, &TreeNode{-300, &TreeNode{-10, nil, nil}, nil}}
	fmt.Println(maxLevelSum(root))
}
