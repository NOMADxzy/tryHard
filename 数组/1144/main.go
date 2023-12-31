package main

import "fmt"

func movesToMakeZigzag(nums []int) int {
	ans1, ans2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			dec := 0
			if i-1 >= 0 && nums[i-1] <= nums[i] {
				dec = nums[i] - nums[i-1] + 1
			}
			if i+1 < len(nums) && nums[i+1] <= nums[i] {
				dec = max(dec, nums[i]-nums[i+1]+1)
			}
			ans1 += dec
		} else {
			dec := 0
			if i-1 >= 0 && nums[i-1] <= nums[i] {
				dec = nums[i] - nums[i-1] + 1
			}
			if i+1 < len(nums) && nums[i+1] <= nums[i] {
				dec = max(dec, nums[i]-nums[i+1]+1)
			}
			ans2 += dec
		}
	}
	return min(ans2, ans1)
}

func main() {
	fmt.Println(movesToMakeZigzag([]int{10, 4, 4, 10, 10, 6, 2, 3}))
}
