package main

import "fmt"

func longestSubarray(nums []int) int {
	preMaxLen := 0
	curMax := nums[0]
	acc := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < curMax {
			preMaxLen = max(preMaxLen, acc)
			acc = 0
			continue
		} else if nums[i] == curMax {
			acc++
		} else {
			curMax = nums[i]
			acc = 1
			preMaxLen = 0
		}
	}
	return max(preMaxLen, acc)
}

func main() {
	fmt.Println(longestSubarray([]int{1, 2, 3, 3, 1}))
}
