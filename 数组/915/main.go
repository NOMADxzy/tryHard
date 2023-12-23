package main

import "fmt"

func partitionDisjoint(nums []int) int {
	leftMaxVal, lb := nums[0], 0
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < leftMaxVal {
			lb = i
			leftMaxVal = maxVal
		}
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	return lb + 1
}

func main() {
	fmt.Println(partitionDisjoint([]int{5, 0, 3, 8, 6}))
}
