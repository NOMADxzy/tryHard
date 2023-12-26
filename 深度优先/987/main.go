package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func VerticalOrder(root *TreeNode, r, c int, mat *[][][]int) {
	if root == nil {
		return
	}
	(*mat)[c][r] = append((*mat)[c][r], root.Val)
	VerticalOrder(root.Left, r+1, c-1, mat)
	VerticalOrder(root.Right, r+1, c+1, mat)
}

func verticalTraversal(root *TreeNode) [][]int {
	var findBorder func(root *TreeNode, r, c int) (int, int, int)
	findBorder = func(root *TreeNode, r, c int) (int, int, int) {
		if root == nil {
			return 0, 0, 0
		}
		ll, lr, ldep := findBorder(root.Left, r-1, c-1)
		rl, rr, rdep := findBorder(root.Right, r-1, c+1)
		return min(c, min(ll, rl)), max(c, max(lr, rr)), max(ldep, rdep) + 1
	}
	left, right, dep := findBorder(root, 0, 0)
	mat := make([][][]int, right-left+1)
	ans := make([][]int, right-left+1)
	for i := 0; i < len(mat); i++ {
		mat[i] = make([][]int, dep)
	}
	VerticalOrder(root, 0, -left, &mat)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < dep; j++ {
			if len(mat[i][j]) > 1 {
				sort.Slice(mat[i][j], func(l, r int) bool {
					return mat[i][j][l] < mat[i][j][r]
				})
			}
			for k := 0; k < len(mat[i][j]); k++ {
				ans[i] = append(ans[i], mat[i][j][k])
			}
		}
	}
	return ans
}

func main() {
	root := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(verticalTraversal(root))
}
