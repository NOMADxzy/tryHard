package main

import (
	"fmt"
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	pair := make([][]int, len(difficulty))
	for i := 0; i < len(difficulty); i++ {
		pair[i] = []int{difficulty[i], i}
	}
	sort.Slice(pair, func(i, j int) bool {
		return pair[i][0] < pair[j][0]
	})
	for i := 0; i < len(pair); i++ {
		pair[i][1] = profit[pair[i][1]]
	}
	idx := 0
	for i := 0; i < len(pair); i++ {

		if idx > 0 && pair[i][1] <= pair[idx-1][1] {
			continue
		} else {
			pair[idx][0], pair[idx][1] = pair[i][0], pair[i][1]
			idx++
		}
	}
	pair = pair[:idx]

	sort.Slice(worker, func(i, j int) bool {
		return worker[i] < worker[j]
	})

	ans := 0
	i := 0
	for i < len(worker) && worker[i] < pair[0][0] {
		i++
	}
	leftBorder := 0
	for ; i < len(worker); i++ {
		w := worker[i]
		l, r := leftBorder, len(pair)-1
		if w >= pair[r][0] {
			ans += pair[r][1]
		} else {
			for l < r {
				mid := (l + r) / 2
				if w >= pair[mid][0] {
					l = mid + 1
				} else {
					r = mid
				}
			}
			ans += pair[l-1][1]
			leftBorder = l - 1
		}
	}
	return ans
}

func main() {
	difficulty := []int{68, 35, 52, 47, 86}
	profit := []int{67, 17, 1, 81, 3}
	worker := []int{92, 10, 85, 84, 82}
	fmt.Println(maxProfitAssignment(difficulty, profit, worker))
}
