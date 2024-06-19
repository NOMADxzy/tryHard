package main

import "fmt"

func maxSubarrays(nums []int) int {
	target := nums[0]
	for i := 1; i < len(nums); i++ {
		target &= nums[i]
	}
	if target > 0 {
		return 1
	}
	acc := nums[0]
	ans := 0
	for i := 0; i < len(nums); i++ {
		acc &= nums[i]
		if acc == 0 {
			ans++
			if i+1 < len(nums) {
				acc = nums[i+1]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxSubarrays([]int{1, 0, 2, 0, 1, 2}))
}
