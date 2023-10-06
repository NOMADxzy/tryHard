package main

import "fmt"

func findMaxLength(nums []int) int {
	newNums := make([]int, len(nums))
	newNums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		newNums[i] = newNums[i-1] + nums[i]
	}
	for i := 0; i < len(nums); i++ {
		newNums[i] = 2*newNums[i] - i
	}
	vMap := map[int]int{}
	vMap[1] = -1
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		val := vMap[newNums[i]]
		if val != 0 {
			if val == -1 {
				val++
			}
			if (i - val + 1) > maxLen {
				maxLen = i - val + 1
			}
		} else {
			vMap[newNums[i]] = i + 1
		}
	}
	return maxLen
}

func main() {
	fmt.Println(findMaxLength([]int{0, 0, 1, 0, 0, 0, 1, 1}))
}
