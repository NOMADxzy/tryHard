package main

import "fmt"

func waysToMakeFair(nums []int) int {
	m := len(nums)
	if m == 1 {
		return 1
	}
	dp0 := make([]int, m)
	dp1 := make([]int, m)
	dp0[0], dp0[1] = nums[0], nums[0]
	dp1[1] = nums[1]
	for i := 2; i < m; i++ {
		if i%2 == 0 {
			dp0[i] = dp0[i-1] + nums[i]
			dp1[i] = dp1[i-1]
		} else {
			dp1[i] = dp1[i-1] + nums[i]
			dp0[i] = dp0[i-1]
		}
	}
	total := dp0[m-1] + dp1[m-1]
	ans := 0

	for i := 0; i < m; i++ {
		target := total - nums[i]
		if target%2 == 0 {
			target /= 2

			sum := 0
			if i-1 >= 0 {
				sum += dp0[i-1]
			}
			sum += dp1[m-1] - dp1[i]
			if sum == target {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(waysToMakeFair([]int{1, 2, 3}))
}
