package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goToNextLayer(root *TreeNode, pos, layer int, idxsL, idxsR *[]int) {
	if root == nil {
		return
	}
	if len(*idxsL) <= layer {
		*idxsL = append(*idxsL, pos)
	} else if pos < (*idxsL)[layer] {
		(*idxsL)[layer] = pos
	}
	if len(*idxsR) <= layer {
		*idxsR = append(*idxsR, pos)
	} else if pos > (*idxsR)[layer] {
		(*idxsR)[layer] = pos
	}
	goToNextLayer(root.Left, pos*2, layer+1, idxsL, idxsR)
	goToNextLayer(root.Right, pos*2+1, layer+1, idxsL, idxsR)
}

func widthOfBinaryTree(root *TreeNode) int {
	var idxsL, idxsR []int
	goToNextLayer(root, 0, 0, &idxsL, &idxsR)
	ans := 0
	for l := 0; l < len(idxsL); l++ {
		if d := idxsR[l] - idxsL[l] + 1; d > ans {
			ans = d
		}
	}
	return ans
}

func main() {
	root := &TreeNode{1,
		&TreeNode{3, &TreeNode{5, nil, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{2, nil, &TreeNode{9, nil, nil}}}
	fmt.Println(widthOfBinaryTree(root))
}
