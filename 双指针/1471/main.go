package main

import (
	"fmt"
	"sort"
)

func getStrongest(arr []int, k int) []int {
	if len(arr) == k {
		return arr
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	mid := (len(arr) - 1) / 2
	ans := make([]int, k)
	left, right := 0, len(arr)-1
	remain := k
	for remain > 0 {
		if arr[mid]-arr[left] > arr[right]-arr[mid] {
			t := arr[right] - arr[mid]
			l, r := left, mid
			for l < r {
				m := (l + r) / 2
				if arr[mid]-arr[m] > t {
					l = m + 1
				} else {
					r = m
				}
			}
			copy(ans[k-remain:], arr[left:l])
			left = l
		} else {
			t := arr[mid] - arr[left]
			l, r := mid, right
			for l < r {
				m := (l + r) / 2
				if arr[m]-arr[mid] >= t {
					r = m
				} else {
					l = m + 1
				}
			}
			r = max(r, right+1-remain)
			copy(ans[k-remain:], arr[r:right+1])
			right = r - 1
		}
		remain = k - (left + len(arr) - 1 - right)
	}
	return ans
}

func main() {
	fmt.Println(getStrongest([]int{-2, -4, -6, -8, -9, -7, -5, -3, -1}, 3))
}
