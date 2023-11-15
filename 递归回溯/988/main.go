package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func violence(best *string, pre string, root *TreeNode) {
	cur := string(rune('a'+root.Val)) + pre
	if root.Left == nil && root.Right == nil {
		bestSmaller := false
		var p int
		for p = 0; p < len(cur) && p < len(*best); p++ {
			if cur[p] > (*best)[p] {
				bestSmaller = true
				break
			} else if cur[p] < (*best)[p] {
				break
			}
		}
		if bestSmaller || p == len(*best) {
		} else {
			*best = cur
		}
	} else {
		if root.Left != nil {
			violence(best, cur, root.Left)
		}
		if root.Right != nil {
			violence(best, cur, root.Right)
		}
	}
}

func smallestFromLeaf(root *TreeNode) string {
	best := "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"
	violence(&best, "", root)
	return best
}

func main() {
	root := &TreeNode{0,
		&TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}}
	fmt.Println(smallestFromLeaf(root))
}
