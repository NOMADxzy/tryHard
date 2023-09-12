package main

import (
	"fmt"
	"sort"
)

// 垃圾方法
func find(nums []int, pos int, target int) bool {
	if target == 0 {
		return true
	} else if pos == len(nums) {
		return false
	} else if nums[pos] <= target {
		for i := pos; i < len(nums); i++ {
			if find(nums, i+1, target-nums[i]) {
				return true
			}
		}
	}
	return false
}

func canPartition1(nums []int) bool {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	} else {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		target := sum / 2
		return find(nums, 0, target)

	}
}

func canPartition(nums []int) bool {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	} else {
		target := sum / 2

		dp := make([][]bool, len(nums))
		for i := 0; i < len(nums); i++ {
			dp[i] = make([]bool, target+1)
			dp[i][0] = true
		}
		if nums[0] == target {
			return true
		} else if nums[0] < target {
			dp[0][nums[0]] = true
		}

		for i := 1; i < len(nums); i++ {
			for j := 0; j < target+1; j++ {
				if dp[i-1][j] {
					dp[i][j] = true
					if j+nums[i] <= target {
						dp[i][j+nums[i]] = true
					}
				}
				if dp[i][target] {
					return true
				}
			}
		}
		return false
	}
}

func main() {
	fmt.Println(canPartition([]int{14, 9, 8, 4, 3, 2}))
}
