package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	modMap := map[int]bool{}
	preRes := nums[0] % k
	prepreRes := 0
	for i := 1; i < len(nums); i++ {
		t := sums[i] % k
		modMap[prepreRes] = true
		if modMap[t] {
			return true
		} else {
			prepreRes = preRes
			preRes = t
		}
	}
	return false
}

//动态规划，内存溢出
func checkSubarraySum1(nums []int, k int) bool {
	dp := make([][]bool, 2)
	dp[0] = make([]bool, k)
	//dp[0][nums[0]%k] = true

	for i := 1; i < len(nums); i++ {
		dp[1] = make([]bool, k)
		dp[1][(nums[i-1]+nums[i])%k] = true
		for j := 0; j < k; j++ {
			if dp[0][j] {
				dp[1][(j+nums[i])%k] = true
			}
		}
		if dp[1][0] {
			return true
		}
		copy(dp[0], dp[1])
	}
	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{1, 2, 12}, 6))
}
