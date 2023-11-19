package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func seqNext(root *TreeNode, sMap *map[string]int, ans *[]*TreeNode) string {
	if root == nil {
		return ""
	}
	s := fmt.Sprintf("%d(%s)(%s)", root.Val, seqNext(root.Left, sMap, ans), seqNext(root.Right, sMap, ans))
	(*sMap)[s]++
	if (*sMap)[s] == 2 {
		*ans = append(*ans, root)
	}
	return s
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var ans []*TreeNode
	sMap := map[string]int{}
	seqNext(root, &sMap, &ans)
	return ans
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2, &TreeNode{4, nil, nil}, nil},
		&TreeNode{3, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{4, nil, nil}}}
	fmt.Println(findDuplicateSubtrees(root))
}
