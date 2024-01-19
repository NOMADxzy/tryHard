package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxValue(root *TreeNode, k int) int {
	var dfs func(root *TreeNode)
	empty := make([]int, k+1)
	hist := map[*TreeNode][]int{nil: empty}
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if hist[root.Left] == nil {
			dfs(root.Left)
		}
		if hist[root.Right] == nil {
			dfs(root.Right)
		}
		leftDP, rightDP := hist[root.Left], hist[root.Right]

		dp := make([]int, k+1)
		dp[0] = leftDP[k] + rightDP[k]
		preMax := dp[0]
		for n := 1; n <= k; n++ {
			for i := 0; i <= n-1; i++ {
				dp[n] = max(leftDP[i]+rightDP[n-1-i]+root.Val, dp[n])
			}
			if dp[n] > preMax {
				preMax = dp[n]
			} else {
				dp[n] = preMax
			}
			preMax = max(preMax, dp[n])
			dp[n] = preMax
		}
		hist[root] = dp
	}
	dfs(root)
	ans := 0
	for i := 0; i <= k; i++ {
		ans = max(ans, hist[root][i])
	}
	return ans
}

func main() {
	root := &TreeNode{5,
		&TreeNode{2, &TreeNode{4, nil, nil}, nil},
		&TreeNode{3, nil, nil}}
	fmt.Println(maxValue(root, 2))
}
