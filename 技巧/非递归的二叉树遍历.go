package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preOrder(root *TreeNode) []int {
	var stack []*TreeNode
	var ans []int
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return ans
}

// 后序遍历
func postOrder(root *TreeNode) []int {
	var stack []*TreeNode
	var ans []int
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}

// 中序遍历
func midOrder(root *TreeNode) []int {
	var stack []*TreeNode
	var ans []int
	cur := root
	for {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			if len(stack) == 0 {
				break
			}
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, node.Val)
			cur = node.Right
		}
	}
	return ans
}

//   1
// 2   3
//4 5 6 7

func main() {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil}},
		&TreeNode{3,
			&TreeNode{6, nil, nil},
			&TreeNode{7, nil, nil}}}
	fmt.Println(midOrder(root))
}
