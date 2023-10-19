package main

import "fmt"

func longestSubsequence(arr []int, difference int) int {
	m := len(arr)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = 1

	idxMap := map[int]int{}
	idxMap[arr[0]] = 0

	for i := 1; i < m; i++ {
		use, abandon := 0, dp[i-1][1]
		if dp[i-1][0] > dp[i-1][1] {
			abandon = dp[i-1][0]
		}
		if idx, ok := idxMap[arr[i]-difference]; ok {
			use = dp[idx][1] + 1
		} else {
			use = 1
		}
		dp[i][0], dp[i][1] = abandon, use
		idxMap[arr[i]] = i
	}

	ans := dp[m-1][0]
	if dp[m-1][1] > ans {
		ans = dp[m-1][1]
	}
	return ans
}

func main() {
	fmt.Println(longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}
