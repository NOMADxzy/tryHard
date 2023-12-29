package main

import "fmt"

// 动态规划 超时
func shipWithinDays1(weights []int, days int) int {
	n := len(weights)
	dp := make([]int, n)
	dp[0] = weights[0]
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + weights[i-1]
	}
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + weights[i]
	}
	for d := 1; d < days; d++ {
		newDp := make([]int, n)
		for i := d; i < n; i++ {
			//lastSum := weights[i]
			//best := max(lastSum, dp[i-1])
			left, right := d, i
			for left < right {
				mid := (left + right) / 2
				if sums[i+1]-sums[mid] < dp[mid-1] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			left--
			newDp[i] = max(sums[i+1]-sums[right], dp[right-1])
			if left >= d && max(sums[i+1]-sums[left], dp[left-1]) < newDp[i] {
				newDp[i] = max(sums[i+1]-sums[left], dp[left-1])
			}
		}
		dp = newDp
	}
	return dp[n-1]
}

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	n := len(weights)
	for i := 0; i < n; i++ {
		if weights[i] > left {
			left = weights[i]
		}
		right += weights[i]
	}
	check := func(w int) bool {
		cost := 0
		sum := 0
		for i := 0; i < n; i++ {
			if sum+weights[i] <= w {
				sum += weights[i]
			} else {
				cost++
				sum = 0
				i--
			}
		}
		return cost+1 <= days
	}
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
