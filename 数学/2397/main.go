package main

import (
	"fmt"
	"math/bits"
)

func maximumRows(matrix [][]int, numSelect int) int {
	m, n := len(matrix), len(matrix[0])
	nums := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nums[i] += matrix[i][j] << (n - 1 - j)
		}
	}
	ans := 0
	for state := 0; state < 1<<n; state++ {
		if bits.OnesCount(uint(state)) != numSelect {
			continue
		}
		cur := 0
		for i := 0; i < m; i++ {
			if nums[i]&state == nums[i] {
				cur++
			}
		}
		ans = max(ans, cur)
	}
	return ans
}

func main() {
	fmt.Println(maximumRows([][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}, 2))
}
