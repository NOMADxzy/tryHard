package main

import "fmt"

func minOperations(nums []int) int {
	ans := 0
	opts := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if opts[i]%2 == 1 {
			if nums[i] == 0 {
				continue
			} else {
				if i+2 >= len(nums) {
					return -1
				} else {
					opts[i]++
					opts[i+1]++
					opts[i+2]++
					ans++
				}
			}
		} else {
			if nums[i] == 1 {
				continue
			} else {
				if i+2 >= len(nums) {
					return -1
				} else {
					opts[i]++
					opts[i+1]++
					opts[i+2]++
					ans++
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{1, 0, 1}))
}
