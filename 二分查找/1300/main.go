package main

import (
	"fmt"
	"sort"
)

func bSearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	if arr[l] >= target {
		return -1
	}
	for l < r {
		mid := (l + r) / 2
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}

func findBestValue(arr []int, target int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	m := len(arr)
	sums := make([]int, m+1)
	for i := 1; i <= m; i++ {
		sums[i] = sums[i-1] + arr[i-1]
	}

	left, right := 0, arr[m-1]
	if sums[m] <= target {
		return right
	} else {
		for left < right {
			mid := (left + right) / 2
			idx := bSearch(arr, mid)
			if sums[idx+1]+(m-1-idx)*mid > target {
				right = mid
			} else {
				left = mid + 1
			}
		}
		left--
		idx1 := bSearch(arr, left)
		idx2 := bSearch(arr, right)
		best := target - (sums[idx1+1] + (m-1-idx1)*left)
		if (sums[idx2+1]+(m-1-idx2)*right)-target < best {
			return right
		}
	}
	return left
}

func main() {
	fmt.Println(findBestValue([]int{3, 3, 3}, 1))
}
