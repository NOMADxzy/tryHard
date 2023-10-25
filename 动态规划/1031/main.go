package main

import "fmt"

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	m := len(nums)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 2)
	}

	sums := make([]int, m)
	sums[0] = nums[0]
	for i := 1; i < m; i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	if secondLen < firstLen {
		firstLen, secondLen = secondLen, firstLen
	}

	for i := m - firstLen; i >= 0; i-- {
		rightSum := sums[i-1+firstLen]
		if i > 0 {
			rightSum -= sums[i-1]
		}
		if i >= m-1 || rightSum > dp[i+1][0] {
			dp[i][0] = rightSum
		} else {
			dp[i][0] = dp[i+1][0]
		}
	}
	for i := m - secondLen; i >= 0; i-- {
		rightSum := sums[i-1+secondLen]
		if i > 0 {
			rightSum -= sums[i-1]
		}
		if i >= m-1 || rightSum > dp[i+1][1] {
			dp[i][1] = rightSum
		} else {
			dp[i][1] = dp[i+1][1]
		}
	}

	ans := 0
	for i := 0; i <= m-firstLen-secondLen; i++ {
		val1 := sums[i-1+firstLen]
		if i > 0 {
			val1 -= sums[i-1]
		}
		l1 := val1 + dp[i+firstLen][1]

		val2 := sums[i-1+secondLen]
		if i > 0 {
			val2 -= sums[i-1]
		}
		l2 := val2 + dp[i+secondLen][0]

		if l1 > ans {
			ans = l1
		}
		if l2 > ans {
			ans = l2
		}
	}
	return ans
}

func main() {
	fmt.Println(maxSumTwoNoOverlap([]int{1, 16, 20, 13, 2, 11, 12, 18, 8, 5, 18, 15, 0, 7}, 3, 5))
}
