package main

import "fmt"

func maxOperations(nums []int) int {
	m := len(nums)
	dp := map[int]int{}
	var dfs func(start, end int, target int) int
	dfs = func(start, end int, target int) int {
		key := (start*(m+1)+end)*(2*m+1) + target
		if val, ok := dp[key]; ok {
			return val
		}
		if end-start+1 <= 1 {
			return 0
		}
		maxVal := 0
		if nums[start]+nums[end] == target {
			maxVal = 1 + dfs(start+1, end-1, target)
		}
		if nums[start]+nums[start+1] == target {
			maxVal = max(maxVal, 1+dfs(start+2, end, target))
		}
		if nums[end-1]+nums[end] == target {
			maxVal = max(maxVal, 1+dfs(start, end-2, target))
		}
		dp[key] = maxVal
		return maxVal
	}
	ans := max(dfs(0, m-1, nums[0]+nums[m-1]), max(dfs(0, m-1, nums[0]+nums[1]), dfs(0, m-1, nums[m-2]+nums[m-1])))
	return ans
}

func main() {
	fmt.Println(maxOperations([]int{3, 2, 1, 2, 3, 4}))
}
