package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s string)int{
	ans := 0
	n := len(s)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if s[i] == '1'{
			dp[i][0] = 1
		}else{
			dp[i][1] = 1
			ans += 1
		}
	}
	for delta := 1; delta < n; delta++ {
		newDp := make([][2]int, n-delta)
		for i := 0; i < n-delta; i++ {
			l, r := i, i+delta
			if s[r] == '0'{
				newDp[l][0] = min(dp[l][0], dp[l][1]+1)
				newDp[l][1] = min(dp[l][0]+1, dp[l][1]+2)
			}else{
				newDp[l][1] = min(dp[l][1], dp[l][0]+1)
				newDp[l][0] = min(dp[l][1]+1, dp[l][0]+2)
			}
			if  (r-l+1) % 2 == 1 && newDp[l][1] % 2 == 1{
				ans += 1
			}
		}
		dp = newDp
	}
	return ans
}

func main() {
	//var n int
	var s string
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	_ = input.Text()
	input.Scan()
	s = input.Text()
	fmt.Println(solve(s))
}



