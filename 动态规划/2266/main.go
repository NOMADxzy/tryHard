package main

import "fmt"

func countTexts(pressedKeys string) int {
	m := len(pressedKeys)
	dp := make([]int, m+1)
	dp[0] = 1
	MOD := 1000000007

	for i := 1; i <= m; i++ {
		preMax := 3
		if pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9' {
			preMax++
		}
		cnt := 0
		for j := i - 1; j >= 0 && pressedKeys[j] == pressedKeys[i-1] && i-j <= preMax; j-- {
			cnt = (cnt + dp[j]) % MOD
		}
		dp[i] = cnt
	}
	return dp[m]
}

func main() {
	fmt.Println(countTexts("444479999555588866"))
}
