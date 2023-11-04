package main

import (
	"fmt"
	"sort"
)

func findDiagonalOrder(nums [][]int) []int {
	size := 0
	for i := 0; i < len(nums); i++ {
		size += len(nums[i])
	}
	pos := make([][]int, size)
	idx := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			pos[idx] = []int{i, j}
			idx++
		}
	}
	sort.Slice(pos, func(i, j int) bool {
		return pos[i][0]+pos[i][1] < pos[j][0]+pos[j][1] || pos[i][0]+pos[i][1] == pos[j][0]+pos[j][1] && pos[i][1] < pos[j][1]
	})

	ans := make([]int, size)
	for i := 0; i < size; i++ {
		ans[i] = nums[pos[i][0]][pos[i][1]]
	}
	return ans
}

func main() {
	arr := [][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}
	fmt.Println(findDiagonalOrder(arr))
}
