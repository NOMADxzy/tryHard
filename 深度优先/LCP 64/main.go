package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closeLampInTree(root *TreeNode) (ans int) {

	type tuple struct {
		node *TreeNode

		switch2, switch3 bool
	}

	dp := map[tuple]int{} // 记忆化搜索

	var f func(*TreeNode, bool, bool) int

	f = func(node *TreeNode, switch2, switch3 bool) int {

		if node == nil {

			return 0

		}

		p := tuple{node, switch2, switch3}

		if res, ok := dp[p]; ok {

			return res

		}

		if node.Val == 1 == (switch2 == switch3) { // 当前节点为开灯

			res1 := f(node.Left, switch2, false) + f(node.Right, switch2, false) + 1

			res2 := f(node.Left, !switch2, false) + f(node.Right, !switch2, false) + 1

			res3 := f(node.Left, switch2, true) + f(node.Right, switch2, true) + 1

			r123 := f(node.Left, !switch2, true) + f(node.Right, !switch2, true) + 3

			dp[p] = min(res1, res2, res3, r123)

		} else { // 当前节点为关灯

			res0 := f(node.Left, switch2, false) + f(node.Right, switch2, false)

			res12 := f(node.Left, !switch2, false) + f(node.Right, !switch2, false) + 2

			res13 := f(node.Left, switch2, true) + f(node.Right, switch2, true) + 2

			res23 := f(node.Left, !switch2, true) + f(node.Right, !switch2, true) + 2

			dp[p] = min(res0, res12, res13, res23)

		}

		return dp[p]

	}

	return f(root, false, false)

}

func min(a, b, c, d int) int {

	if b < a {

		a = b

	}

	if c < a {

		a = c

	}

	if d < a {

		a = d

	}

	return a

}

//func closeLampInTree1(root *TreeNode) int {
//	// 状态1：根开，其余闭
//	// 状态2：根闭，其余开
//	// 状态3：根开，其余开
//	// 状态4：根闭，其余闭
//	hist := map[int]map[*TreeNode]int{}
//	for i := 0; i < 4; i++ {
//		hist[i] = map[*TreeNode]int{}
//	}
//	var dfs func(root *TreeNode, target int, same bool) int
//	dfs = func(root *TreeNode, target int, same bool) int {
//		key := target
//		if same {
//			key += 2
//		}
//		if v, ok := hist[key][root]; ok {
//			return v
//		}
//		if root == nil {
//			hist[key][root] = 0
//			return 0
//		}
//		minOpts := 1 << 31
//		childTarget := target
//		if !same {
//			childTarget = 1 - target
//		}
//		if root.Val == target {
//			opts := dfs(root.Left, childTarget, true) + dfs(root.Right, childTarget, true)
//			minOpts = min(minOpts, opts)
//		} else {
//			// 操作1
//			opts := 1 + dfs(root.Left, childTarget, true) + dfs(root.Right, childTarget, true)
//			minOpts = min(minOpts, opts)
//			// 操作2
//			opts = 1 + dfs(root.Left, 1-childTarget, true) + dfs(root.Right, 1-childTarget, true)
//			minOpts = min(minOpts, opts)
//			// 操作3
//			opts = 1 + dfs(root.Left, 1-childTarget, false) + dfs(root.Right, 1-childTarget, false)
//			minOpts = min(minOpts, opts)
//		}
//		hist[key][root] = minOpts
//		return minOpts
//	}
//	return min(dfs(root, 0, true), dfs(root, 1, true)+1)
//}

func main() {
	root := &TreeNode{1,
		&TreeNode{1, nil, nil},
		&TreeNode{0, nil, &TreeNode{1, nil, nil}}}
	fmt.Println(closeLampInTree(root))
}
