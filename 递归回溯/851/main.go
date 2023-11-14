package main

import "fmt"

func findBest(richerList [][]int, pos int, ans []int, quiet []int) int {
	if ans[pos] >= 0 {
		return ans[pos]
	}
	best := pos
	for _, p := range richerList[pos] {
		if val := findBest(richerList, p, ans, quiet); quiet[val] < quiet[best] {
			best = val
		}
	}
	ans[pos] = best
	return best
}

func loudAndRich(richer [][]int, quiet []int) []int {
	m := len(quiet)
	richerList := make([][]int, m)

	ans := make([]int, m)
	for i := 0; i < m; i++ {
		ans[i] = -1
	}
	for _, r := range richer {
		richerList[r[1]] = append(richerList[r[1]], r[0])
	}

	for i := 0; i < m; i++ {
		findBest(richerList, i, ans, quiet)
	}
	return ans
}

func main() {
	fmt.Println(loudAndRich([][]int{{0, 1}, {1, 2}}, []int{0, 1, 2}))
}
