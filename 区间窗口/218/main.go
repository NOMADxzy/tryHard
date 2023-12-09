package main

import (
	"fmt"
	"sort"
)

func upCheck(arr [][]int, i int) {
	f := (i - 1) / 2
	if f >= 0 && arr[i][0] > arr[f][0] {
		arr[i], arr[f] = arr[f], arr[i]
		upCheck(arr, f)
	}
}

func downCheck(arr [][]int, i int) {
	left, right := 2*i+1, 2*i+2
	lar := i
	if left < len(arr) && arr[left][0] > arr[lar][0] {
		lar = left
	}
	if right < len(arr) && arr[right][0] > arr[lar][0] {
		lar = right
	}
	if lar != i {
		arr[lar], arr[i] = arr[i], arr[lar]
		downCheck(arr, lar)
	}
}

func getSkyline(buildings [][]int) [][]int {
	var queue [][]int
	sort.Slice(buildings, func(i, j int) bool {
		return buildings[i][0] < buildings[j][0]
	})
	left := buildings[0][0]
	right := buildings[len(buildings)-1][0]
	var ans [][]int
	idx := 0
	for i := left; len(queue) > 0 || i <= right; { // 扫描线i
		for idx < len(buildings) && i >= buildings[idx][0] {
			building := buildings[idx]
			queue = append(queue, []int{building[2], building[1]})
			upCheck(queue, len(queue)-1)
			idx++
		}

		for len(queue) > 0 && queue[0][1] <= i {
			queue[0] = queue[len(queue)-1]
			queue = queue[0 : len(queue)-1]
			downCheck(queue, 0)
		}
		if len(queue) == 0 && ans[len(ans)-1][1] != 0 {
			ans = append(ans, []int{i, 0})
		} else if len(ans) == 0 || queue[0][0] != ans[len(ans)-1][1] {
			ans = append(ans, []int{i, queue[0][0]})
		}
		i = 1<<31 - 1
		if len(queue) > 0 {
			i = queue[0][1]
		}
		if idx < len(buildings) && buildings[idx][0] <= i {
			i = buildings[idx][0]
		}
	}
	return ans
}

func main() {
	builds := [][]int{{3, 7, 8}, {3, 8, 7}, {3, 9, 6}, {3, 10, 5}}
	fmt.Println(getSkyline(builds))
}
