package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(arr [][]int) []int {
	m, n := len(arr), len(arr[0])
	dp := make([][]int, m)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
	}
	for j := 1; j <= n; j++ {
		for i := 0; i < m; i++ {
			dp[i][j] = arr[i][j-1] + dp[i][j-1]
		}
	}
	ans := make([]int, n)
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			continue
		}
		target := i * i / 2
		for l := 0; l+i-1 < n; l++ {
			r := l + i - 1
			sum := 0
			for tmp := 0; tmp < i; tmp++ {
				sum += dp[tmp][r+1] - dp[tmp][l]
			}
			t, b := 0, i-1
			for {
				if sum == target {
					ans[i-1]++
				}
				sum -= dp[t][r+1] - dp[t][l]
				t++
				b++
				if b == n {
					break
				}
				sum += dp[b][r+1] - dp[b][l]
			}
		}
	}
	return ans
}

func main() {
	//arr := [][]int{
	//	{1, 0, 1, 0},
	//	{0, 1, 0, 1},
	//	{1, 1, 0, 0},
	//	{0, 0, 1, 1},
	//}
	scan := bufio.NewScanner(os.Stdin)
	var n int
	_, _ = fmt.Scan(&n)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		scan.Scan()
		s := scan.Text()
		for j := 0; j < len(s); j++ {
			if s[j] == '1' {
				arr[i][j] = 1
			}
		}
	}
	ans := solve(arr)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}
