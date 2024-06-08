package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	ans := 0
	var dfs func(node *TreeNode)
	hists := map[*TreeNode]map[int]int{}
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if hists[node] == nil {
			hists[node] = map[int]int{node.Val: 1}
		}
		if node.Val == targetSum {
			ans++
		}
		if node.Left != nil {
			dfs(node.Left)
			hist := hists[node.Left]
			for k, v := range hist {
				hists[node][k+node.Val] += v
				if k+node.Val == targetSum {
					ans += v
				}
			}
		}
		if node.Right != nil {
			dfs(node.Right)
			hist := hists[node.Right]
			for k, v := range hist {
				hists[node][k+node.Val] += v
				if k+node.Val == targetSum {
					ans += v
				}
			}
		}

	}
	dfs(root)
	return ans
}

func main() {
	root := &TreeNode{Val: 10,
		Left: &TreeNode{5, &TreeNode{3, &TreeNode{3, nil, nil}, &TreeNode{-2, nil, nil}}, &TreeNode{2, nil, &TreeNode{1, nil, nil}}}}
	fmt.Println(pathSum(root, 8))
}
