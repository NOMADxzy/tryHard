package main

import (
	"fmt"
	"sort"
)

func minProcessingTime(processorTime []int, tasks []int) int {
	sort.Slice(processorTime, func(i, j int) bool {
		return processorTime[i] < processorTime[j]
	})
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})
	ans := 0
	j := 0
	for i := 0; i < len(tasks); i += 4 {
		ans = max(ans, processorTime[j]+tasks[i])
		j++
	}
	return ans
}

func main() {
	fmt.Println(minProcessingTime([]int{8, 10}, []int{2, 2, 3, 1, 8, 7, 4, 5}))
}
