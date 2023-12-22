package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type MyTreeNode struct {
	Left  *MyTreeNode
	Right *MyTreeNode
	Val   int
}

func buildTree(root *MyTreeNode, val int, layer int) {
	if layer == 0 {
		return
	} else {
		if val < layer {
			if root.Left == nil {
				root.Left = &MyTreeNode{nil, nil, 0}
			}
			buildTree(root.Left, val, layer/2)
		} else {
			if root.Right == nil {
				root.Right = &MyTreeNode{nil, nil, 1}
			}
			buildTree(root.Right, val-layer, layer/2)
		}
	}
}

func findMaximumXOR(nums []int) int {

	maxNum := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	mask := 1
	for i := maxNum; i > 1; i = i >> 1 {
		mask *= 2
	}

	root := &MyTreeNode{nil, nil, 0}
	for i := 0; i < len(nums); i++ {
		buildTree(root, nums[i], mask)
	}
	maxVal := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		k := nums[i]
		r := root
		for m := mask; m > 0; m /= 2 {
			if k&m > 0 { //剪枝
				if r.Left != nil {
					sum += m
					r = r.Left
				} else {
					r = r.Right
				}
			} else {
				if r.Right != nil {
					sum += m
					r = r.Right
				} else {
					r = r.Left
				}
			}
		}
		if sum > maxVal {
			maxVal = sum
		}
	}
	return maxVal
}

func main() {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
}
