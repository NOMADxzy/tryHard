package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(root *TreeNode, dp []int, target int, cnt *int) {
	dp1 := []int{root.Val}
	if root.Val == target {
		*cnt++
	}
	//dp = append(dp, 0)
	for i := 0; i < len(dp); i++ {
		//dp[i] += root.Val
		dp1 = append(dp1, dp[i]+root.Val)
		if dp1[i+1] == target {
			*cnt++
		}
	}
	if root.Left != nil {
		check(root.Left, dp1, target, cnt)
	}
	if root.Right != nil {
		check(root.Right, dp1, target, cnt)
	}

}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	cnt := 0
	var dp []int
	check(root, dp, targetSum, &cnt)
	return cnt
}

func main() {
	root := &TreeNode{5,
		&TreeNode{4,
			&TreeNode{11,
				&TreeNode{7, nil, nil},
				&TreeNode{2, nil, nil}},
			nil},
		&TreeNode{8,
			&TreeNode{13, nil, nil},
			&TreeNode{4, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}}}}
	fmt.Println(pathSum(root, 22))
}
