package main

import "fmt"

func bombNext(bombs [][]int, pos int, mark []bool, dists [][]int) {
	mark[pos] = true
	for i := 0; i < len(bombs); i++ {
		if !mark[i] && bombs[pos][2]*bombs[pos][2] >= dists[pos][i] {
			bombNext(bombs, i, mark, dists)
		}
	}
}

func maximumDetonation(bombs [][]int) int {
	m := len(bombs)
	dists := make([][]int, m)
	for i := 0; i < m; i++ {
		dists[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			d := (bombs[i][0]-bombs[j][0])*(bombs[i][0]-bombs[j][0]) + (bombs[i][1]-bombs[j][1])*(bombs[i][1]-bombs[j][1])
			dists[i][j] = d
			dists[j][i] = d
		}
	}
	ans := 0
	mark := make([]bool, m)
	for i := 0; i < m; i++ {
		bombNext(bombs, i, mark, dists)
		cur := 0
		for j := 0; j < m; j++ {
			if mark[j] {
				cur++
			}
		}
		if cur > ans {
			ans = cur
		}
		clear(mark)
	}
	return ans
}

func main() {
	fmt.Println(maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}}))
}
