package main

import "fmt"

func cntNext(childs [][]int, cntChilds []int, pos int) int {
	cnt := 1
	for _, nextP := range childs[pos] {
		cnt += cntNext(childs, cntChilds, nextP)
	}
	cntChilds[pos] = cnt
	return cnt
}

func countHighestScoreNodes(parents []int) int {
	m := len(parents)
	childs := make([][]int, m)
	cntChilds := make([]int, m)

	for i, p := range parents {
		if p >= 0 {
			childs[p] = append(childs[p], i)
		}
	}

	total := cntNext(childs, cntChilds, 0)
	ans := 0
	curMax := 0

	for i := 0; i < m; i++ {
		score := 1
		for _, c := range childs[i] {
			score *= cntChilds[c]
		}
		if i > 0 {
			score *= total - cntChilds[i]
		}
		if score > curMax {
			curMax = score
			ans = 1
		} else if score == curMax {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(countHighestScoreNodes([]int{-1, 2, 0, 2, 0}))
}
