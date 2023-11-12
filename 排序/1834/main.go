package main

import (
	"fmt"
	"sort"
)

func downCheck(arr [][]int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	best := i
	if left < len(arr) && (arr[left][0] < arr[best][0] || arr[left][0] == arr[best][0] && arr[left][1] < arr[best][1]) {
		best = left
	}
	if right < len(arr) && (arr[right][0] < arr[best][0] || arr[right][0] == arr[best][0] && arr[right][1] < arr[best][1]) {
		best = right
	}
	if best != i {
		arr[i], arr[best] = arr[best], arr[i]
		downCheck(arr, best)
	}
}

func upCheck(arr [][]int, i int) {
	f := (i - 1) / 2
	if f >= 0 && (arr[i][0] < arr[f][0] || arr[f][0] == arr[i][0] && arr[i][1] < arr[f][1]) {
		arr[i], arr[f] = arr[f], arr[i]
		upCheck(arr, f)
	}
}

func getOrder(tasks [][]int) []int {
	m := len(tasks)
	ans := make([]int, m)
	idxs := make([]int, m)
	for i := 0; i < m; i++ {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return tasks[idxs[i]][0] < tasks[idxs[j]][0] || tasks[idxs[i]][0] < tasks[idxs[j]][0] && idxs[i] < idxs[j]
	})

	var queue [][]int

	j := 0
	startTime := 0
	r := 0
	for i := tasks[idxs[0]][0]; j < m; {
		for ; r < m && tasks[idxs[r]][0] == i; r++ {
			queue = append(queue, []int{tasks[idxs[r]][1], idxs[r]})
			upCheck(queue, len(queue)-1)
		}
		if len(queue) > 0 && i >= startTime {
			t := queue[0]
			startTime = i + t[0]

			ans[j] = t[1]
			j++

			queue[0] = queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			downCheck(queue, 0)
		}
		i++
		if startTime > i {
			i = startTime
		}
		if r < m && (len(queue) == 0 || tasks[idxs[r]][0] < i) {
			i = tasks[idxs[r]][0]
		}

	}
	return ans
}

func main() {
	fmt.Println(getOrder([][]int{{100, 100}, {1000000000, 1000000000}}))
}
