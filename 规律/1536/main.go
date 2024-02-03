package main

import "fmt"

func minSwaps(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	vals := make([]int, len(grid))
	for i := 0; i < m; i++ {
		cnt0 := 0
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] != 0 {
				break
			}
			cnt0++
		}
		vals[i] = cnt0
	}
	ans := 0
	for i := 0; i < m; i++ {
		target := n - 1 - i
		var j int
		for j = i; j < m && vals[j] < target; j++ {
		}
		if j == m {
			return -1
		} else {
			for j > i {
				vals[j-1], vals[j] = vals[j], vals[j-1]
				j--
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minSwaps([][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}))
}
