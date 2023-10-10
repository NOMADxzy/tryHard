package main

import "fmt"

func subarraysDivByK(nums []int, k int) int {
	m := len(nums)
	sums := make([]int, m+1)
	sums[1] = nums[0]
	for i := 2; i <= m; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	ans := 0
	dp := make([]int, m+1)
	for i := 1; i <= m; i++ {
		cnt := 0
		for j := i - 1; j >= 0; j-- {
			right := sums[i] - sums[j]
			if right%k == 0 {
				cnt += dp[j] + 1
				break
			}
		}
		dp[i] = cnt
		ans += dp[i]
	}
	return ans
}

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}
