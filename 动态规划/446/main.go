package main

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	m := len(nums)
	ans := 0
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}
	for k := 2; k < m; k++ {
		for j := 1; j < k; j++ {
			d := nums[k] - nums[j]
			sum := 0
			for i := j - 1; i >= 0; i-- {
				if nums[j]-nums[i] == d {
					sum += 1 + dp[i][j]
				}
			}
			dp[j][k] = sum
			ans += sum
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
}
