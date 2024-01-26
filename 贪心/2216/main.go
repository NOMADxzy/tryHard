package main

import "fmt"

func minDeletion(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); {
		j := i + 1
		for j < len(nums) && nums[j] == nums[i] {
			ans++
			j++
		}
		if j == len(nums) {
			ans++
			break
		}
		i = j + 1
	}
	return ans
}

func main() {
	fmt.Println(minDeletion([]int{1, 1, 2, 2, 3, 3}))
}
