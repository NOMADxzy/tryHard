package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var stack []*TreeNode
	cur := root
	ok := false
	for {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			if len(stack) == 0 {
				break
			}
			e := stack[len(stack)-1]
			if e == p {
				ok = true
			} else if ok {
				return e
			}
			stack = stack[:len(stack)-1]
			cur = e.Right
		}
	}
	return nil
}

// 1
// 2 3
// 4
func main() {
	root := &TreeNode{1,
		&TreeNode{2, nil, &TreeNode{4, nil, nil}},
		&TreeNode{3, nil, nil}}
	fmt.Println(inorderSuccessor(root, root.Left).Val)
}
