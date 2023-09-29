package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(nums, nextGreater []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	maxIdx := l
	for nextGreater[maxIdx] <= r && nextGreater[maxIdx] != -1 {
		maxIdx = nextGreater[maxIdx]
	}
	return &TreeNode{nums[maxIdx],
		build(nums, nextGreater, l, maxIdx-1), build(nums, nextGreater, maxIdx+1, r)}
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var stack, nextGreater []int
	nextGreater = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			idx := stack[len(stack)-1]
			nextGreater[idx] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		idx := stack[len(stack)-1]
		nextGreater[idx] = -1
		stack = stack[:len(stack)-1]
	}
	root := build(nums, nextGreater, 0, len(nums)-1)
	return root
}

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
}
