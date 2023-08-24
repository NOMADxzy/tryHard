package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generate(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var Trees []*TreeNode
	for r := start; r <= end; r++ {
		leftTrees := generate(start, r-1)
		rightTrees := generate(r+1, end)

		for i := 0; i < len(leftTrees); i++ {
			for j := 0; j < len(rightTrees); j++ {
				Trees = append(Trees, &TreeNode{r, leftTrees[i], rightTrees[j]})
			}
		}
	}
	return Trees
}

func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}

func main() {
	a := generateTrees(8)
	fmt.Println(len(a))
}
