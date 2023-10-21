package main

import (
	"fmt"
	"sort"
)

func numFactoredBinaryTrees(arr []int) int {
	m := len(arr)
	MOD := 1000000007
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	dp := make([]int, m)
	dp[0] = 1
	ans := 1

	idxMap := map[int]int{arr[0]: 0}

	for i := 1; i < m; i++ {
		total := 1
		limit := 0
		for j := i - 1; j >= 0 && arr[j] > limit; j-- {
			if arr[i]%arr[j] == 0 {
				if cntIdx, ok := idxMap[arr[i]/arr[j]]; ok {
					cnt := (dp[j] * dp[cntIdx]) % MOD
					if j != cntIdx {
						cnt = (cnt * 2) % MOD
					}
					total = (total + cnt) % MOD
					limit = arr[cntIdx]
				}
			}
		}
		dp[i] = total
		ans = (ans + total) % MOD
		idxMap[arr[i]] = i
	}
	return ans
}

func main() {
	fmt.Println(numFactoredBinaryTrees([]int{45, 42, 2, 18, 23, 1170, 12, 41, 40, 9, 47, 24, 33, 28, 10, 32, 29, 17, 46, 11, 759, 37, 6, 26, 21, 49, 31, 14, 19, 8, 13, 7, 27, 22, 3, 36, 34, 38, 39, 30, 43, 15, 4, 16, 35, 25, 20, 44, 5, 48}))
}
