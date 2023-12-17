package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func cntNext(root *TreeNode, vMap map[int]int) int {
	val := root.Val
	if root.Left != nil {
		val += cntNext(root.Left, vMap)
	}
	if root.Right != nil {
		val += cntNext(root.Right, vMap)
	}
	vMap[val]++
	return val
}

func findFrequentTreeSum(root *TreeNode) []int {
	vMap := map[int]int{}
	cntNext(root, vMap)
	var ans []int
	maxCnt := 0
	for _, v := range vMap {
		if v > maxCnt {
			maxCnt = v
		}
	}
	for k, v := range vMap {
		if v == maxCnt {
			ans = append(ans, k)
		}
	}
	return ans
}

func main() {
	root := &TreeNode{5, &TreeNode{2, nil, nil}, &TreeNode{-3, nil, nil}}
	fmt.Println(findFrequentTreeSum(root))
}
