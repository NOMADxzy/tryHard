package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var ans []*TreeNode
	var dfs func(root, father *TreeNode)
	isDeleted := map[int]bool{-1: true}
	for i := 0; i < len(to_delete); i++ {
		isDeleted[to_delete[i]] = true
	}

	unLink := func(father, child *TreeNode) {
		if father.Left == child {
			father.Left = nil
		} else {
			father.Right = nil
		}
	}

	dfs = func(root, father *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left, root)
		dfs(root.Right, root)
		if isDeleted[root.Val] {
			unLink(father, root)
		} else if isDeleted[father.Val] {
			ans = append(ans, root)
		}
	}

	dfs(root, &TreeNode{-1, root, nil})
	return ans
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}},
	}
	lists := delNodes(root, []int{3, 5})
	fmt.Println(len(lists))
}
