package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	var min int
	first := true
	sum := nums[0]
	for l < len(nums) {
		if sum >= target {
			if first {
				min = r - l + 1
				first = false
			} else if r-l+1 < min {
				min = r - l + 1
			}
			sum -= nums[l]
			l++
		} else {
			if r < len(nums)-1 {
				r++
				sum += nums[r]
			} else {
				break
			}
		}
	}
	return min
}

func main() {
	fmt.Println(minSubArrayLen(4, []int{5}))
}
