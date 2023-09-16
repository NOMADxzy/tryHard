package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{{0, nil, nil}}
	}
	var trees []*TreeNode
	for i := 0; i < n/2; i++ {
		leftNodeNums := 2*i + 1
		leftTrees := build(leftNodeNums)
		rightTrees := build(n - 1 - leftNodeNums)
		for j := 0; j < len(leftTrees); j++ {
			for k := 0; k < len(rightTrees); k++ {
				trees = append(trees, &TreeNode{0, leftTrees[j], rightTrees[k]})
			}
		}
	}
	return trees
}

func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}
	return build(n)
}

func main() {
	trees := allPossibleFBT(1)
	fmt.Println(len(trees))
}
