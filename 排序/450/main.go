package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLast(root *TreeNode, father *TreeNode) *TreeNode {

	if root.Right != nil {
		return findLast(root.Right, root)
	} else {
		if root.Left == nil {
			if root.Val > father.Val {
				father.Right = nil
			} else {
				father.Left = nil
			}
			return root
		} else {
			if root.Left.Val > father.Val {
				father.Right = root.Left
			} else {
				father.Left = root.Left
			}
			return root
		}
	}
}

func findNode(root *TreeNode, isLeft bool, father *TreeNode, key int) {
	if root == nil {
		return
	}
	if root.Val == key {
		if root.Left != nil {
			node := findLast(root.Left, root)
			node.Left = root.Left
			node.Right = root.Right
			if father != nil {
				if isLeft {
					father.Left = node
				} else {
					father.Right = node
				}
			}
		} else if root.Right != nil {
			if father != nil {
				if isLeft {
					father.Left = root.Right
				} else {
					father.Right = root.Right
				}
			}
		} else {
			if isLeft {
				father.Left = nil
			} else {
				father.Right = nil
			}
		}
	} else if key < root.Val {
		findNode(root.Left, true, root, key)
	} else {
		findNode(root.Right, false, root, key)
	}
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	head := &TreeNode{0, root, nil}
	findNode(root, true, head, key)
	return head.Left
}

func main() {
	root := &TreeNode{5,
		&TreeNode{3,
			&TreeNode{2, nil, nil},
			&TreeNode{4, nil, nil}},
		&TreeNode{6, nil, &TreeNode{7, nil, nil}}}
	mroot := deleteNode(root, 2)
	fmt.Println(mroot.Val)
}
