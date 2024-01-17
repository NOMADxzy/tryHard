package main

import (
	"fmt"
	"sort"
)

func smallestChair(times [][]int, targetFriend int) int {
	// [到期时间，idx]
	idxs := make([]int, len(times))
	for i := 0; i < len(idxs); i++ {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return times[idxs[i]][0] < times[idxs[j]][0]
	})
	var upCheck func(arr [][]int, i int)
	var downCheck func(arr [][]int, i int)
	upCheck = func(arr [][]int, i int) {
		f := (i - 1) / 2
		if f >= 0 && arr[i][0] < arr[f][0] {
			arr[i], arr[f] = arr[f], arr[i]
			upCheck(arr, f)
		}
	}
	downCheck = func(arr [][]int, i int) {
		left, right := 2*i+1, 2*i+2
		largest := i
		if left < len(arr) && arr[left][0] < arr[largest][0] {
			largest = left
		}
		if right < len(arr) && arr[right][0] < arr[largest][0] {
			largest = right
		}
		if largest != i {
			arr[largest], arr[i] = arr[i], arr[largest]
			downCheck(arr, largest)
		}
	}
	var hp1, hp2 [][]int
	n := len(times)
	for i := 0; i < n; i++ {
		hp2 = append(hp2, []int{i, 0})
		upCheck(hp2, len(hp2)-1)
	}
	for i := 0; i < len(idxs); i++ {
		time, end_time := times[idxs[i]][0], times[idxs[i]][1]
		for len(hp1) > 0 && hp1[0][0] <= time {
			id := hp1[0][1]
			hp1[0] = hp1[len(hp1)-1]
			hp1 = hp1[:len(hp1)-1]
			downCheck(hp1, 0)
			hp2 = append(hp2, []int{id, 0})
			upCheck(hp2, len(hp2)-1)
		}
		id := hp2[0][0]
		if idxs[i] == targetFriend {
			return id
		}
		hp2[0] = hp2[len(hp2)-1]
		hp2 = hp2[:len(hp2)-1]
		downCheck(hp2, 0)
		hp1 = append(hp1, []int{end_time, id})
		upCheck(hp1, len(hp1)-1)
	}
	return -1
}

func main() {
	fmt.Println(smallestChair([][]int{{3, 10}, {1, 5}, {2, 6}}, 0))
}
