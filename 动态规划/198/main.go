package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	nums[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		choice1 := nums[i] + nums[i-2]
		choice2 := nums[i-1]
		nums[i] = max(choice1, choice2)
	}
	return nums[len(nums)-1]
}

func main() {
	fmt.Println(rob([]int{2, 3}))
}
