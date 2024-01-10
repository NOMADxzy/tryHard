package main

import "fmt"

func restoreArray(adjacentPairs [][]int) []int {
	nexts := map[int][]int{}
	for _, pair := range adjacentPairs {
		a, b := pair[0], pair[1]
		nexts[a] = append(nexts[a], b)
		nexts[b] = append(nexts[b], a)
	}
	cur := 0
	for v, next := range nexts {
		if len(next) == 1 {
			cur = v
			break
		}
	}
	pre := -1
	ans := make([]int, len(adjacentPairs)+1)
	for i := 0; i < len(ans); i++ {
		ans[i] = cur
		for _, v := range nexts[cur] {
			if v != pre {
				cur = v
				pre = ans[i]
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(restoreArray([][]int{{2, 1}, {3, 4}, {3, 2}}))
}
