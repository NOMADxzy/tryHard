package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var vals []int
	var cur []*TreeNode
	var next []*TreeNode

	cur = append(cur, root)
	max := root.Val

	for len(cur) > 0 {
		node := cur[0]
		cur = cur[1:]
		if node.Val > max {
			max = node.Val
		}

		if node.Right != nil {
			next = append(next, node.Right)
		}
		if node.Left != nil {
			next = append(next, node.Left)
		}

		if len(cur) == 0 {
			vals = append(vals, max)
			if len(next) == 0 {
				break
			}
			max = next[0].Val
			cur = next
			next = nil
		}

	}
	return vals
}

func main() {
	root := &TreeNode{2,
		&TreeNode{1, nil, nil},
		&TreeNode{3, nil, nil}}
	fmt.Println(largestValues(root))
}
