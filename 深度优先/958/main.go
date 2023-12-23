package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	var dfs func(root *TreeNode) (int, bool, bool)
	dfs = func(root *TreeNode) (int, bool, bool) {
		if root == nil {
			return 0, true, true
		}
		ldeep, lperfect, lok := dfs(root.Left)
		rdeep, rperfect, rok := dfs(root.Right)
		if !lok || !rok {
			return 0, false, false
		}
		if ldeep == rdeep+1 && rperfect {
			return ldeep + 1, false, true
		} else if ldeep == rdeep && lperfect {
			return ldeep + 1, rperfect, true
		} else {
			return 0, false, false
		}
	}
	_, _, ok := dfs(root)
	return ok
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, nil}}
	fmt.Println(isCompleteTree(root))
}
