package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(father map[*TreeNode]*TreeNode, root *TreeNode, start int, target **TreeNode) {
	if root.Val == start {
		*target = root
	}
	if root.Left != nil {
		father[root.Left] = root
		findTarget(father, root.Left, start, target)
	}
	if root.Right != nil {
		father[root.Right] = root
		findTarget(father, root.Right, start, target)
	}
}

func nextDoneTime(root *TreeNode, father map[*TreeNode]*TreeNode, visited map[*TreeNode]bool) int {
	if root == nil || visited[root] {
		return 0
	}
	visited[root] = true
	doneTime1 := nextDoneTime(root.Left, father, visited)
	doneTime2 := nextDoneTime(root.Right, father, visited)
	doneTime3 := nextDoneTime(father[root], father, visited)
	return max(doneTime1, max(doneTime2, doneTime3)) + 1
}

func amountOfTime(root *TreeNode, start int) int {
	var startNode *TreeNode
	father := map[*TreeNode]*TreeNode{}
	findTarget(father, root, start, &startNode)
	return nextDoneTime(startNode, father, map[*TreeNode]bool{}) - 1
}

func main() {
	root := &TreeNode{1,
		&TreeNode{5, nil,
			&TreeNode{4,
				&TreeNode{9, nil, nil},
				&TreeNode{2, nil, nil}}},
		&TreeNode{3,
			&TreeNode{10, nil, nil},
			&TreeNode{6, nil, nil}}}
	fmt.Println(amountOfTime(root, 3))
}
