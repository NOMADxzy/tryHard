package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m, n := len(students), len(students[0])
	scoreMatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		scoreMatrix[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			val := 0
			for k := 0; k < n; k++ {
				if students[i][k] == mentors[j][k] {
					val++
				}
			}
			scoreMatrix[i][j] = val
		}
	}

	dp := map[int]int{0: 0}

	markTop := 1<<m - 1

	for mark := 0; mark < markTop; mark++ {

		cnt1 := 0
		for i := mark; i > 0; i = i >> 1 {
			if i%2 == 1 {
				cnt1++
			}
		}

		mask := 1
		for p := 0; p < m; p++ {
			if mark&mask == 0 {
				dp[mark+mask] = max(dp[mark+mask], dp[mark]+scoreMatrix[cnt1][p])
			}
			mask *= 2
		}
	}
	return dp[markTop]
}

func main() {
	fmt.Println(maxCompatibilitySum([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 1}}, [][]int{{1, 0, 0}, {0, 0, 1}, {1, 1, 0}}))
}
