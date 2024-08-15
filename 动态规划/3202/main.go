package main

import "fmt"

func maximumLength(nums []int, k int) int {
	m := len(nums)
	dp := make([]map[int]int, m)
	ans := 0
	for i := 0; i < m; i++ {
		dp[i] = map[int]int{}
	}
	for j := 1; j < m; j++ {
		//dp[j][nums[0]+nums[1]%m] = 1
		for i := j - 1; i >= 0; i-- {
			v := (nums[i] + nums[j]) % k
			dp[j][v] = max(dp[j][v], dp[i][v]+1)
			ans = max(ans, dp[j][v]+1)
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 2, 3, 4, 5}, 2))
}
