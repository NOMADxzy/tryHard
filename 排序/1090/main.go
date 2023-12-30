package main

import (
	"fmt"
	"sort"
)

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	n := len(values)
	idxs := make([]int, n)
	for i := 0; i < n; i++ {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return values[idxs[i]] > values[idxs[j]]
	})
	ans := 0
	labelCnt := map[int]int{}
	cnt := 0
	for i := 0; i < n && cnt < numWanted; i++ {
		idx := idxs[i]
		if labelCnt[labels[idx]] < useLimit {
			labelCnt[labels[idx]]++
			cnt++
			ans += values[idx]
		}
	}
	return ans
}

func main() {
	fmt.Println(largestValsFromLabels([]int{2, 1, 9, 3, 5}, []int{1, 2, 2, 2, 2}, 1, 3))
}
