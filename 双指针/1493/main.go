package main

import "fmt"

func longestSubarray(nums []int) int {
	l, r := 0, 0
	max := 0
	last0 := -1

	for ; r < len(nums); r++ {
		if nums[r] == 0 {
			l = last0 + 1
			last0 = r
		}
		if r-l+1 > max {
			max = r - l + 1
		}
	}
	return max - 1
}

func main() {
	fmt.Println(longestSubarray([]int{0, 0}))
}
