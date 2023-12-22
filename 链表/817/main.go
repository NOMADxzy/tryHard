package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	ans := 0
	node := head
	nMap := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		nMap[nums[i]] = true
	}

	for node != nil {
		for node != nil && !nMap[node.Val] {
			node = node.Next
		}
		if node == nil {
			break
		}
		ans++
		for node != nil && nMap[node.Val] {
			node = node.Next
		}
	}
	return ans
}

func main() {
	head := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}
	fmt.Println(numComponents(head, []int{0, 1, 3}))
}
