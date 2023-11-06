package main

import (
	"fmt"
	"sort"
)

func upCheck(arr []int, i int) {
	f := (i - 1) / 2
	if f >= 0 && arr[f] > arr[i] {
		arr[f], arr[i] = arr[i], arr[f]
		upCheck(arr, f)
	}
}

func downCheck(arr []int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	lar := i
	if left < len(arr) && arr[left] < arr[lar] {
		lar = left
	}
	if right < len(arr) && arr[right] < arr[lar] {
		lar = right
	}
	if lar != i {
		arr[lar], arr[i] = arr[i], arr[lar]
		downCheck(arr, lar)
	}
}

func maxEvents(events [][]int) int {
	var ans int
	var queue []int
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	curDay := 1
	for i := 0; i < len(events) || len(queue) > 0; curDay++ {
		for ; i < len(events) && events[i][0] == curDay; i++ {
			queue = append(queue, events[i][1])
			upCheck(queue, len(queue)-1)
		}
		for len(queue) > 0 && queue[0] < curDay {
			queue[0] = queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			downCheck(queue, 0)
		}
		if len(queue) > 0 {
			queue[0] = queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			downCheck(queue, 0)
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(maxEvents([][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}}))
}
