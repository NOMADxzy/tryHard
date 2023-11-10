package main

import (
	"fmt"
	"sort"
)

func getVal(sums []int, i int, target int) int {
	return i*target - sums[i] + sums[len(sums)-1] - sums[i+1] - (len(sums)-2-i)*target

}

func minOperations(grid [][]int, x int) int {
	m, n := len(grid), len(grid[0])
	arr := make([]int, m*n)
	sums := make([]int, m*n+1)
	idx := 0

	delta := grid[0][0] % x

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr[idx] = grid[i][j]
			if arr[idx]%x != delta {
				return -1
			}
			idx++
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 1; i <= len(arr); i++ {
		sums[i] += sums[i-1] + arr[i-1]
	}

	minVal := 1<<31 - 1
	left, right := 0, len(arr)-1
	for left < right-1 {
		mid := (left + right) / 2
		v1, v2, v3 := getVal(sums, mid-1, arr[mid-1]), getVal(sums, mid, arr[mid]), getVal(sums, mid+1, arr[mid+1])
		if v1 > v2 && v2 > v3 {
			left = mid
		} else if v1 < v2 && v2 < v3 {
			right = mid
		} else {
			minVal = v2
			break
		}
	}
	if val := getVal(sums, left, arr[left]); val < minVal {
		minVal = val
	}
	if val := getVal(sums, right, arr[right]); val < minVal {
		minVal = val
	}
	return minVal / x
}

func main() {
	grid := [][]int{{2, 2}, {2, 2}}
	fmt.Println(minOperations(grid, 2))
}
