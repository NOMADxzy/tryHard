package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathTarget(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	var dfs func(root *TreeNode, path *[]int, acc int)
	dfs = func(root *TreeNode, path *[]int, acc int) {
		if root.Left == nil && root.Right == nil {
			if acc+root.Val == target {
				tmp := make([]int, len(*path))
				copy(tmp, *path)
				tmp = append(tmp, root.Val)
				ans = append(ans, tmp)
			}
			return
		}
		*path = append(*path, root.Val)
		if root.Left != nil {
			dfs(root.Left, path, acc+root.Val)
		}
		if root.Right != nil {
			dfs(root.Right, path, acc+root.Val)
		}
		*path = (*path)[:len(*path)-1]
	}
	var path []int
	dfs(root, &path, 0)
	return ans
}

func main() {
	root := &TreeNode{5,
		&TreeNode{3, nil, nil},
		&TreeNode{5, nil, nil}}
	fmt.Println(pathTarget(root, 8))
}
