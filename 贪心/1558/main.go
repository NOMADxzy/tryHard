package main

import (
	"fmt"
	"sort"
)

func minOperations(nums []int) int {
	maxDiv := 0
	delCnt := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	hist := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		div, del := 0, 0
		for j := nums[i]; j > 0; {
			if hist[j] != nil {
				div += hist[j][0]
				del += hist[j][1]
				break
			}
			if j%2 == 0 {
				j /= 2
				div++
			} else {
				j -= 1
				del++
			}
		}
		hist[nums[i]] = []int{div, del}
		maxDiv = max(maxDiv, div)
		delCnt += del
	}
	return maxDiv + delCnt
}

func main() {
	fmt.Println(minOperations([]int{2, 4, 8, 16}))

}
