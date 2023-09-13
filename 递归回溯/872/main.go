package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostOrder(root *TreeNode, leafs *[]int) {
	if root.Left == nil && root.Right == nil {
		*leafs = append(*leafs, root.Val)
	} else {
		if root.Left != nil {
			PostOrder(root.Left, leafs)
		}
		if root.Right != nil {
			PostOrder(root.Right, leafs)
		}
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leafs1 []int
	var leafs2 []int
	PostOrder(root1, &leafs1)
	PostOrder(root2, &leafs2)
	if len(leafs2) != len(leafs1) {
		return false
	} else {
		for i := 0; i < len(leafs1); i++ {
			if leafs1[i] != leafs2[i] {
				return false
			}
		}
	}
	return true

}

func main() {
	root1 := &TreeNode{3,
		&TreeNode{5,
			&TreeNode{6, nil, nil},
			&TreeNode{2,
				&TreeNode{7, nil, nil},
				&TreeNode{4, nil, nil}}},
		&TreeNode{1,
			&TreeNode{9, nil, nil},
			&TreeNode{8, nil, nil}}}
	root2 := &TreeNode{3,
		&TreeNode{5,
			&TreeNode{6, nil, nil},
			&TreeNode{2,
				&TreeNode{7, nil, nil},
				&TreeNode{4, nil, nil}}},
		&TreeNode{1,
			&TreeNode{9, nil, nil},
			&TreeNode{9, nil, nil}}}
	fmt.Println(leafSimilar(root1, root2))
}
