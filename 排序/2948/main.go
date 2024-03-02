package main

import (
	"fmt"
	"sort"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {
	m := len(nums)
	pairs := make([][2]int, m)

	for i := 0; i < m; i++ {
		pairs[i][0] = nums[i]
		pairs[i][1] = i
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	for i := 0; i < m; {
		j := i + 1
		for j < m && pairs[j][0]-pairs[j-1][0] <= limit {
			j++
		}
		j--
		var tmp []int
		for k := i; k <= j; k++ {
			tmp = append(tmp, pairs[k][1])
		}
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		for k := i; k <= j; k++ {
			nums[tmp[k-i]] = pairs[k][0]
		}
		i = j + 1
	}
	return nums
}

func main() {
	fmt.Println(lexicographicallySmallestArray([]int{1, 5, 3, 9, 8, 1}, 2))
}
