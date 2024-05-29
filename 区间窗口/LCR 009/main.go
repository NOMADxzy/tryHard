package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	l, r := 0, 0
	acc := 1
	for r < len(nums) {
		acc *= nums[r]
		for acc >= k && l <= r {
			acc /= nums[l]
			l++
		}
		ans += r - l + 1
		r++
	}
	return ans
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
