package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func part(nums []int, left int, right int) *TreeNode {
	mid := (left + right) / 2
	if left == right {
		return &TreeNode{nums[mid], nil, nil}
	} else if left == mid {
		return &TreeNode{nums[mid], nil, part(nums, mid+1, right)}
	} else {
		return &TreeNode{nums[mid], part(nums, left, mid-1), part(nums, mid+1, right)}
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	return part(nums, 0, len(nums)-1)
}

func main() {
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(root.Val)
}
