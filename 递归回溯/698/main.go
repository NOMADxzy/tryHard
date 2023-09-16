package main

import "fmt"

func find(nums []int, acc int, k int, marked int, choosed int, cache map[int]int) bool {
	if choosed == len(nums) {
		cache[marked] = 2
		return true
	}
	if cache[marked] == 1 {
		return false
	} else if cache[marked] == 2 {
		return true
	}
	mask := 1
	for i := 0; i < len(nums); i++ {
		if marked&mask == 0 {
			if acc+nums[i] == k {
				if find(nums, 0, k, marked+mask, choosed+1, cache) {
					cache[marked] = 2
					return true
				}
			} else if acc+nums[i] < k {
				if find(nums, acc+nums[i], k, marked+mask, choosed+1, cache) {
					cache[marked] = 2
					return true
				}
			}
		}
		mask *= 2
	}
	cache[marked] = 1
	return false
}

func canPartitionKSubsets(nums []int, k int) bool {
	accTotal := 0
	for i := 0; i < len(nums); i++ {
		accTotal += nums[i]
	}
	if accTotal%k > 0 {
		return false
	} else {
		target := accTotal / k
		cache := map[int]int{}
		return find(nums, 0, target, 0, 0, cache)
	}
}

func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}
