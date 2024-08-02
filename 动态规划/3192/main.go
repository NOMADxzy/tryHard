package main

import "fmt"

func minOperations(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if ans%2 == 1 {
				ans++
			}
		} else {
			if ans%2 == 0 {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{1, 0, 1, 0, 1}))
}
