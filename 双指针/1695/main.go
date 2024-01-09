package main

import "fmt"

func maximumUniqueSubarray(nums []int) int {
	lastIdx := map[int]int{}
	l, r := 0, 0
	ans := 0
	sum := 0
	for r < len(nums) {
		sum += nums[r]
		if idx, ok := lastIdx[nums[r]]; ok {
			for l <= idx {
				sum -= nums[l]
				l++
			}
		}
		lastIdx[nums[r]] = r
		ans = max(ans, sum)
		r++
	}
	return ans
}

func main() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))
}
