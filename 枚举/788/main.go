package main

import (
	"fmt"
	"strconv"
)

func rotatedDigits1(n int) int {

	w := 0
	pow := 1
	for i := n; i > 0; i /= 10 {
		w++
		pow *= 10
	}

	dp := make([][]int, w)
	for i := 0; i < w; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 3
	dp[0][1] = 4

	ans := 4
	for i := 1; i < w; i++ {
		dp[i][0] = dp[i-1][0] * 3
		dp[i][1] = dp[i-1][1]*7 + dp[i-1][0]*4
		ans += dp[i][1]
	}

	invalidLargerMap := []int{2, 1, 1, 1, 1, 1, 1, 1, 0, 0}
	validLargerMap := []int{4, 4, 3, 3, 3, 2, 1, 1, 1, 0}
	if w == 1 {
		return 4 - validLargerMap[n]
	}
	flag := false
	badAns := dp[w-2][1] // 首字母为0的去除
	for i := 0; i < w; i++ {
		pow /= 10
		v := n / pow
		n -= pow * v
		if w-2-i >= 0 {
			if flag {
				badAns += (invalidLargerMap[v] + validLargerMap[v]) * (dp[w-2-i][1] + dp[w-2-i][0])
			} else {
				badAns += invalidLargerMap[v]*dp[w-2-i][1] + validLargerMap[v]*(dp[w-2-i][1]+dp[w-2-i][0])
			}
		} else {
			if flag {
				badAns += invalidLargerMap[v] + validLargerMap[v]
			} else {
				badAns += validLargerMap[v]
			}
		}
		if v == 5 || v == 2 || v == 6 || v == 9 {
			flag = true
		}

	}
	return ans - badAns
}

var check = [10]int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}

func rotatedDigits(n int) (ans int) {
	for i := 1; i <= n; i++ {
		valid, diff := true, false
		for _, c := range strconv.Itoa(i) {
			if check[c-'0'] == -1 {
				valid = false
			} else if check[c-'0'] == 1 {
				diff = true
			}
		}
		if valid && diff {
			ans++
		}
	}
	return
}

func main() {
	fmt.Println(rotatedDigits(857))
}
