package main

import "fmt"

func sumSubarrayMins(arr []int) int {
	m := len(arr)
	MOD := 1000000007
	var stack []int
	beforeSmaller := make([]int, m)
	for i := m - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			idx := stack[len(stack)-1]
			beforeSmaller[idx] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		beforeSmaller[i] = -1
	}

	dp := make([]int, m) // dp[i]表示以i为右端点所有子数组的最小值之和
	dp[0] = arr[0]
	ans := arr[0]

	for i := 1; i < m; i++ {
		j := beforeSmaller[i]
		if j >= 0 {
			dp[i] = (((i-j)*arr[i])%MOD + dp[j]) % MOD
		} else {
			dp[i] = ((i + 1) * arr[i]) % MOD
		}
		ans = (ans + dp[i]) % MOD
	}
	return ans
}

func main() {
	fmt.Println(sumSubarrayMins([]int{11, 81, 94, 43, 3}))
}
