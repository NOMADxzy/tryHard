package main

import (
	"fmt"
	"sort"
)

// 树状数组 离散化
func countRangeSum(nums []int, lower int, upper int) int {
	m := len(nums)
	set := map[int]bool{}
	preSum := make([]int, m)
	preSum[0] = nums[0]
	set[preSum[0]] = true
	set[preSum[0]-lower] = true
	set[preSum[0]-upper] = true
	for i := 1; i < m; i++ {
		preSum[i] = preSum[i-1] + nums[i]
		set[preSum[i]] = true
		set[preSum[i]-lower] = true
		set[preSum[i]-upper] = true
	}
	var uniqueVals []int
	for v, _ := range set {
		uniqueVals = append(uniqueVals, v)
	}
	sort.Slice(uniqueVals, func(i, j int) bool {
		return uniqueVals[i] < uniqueVals[j]
	})
	idxs := map[int]int{}
	idx := 1
	for i := 0; i < len(uniqueVals); i++ {
		idxs[uniqueVals[i]] = idx
		idx++
	}
	tree := make([]int, 3*m+4)
	lowBit := func(x int) int {
		return x & (-x)
	}
	update := func(x int, val int) {
		for x < len(tree) {
			tree[x] += val
			x += lowBit(x)
		}
	}
	query := func(x int) int {
		sum := 0
		for x > 0 {
			sum += tree[x]
			x -= lowBit(x)
		}
		return sum
	}
	ans := 0
	for i := 0; i < m; i++ {
		if preSum[i] >= lower && preSum[i] <= upper {
			ans++
		}
		l, r := idxs[preSum[i]-upper], idxs[preSum[i]-lower]
		ans += query(r) - query(l-1)

		update(idxs[preSum[i]], 1)
	}
	return ans
}

// low < preSum - x < high
// preSum - high < x < preSum - low

func main() {
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
}
