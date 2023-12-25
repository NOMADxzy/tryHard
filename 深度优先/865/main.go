package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfsDepAndPlace(root *TreeNode) (int, *TreeNode) {
	if root.Left == nil && root.Right == nil {
		return 1, root
	}
	var ldep, rdep int
	var ltree, rtree *TreeNode
	if root.Left != nil {
		ldep, ltree = dfsDepAndPlace(root.Left)
	}
	if root.Right != nil {
		rdep, rtree = dfsDepAndPlace(root.Right)
	}
	if ldep == rdep {
		return ldep + 1, root
	} else if ldep > rdep {
		return ldep + 1, ltree
	} else {
		return rdep + 1, rtree
	}
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, ans := dfsDepAndPlace(root)
	return ans
}

func main() {
	root := &TreeNode{1,
		&TreeNode{1, &TreeNode{1, nil, nil}, nil},
		&TreeNode{1, nil, nil}}
	tree := subtreeWithAllDeepest(root)
	fmt.Println(tree.Val)
}
