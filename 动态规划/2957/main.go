package main

import (
	"fmt"
)

func removeAlmostEqualCharacters(word string) int {
	m := len(word)
	dp := []int{0, 1}
	isNeighbor := func(a, b byte) bool {
		x := int(a - 'a')
		y := int(b - 'a')
		if x-y <= 1 && x-y >= -1 {
			return true
		}
		return false
	}

	for i := 1; i < m; i++ {
		nextDp := []int{0, 0}
		if isNeighbor(word[i-1], word[i]) {
			nextDp[0] = dp[1]
			nextDp[1] = min(dp[0], dp[1]) + 1
		} else {
			nextDp[0] = min(dp[0], dp[1])
			nextDp[1] = min(dp[0], dp[1]) + 1
		}
		dp = nextDp
	}
	return min(dp[0], dp[1])
}

func main() {
	fmt.Println(removeAlmostEqualCharacters("abddez"))
}
