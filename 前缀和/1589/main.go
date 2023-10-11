package main

import (
	"fmt"
	"sort"
)

func maxSumRangeQuery(nums []int, requests [][]int) int {
	m := len(nums)
	deferArr := make([]int, m)
	for i := 0; i < len(requests); i++ {
		rq := requests[i]
		deferArr[rq[0]]++
		if rq[1] < m-1 {
			deferArr[rq[1]+1]--
		}
	}
	for i := 1; i < m; i++ {
		deferArr[i] = deferArr[i-1] + deferArr[i]
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	sortIdxs := make([]int, m)
	for i := 0; i < m; i++ {
		sortIdxs[i] = i
	}
	sort.Slice(sortIdxs, func(i, j int) bool {
		return deferArr[sortIdxs[i]] > deferArr[sortIdxs[j]]
	})

	ans := 0
	for i := 0; i < m; i++ {
		ans += (deferArr[sortIdxs[i]] * nums[i]) % 1000000007
		ans %= 1000000007
	}
	return ans
}

func main() {
	fmt.Println(maxSumRangeQuery([]int{1, 2, 3, 4, 5}, [][]int{{1, 3}, {0, 1}}))
}
