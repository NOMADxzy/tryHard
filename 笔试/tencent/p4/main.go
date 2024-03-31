package main

import "fmt"

func solve(arr []int, k int) int {
	n := len(arr)
	accs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		accs[i] = accs[i-1] ^ arr[i-1]
	}
	dp := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, n)
	}
	dp[0] = accs[1:]

	// i=1,j=1
	for i := 1; i < k; i++ { // i+1段
		for j := i; j < n; j++ {
			maxVal := -1 << 31
			for t := i; t <= j; t++ {
				v1 := dp[i-1][t-1]
				v2 := accs[j+1] ^ accs[t]
				tmp := v1 + v2
				if tmp > maxVal {
					maxVal = tmp
				}
			}
			dp[i][j] = maxVal
		}
	}
	return dp[k-1][n-1]
}

func main() {
	var n, k int
	_, _ = fmt.Scan(&n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
	}
	fmt.Println(solve(arr, k))
}
