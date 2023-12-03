package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goNextLayer(root *TreeNode, curLayer int, layerSum *[]int, childSum map[*TreeNode]int, parents map[*TreeNode]*TreeNode) {
	if len(*layerSum) <= curLayer {
		*layerSum = append(*layerSum, 0)
	}
	(*layerSum)[curLayer] += root.Val
	if root.Left != nil {
		parents[root.Left] = root
		childSum[root] += root.Left.Val
		goNextLayer(root.Left, curLayer+1, layerSum, childSum, parents)
	}
	if root.Right != nil {
		parents[root.Right] = root
		childSum[root] += root.Right.Val
		goNextLayer(root.Right, curLayer+1, layerSum, childSum, parents)
	}
}

func getSumNext(root *TreeNode, curLayer int, layerSum []int, parents map[*TreeNode]*TreeNode, childSum map[*TreeNode]int) {
	if root == nil {
		return
	}
	sum := layerSum[curLayer] - childSum[parents[root]]
	root.Val = sum
	getSumNext(root.Left, curLayer+1, layerSum, parents, childSum)
	getSumNext(root.Right, curLayer+1, layerSum, parents, childSum)
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	parents := map[*TreeNode]*TreeNode{root: &TreeNode{0, nil, nil}}
	childSum := map[*TreeNode]int{parents[root]: root.Val}
	var layerSum []int
	goNextLayer(root, 0, &layerSum, childSum, parents)
	getSumNext(root, 0, layerSum, parents, childSum)
	return root
}

func main() {
	root := &TreeNode{5,
		&TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{10, nil, nil}},
		&TreeNode{9, nil, &TreeNode{7, nil, nil}}}
	fmt.Println(replaceValueInTree(root).Val)
	fmt.Println("ok")
}
