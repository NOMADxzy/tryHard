package main

import "fmt"

func countAlternatingSubarrays(nums []int) int64 {
	ans := int64(1)
	pre := 1
	for i := 1; i < len(nums); i++ {
		cur := 1
		if nums[i] != nums[i-1] {
			cur += pre
		}
		ans += int64(cur)
		pre = cur
	}
	return ans
}

func main() {
	fmt.Println(countAlternatingSubarrays([]int{0, 1, 1, 1}))
}
