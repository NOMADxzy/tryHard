package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func find(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	} else {
		leftL := find(root.Left, max)
		rightL := find(root.Right, max)

		if leftL+rightL > *max {
			*max = leftL + rightL
		}

		better := leftL
		if rightL > better {
			better = rightL
		}
		return better + 1
	}
}

func diameterOfBinaryTree(root *TreeNode) int {
	longest := 0
	find(root, &longest)

	return longest
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{2, nil, nil}, nil}}
	fmt.Println(diameterOfBinaryTree(root))
}
