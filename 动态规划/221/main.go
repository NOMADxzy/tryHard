package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j < n+1; j++ {
			dp[i][j] = []int{0, 0, 0}
		}
	}

	max := 0

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if matrix[i-1][j-1] == '1' {
				tuple := dp[i][j]
				tuple[0] = dp[i-1][j][0] + 1
				tuple[1] = dp[i][j-1][1] + 1

				border := min(tuple[0], tuple[1])
				tuple[2] = min(dp[i-1][j-1][2]+1, border)
				if max < tuple[2] {
					max = tuple[2]
				}
			}
		}
	}
	return max * max
}

func main() {
	fmt.Println(maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '0'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))

	//fmt.Println(maximalSquare([][]byte{
	//	{'0', '1'},
	//	{'1', '0'},
	//}))
}
