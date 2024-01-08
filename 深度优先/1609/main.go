package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs 82% 5%
func isEvenOddTree1(root *TreeNode) bool {
	var dfs func(root *TreeNode, layer int)
	var layerVal []int
	ans := true
	dfs = func(root *TreeNode, layer int) {
		if !ans || root == nil {
			return
		}
		if layer%2 == 0 && root.Val%2 == 1 {
			if layer == len(layerVal) {
				layerVal = append(layerVal, root.Val)
			} else if root.Val <= layerVal[layer] {
				ans = false
			} else {
				layerVal[layer] = root.Val
			}
		} else if layer%2 == 1 && root.Val%2 == 0 {
			if layer == len(layerVal) {
				layerVal = append(layerVal, root.Val)
			} else if root.Val >= layerVal[layer] {
				ans = false
			} else {
				layerVal[layer] = root.Val
			}
		} else {
			ans = false
		}
		if ans {
			dfs(root.Left, layer+1)
			dfs(root.Right, layer+1)
		} else {
			return
		}
	}
	dfs(root, 0)
	return ans
}

// bfs 35% 50%
func isEvenOddTree(root *TreeNode) bool {
	var queue []*TreeNode
	type0 := true
	queue = append(queue, root)
	for len(queue) > 0 {
		var newQueue []*TreeNode
		cur := -1 << 31
		if !type0 {
			cur = -cur
		}
		for len(queue) > 0 {
			e := queue[0]
			queue = queue[1:]
			if type0 && e.Val > cur && e.Val%2 == 1 {
				cur = e.Val
			} else if !type0 && e.Val < cur && e.Val%2 == 0 {
				cur = e.Val
			} else {
				return false
			}
			if e.Left != nil {
				newQueue = append(newQueue, e.Left)
			}
			if e.Right != nil {
				newQueue = append(newQueue, e.Right)
			}
		}
		queue = newQueue
		type0 = !type0
	}
	return true
}

func main() {
	root := &TreeNode{1,
		&TreeNode{10, &TreeNode{3, nil, nil}, nil},
		&TreeNode{4, &TreeNode{7, nil, nil}, &TreeNode{5, nil, nil}}}
	fmt.Println(isEvenOddTree1(root))
}
