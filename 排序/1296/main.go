package main

import (
	"fmt"
	"sort"
)

func isPossibleDivide(nums []int, k int) bool {
	nMap := map[int]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		nMap[nums[i]]++
	}

	for i := 0; i < len(nums); i++ {
		if nMap[nums[i]] > 0 {
			for j := 0; j < k; j++ {
				nMap[nums[i]+j]--
				if nMap[nums[i]+j] < 0 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isPossibleDivide([]int{1, 2, 3, 3, 4, 4, 5, 6}, 4))
}
