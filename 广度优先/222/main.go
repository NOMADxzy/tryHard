package main

import "fmt"

type TreeNode struct {
	int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var queue []*TreeNode
	if root == nil {
		return 0
	}

	queue = append(queue, root)
	bad := 0
	sum := 1
	curLayerNums := 1
	j := 0

	for len(queue) > 0 {

		if j == curLayerNums {
			j = 0
			curLayerNums *= 2
			sum += curLayerNums
			if bad > 0 {
				return sum - bad
			}
		}

		node := queue[0]
		queue = queue[1:]

		badLeft := 1
		badRight := 1
		if node.Right != nil {
			queue = append(queue, node.Right)
			badRight = 0
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			badLeft = 0
		}

		bad = bad + badLeft + badRight
		j++

	}
	return sum
}

func main() {
	fmt.Println(countNodes(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, nil}}))
}
