package main

import "fmt"

func minOperations(nums []int, k int) int {
	totalXorResult := 0
	for i := 0; i < len(nums); i++ {
		totalXorResult ^= nums[i]
	}
	mask := 1
	ans := 0
	for mask <= totalXorResult || mask <= k {
		if totalXorResult&mask != k&mask {
			ans++
		}
		mask *= 2
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{2, 1, 3, 4}, 1))
}
