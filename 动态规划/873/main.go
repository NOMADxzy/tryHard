package main

import "fmt"

func lenLongestFibSubseq(arr []int) int {
	m := len(arr)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][1] = 2
	ans := 2

	for k := 2; k < m; k++ {
		for r := k - 1; r > 0; r-- {
			length := 2
			target := arr[k] - arr[r]
			left, right := 0, r-1
			if arr[left] > target || arr[right] < target {
			} else if arr[left] == target {
				length = dp[left][r] + 1
			} else {
				for left < right {
					mid := (left + right) / 2
					if arr[mid] < target {
						left = mid + 1
					} else {
						right = mid
					}
				}
				if arr[right] == target {
					length = dp[right][r] + 1
				}
			}
			dp[r][k] = length
			if length > ans {
				ans = length
			}
		}
		dp[0][k] = 2
	}
	if ans == 2 {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(lenLongestFibSubseq([]int{2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50}))
}
