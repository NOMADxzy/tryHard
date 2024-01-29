package main

import (
	"fmt"
	"sort"
)

func minimumRemoval(beans []int) int64 {
	sort.Slice(beans, func(i, j int) bool {
		return beans[i] < beans[j]
	})
	total := int64(0)
	for i := 0; i < len(beans); i++ {
		total += int64(beans[i])
	}
	m := len(beans)
	rightSum := total - int64(m*beans[0])
	leftTotalSum := int64(0)
	leftSum := int64(0)
	ans := total
	for i := 0; i < m; i++ {
		if i > 0 && beans[i-1] < beans[i] {
			leftSum = leftTotalSum
		}
		rightSum = total - int64((m-i)*beans[i])
		ans = min(ans, leftSum+rightSum)
		total -= int64(beans[i])
		leftTotalSum += int64(beans[i])
	}
	return ans
}

func main() {
	fmt.Println(minimumRemoval([]int{2, 5, 5}))
}
