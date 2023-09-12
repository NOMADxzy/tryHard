package main

import "fmt"

func pivotIndex(nums []int) int {
	sum := 0
	idx := -1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	leftSum := 0

	for i := 0; i < len(nums); i++ {
		other := sum - nums[i]
		if other%2 == 0 {
			if leftSum == other/2 {
				idx = i
				break
			}
		}
		leftSum += nums[i]

	}
	return idx
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
