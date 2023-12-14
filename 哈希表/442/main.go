package main

import "fmt"

func findDuplicates(nums []int) []int {
	var ans []int
	for i := 0; i < len(nums); i++ {
		rightPlace := nums[i] - 1
		if rightPlace != i {
			if nums[rightPlace] == nums[i] {
				ans = append(ans, nums[i])
				nums[i] = -1
			} else {
				if nums[rightPlace] < 0 {
					nums[rightPlace] = nums[i]
					nums[i] = -1
				} else {
					nums[rightPlace], nums[i] = nums[i], nums[rightPlace]
					i--
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
