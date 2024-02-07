package main

import "fmt"

func longestString(x int, y int, z int) int {
	hist := map[int]int{}
	//strs := []string{"AA", "BB", "AB"}
	nexts := [][]int{{1}, {0, 2}, {0, 2}}
	var dfs func(xRemain, yRemain, zRemain int, pre int) int
	dfs = func(xRemain, yRemain, zRemain int, pre int) int {
		key := 51*2701*pre + 2701*xRemain + (51*yRemain + zRemain)
		if v, ok := hist[key]; ok {
			return v
		}
		best := 0
		for _, idx := range nexts[pre] {
			if idx == 0 && xRemain > 0 {
				best = max(best, 2+dfs(xRemain-1, yRemain, zRemain, 0))
			}
			if idx == 1 && yRemain > 0 {
				best = max(best, 2+dfs(xRemain, yRemain-1, zRemain, 1))
			}
			if idx == 2 && zRemain > 0 {
				best = max(best, 2+dfs(xRemain, yRemain, zRemain-1, 2))
			}
		}
		hist[key] = best
		return best
	}
	cnts := []int{x, y, z}
	ans := 0
	for i := 0; i < 3; i++ {
		if cnts[i] > 0 {
			cnts[i]--
			ans = max(ans, 2+dfs(cnts[0], cnts[1], cnts[2], i))
			cnts[i]++
		}
	}
	return ans
}

func main() {
	fmt.Println(longestString(2, 5, 1))
}
