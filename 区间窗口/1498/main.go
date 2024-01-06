package main

import (
	"fmt"
	"sort"
)

func numSubseq(nums []int, target int) int {
	MOD := 1000000007
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := 0
	addVal := 1
	addValMap := map[int]int{}
	for j := 0; j < len(nums)-1; j++ {
		addVal = (addVal + addVal) % MOD
		addValMap[j+1] = addVal
	}
	for i := 0; i < len(nums) && nums[i] <= target-nums[i]; i++ {
		left, right := i, len(nums)-1
		limit := target - nums[i]
		var n int
		if nums[right] <= limit {
			n = len(nums) - i - 1
		} else {
			for left < right {
				mid := (left + right) / 2
				if nums[mid] <= limit {
					left = mid + 1
				} else {
					right = mid
				}
			}
			n = left - i - 1
		}
		addVal = addValMap[n]
		ans = (ans + addVal) % MOD
	}
	return ans
}

func main() {
	fmt.Println(numSubseq([]int{2, 3, 3, 4, 6, 7}, 12))
}
