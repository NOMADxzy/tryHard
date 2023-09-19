package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func check(nums []int, val int) int {
	lastChoosed := false
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if !lastChoosed && nums[i] <= val {
			lastChoosed = true
			cnt++
		} else {
			lastChoosed = false
		}
	}
	return cnt
}

func minCapability(nums []int, k int) int {
	minVal, maxVal := nums[0], nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] <= minVal {
			minVal = nums[i]
		}
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}

	for minVal < maxVal {
		mid := (maxVal + minVal) / 2
		if check(nums, mid) < k {
			minVal = mid + 1
		} else {
			maxVal = mid
		}
	}
	return minVal
}

// 动态规划 ，超时
func minCapability_(nums []int, k int) int {
	m := len(nums)
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, k+1)
	}

	dp[0][1] = nums[0]
	if m == 1 {
		return nums[0]
	}
	dp[1][1] = min(nums[0], nums[1])
	if m == 2 {
		return dp[1][1]
	}

	for i := 2; i < m; i++ {
		for j := 1; j <= i/2+1 && j <= k; j++ {
			if dp[1][j] > 0 {
				dp[2][j] = min(max(dp[0][j-1], nums[i]), dp[1][j])
			} else {
				dp[2][j] = max(dp[0][j-1], nums[i])
			}
		}
		copy(dp[0], dp[1])
		copy(dp[1], dp[2])
	}
	return dp[1][k]
}

func main() {
	fmt.Println(minCapability([]int{2, 7, 9, 3, 1}, 2))
}
