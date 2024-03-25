package main

import "fmt"

func minimumOperations(leaves string) int {
	m := len(leaves)
	dp := make([]int, 3)
	tmp := 0
	if leaves[0] == 'y' {
		tmp++
	}
	dp[0], dp[1], dp[2] = tmp, tmp, tmp
	if leaves[1] == 'r' {
		dp[1]++
		dp[2]++
	} else {
		dp[0]++
	}
	for i := 2; i < m; i++ {
		newDp := make([]int, 3)
		toR, toY := 0, 0
		if leaves[i] == 'y' {
			toR++
		} else {
			toY++
		}

		newDp[0] = toR + dp[0]
		newDp[1] = toY + min(dp[0], dp[1])
		newDp[2] = toR + min(dp[1], dp[2])
		dp = newDp
	}
	return dp[2]
}

func main() {
	fmt.Println(minimumOperations("rrryyyrryyyrr"))
}
