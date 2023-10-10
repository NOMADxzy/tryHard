package main

import "fmt"

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	sums := make([][]int, m)
	for i := 0; i < m; i++ {
		sums[i] = make([]int, n)
	}

	maxLen := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if maxLen < 1 && mat[i][j] <= threshold {
				maxLen = 1
			}
			if i == 0 && j == 0 {
				sums[i][j] = mat[i][j]
			} else if i == 0 {
				sums[i][j] = sums[i][j-1] + mat[i][j]
			} else if j == 0 {
				sums[i][j] = sums[i-1][j] + mat[i][j]
			} else {
				sums[i][j] = sums[i-1][j] + sums[i][j-1] + mat[i][j] - sums[i-1][j-1]
			}
			if i == j && sums[i][j] <= threshold && (i+1) > maxLen {
				maxLen = i + 1
			}
		}
	}
	if maxLen == 0 {
		return 0
	}

	for i := 0; maxLen < m-i; i++ {
		for j := 0; maxLen < n-j; j++ {
			var l int
			for l = maxLen + 1; i+l-1 < m && j+l-1 < n; l++ {
				var s0, s1, s2 int
				if i == 0 && j == 0 {
				} else if i == 0 {
					s2 = sums[i+l-1][j-1]
				} else if j == 0 {
					s1 = sums[i-1][j+l-1]
				} else {
					s0, s1, s2 = sums[i-1][j-1], sums[i-1][j+l-1], sums[i+l-1][j-1]
				}

				s := sums[i+l-1][j+l-1] - s1 - s2 + s0
				if s <= threshold {
					maxLen = l
				} else {
					break
				}
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(maxSideLength([][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}}, 1))
}
