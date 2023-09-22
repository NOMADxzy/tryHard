package main

import "fmt"

func downCheck(arr [][]int, i int) {
	l := 2*i + 1
	r := 2*i + 2
	smallest := i
	if l < len(arr) && arr[l][0] < arr[smallest][0] {
		smallest = l
	}
	if r < len(arr) && arr[r][0] < arr[smallest][0] {
		smallest = r
	}
	if smallest != i {
		arr[smallest], arr[i] = arr[i], arr[smallest]
		downCheck(arr, smallest)
	}
}

func upCheck(arr [][]int, i int) {
	f := (i - 1) / 2
	if f >= 0 && arr[f][0] > arr[i][0] {
		arr[f], arr[i] = arr[i], arr[f]
		upCheck(arr, f)
	}
}

func rangeSum(nums []int, n int, left int, right int) int {
	var priorityQueue [][]int
	for i := 0; i < len(nums); i++ {
		priorityQueue = append(priorityQueue, []int{nums[i], i})
		upCheck(priorityQueue, len(priorityQueue)-1)
	}

	var sortedArr []int

	for len(sortedArr) < right {
		e := priorityQueue[0]
		sortedArr = append(sortedArr, e[0])
		expandIdx := e[1] + 1

		if expandIdx >= len(nums) {
			priorityQueue[0] = priorityQueue[len(priorityQueue)-1]
			priorityQueue = priorityQueue[:len(priorityQueue)-1]
			downCheck(priorityQueue, 0)
		} else {
			priorityQueue[0] = []int{e[0] + nums[expandIdx], expandIdx}
			downCheck(priorityQueue, 0)
		}
	}
	sum := 0
	for i := left - 1; i < right; i++ {
		sum = (sum + sortedArr[i]) % 1000000007
	}
	return sum
}

func main() {
	fmt.Print(rangeSum([]int{1, 2, 3, 4}, 4, 1, 4))
}
