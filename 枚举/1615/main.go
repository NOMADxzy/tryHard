package main

import (
	"fmt"
	"sort"
)

func maximalNetworkRank(n int, roads [][]int) int {
	cnts := make([]int, n)
	hist := map[int]bool{}
	for _, road := range roads {
		cnts[road[0]]++
		cnts[road[1]]++
		hist[100*road[0]+road[1]] = true
		hist[100*road[1]+road[0]] = true
	}
	idxs := make([]int, n)
	for i := 0; i < n; i++ {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return cnts[idxs[i]] > cnts[idxs[j]]
	})
	maxVal := cnts[idxs[1]]
	best := cnts[idxs[0]] + cnts[idxs[1]]
	if hist[100*idxs[0]+idxs[1]] {
		best--
	} else {
		return best
	}
	for i := 0; i < n && cnts[idxs[i]] >= maxVal; i++ {
		for j := i + 1; j < n && cnts[idxs[j]] >= maxVal; j++ {
			cur := cnts[idxs[i]] + cnts[idxs[j]]
			if !hist[100*idxs[i]+idxs[j]] && cur+1 > best {
				return cur
			} else if cur+1 == best {
				return best
			}
		}
	}
	return best
}

func main() {
	fmt.Println(maximalNetworkRank(3, [][]int{{0, 1}, {0, 2}}))
}
