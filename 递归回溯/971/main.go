package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findChildCnt(root *TreeNode, childCnt map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	childCnt[root] = 1 + findChildCnt(root.Left, childCnt) + findChildCnt(root.Right, childCnt)
	return childCnt[root]
}

func checkNext(ans *[]int, root *TreeNode, voyage []int, childCnt map[*TreeNode]int) bool {
	if root.Val != voyage[0] {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	} else if root.Left == nil {
		return checkNext(ans, root.Right, voyage[1:], childCnt)
	} else if root.Right == nil {
		return checkNext(ans, root.Left, voyage[1:], childCnt)
	} else {
		size1, size2 := childCnt[root.Left], childCnt[root.Right]
		if voyage[1] == root.Left.Val {
			return checkNext(ans, root.Left, voyage[1:size1+1], childCnt) && checkNext(ans, root.Right, voyage[1+size1:size2+size1+1], childCnt)
		} else if voyage[1] == root.Right.Val && voyage[1+size2] == root.Left.Val {
			*ans = append(*ans, root.Val)
			return checkNext(ans, root.Left, voyage[1+size2:size2+size1+1], childCnt) && checkNext(ans, root.Right, voyage[1:size2+1], childCnt)
		} else {
			return false
		}
	}
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	childCnt := map[*TreeNode]int{}
	findChildCnt(root, childCnt)
	var ans []int
	b := checkNext(&ans, root, voyage, childCnt)
	if !b {
		return []int{-1}
	}
	return ans
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(flipMatchVoyage(root, []int{1, 3, 2}))
}
