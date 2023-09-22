package main

import "fmt"

func downCheck(arr [][]int, i int, size int) {
	l := 2*i + 1
	r := 2*i + 2
	minest := i
	if l < size && arr[l][0]*arr[minest][1] < arr[minest][0]*arr[l][1] {
		minest = l
	}
	if r < size && arr[r][0]*arr[minest][1] < arr[minest][0]*arr[r][1] {
		minest = r
	}
	if minest != i {
		arr[minest], arr[i] = arr[i], arr[minest]
		downCheck(arr, minest, size)
	}
}

func upCheck(arr [][]int, i int) {
	f := (i - 1) / 2
	if f >= 0 && arr[f][0]*arr[i][1] > arr[i][0]*arr[f][1] {
		arr[f], arr[i] = arr[i], arr[f]
		upCheck(arr, f)
	}
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	pointers := make([]int, len(arr)-1)
	for i := 0; i < len(pointers); i++ {
		pointers[i] = len(arr) - 1
	}
	var priorityQueue [][]int
	for i := 0; i < len(pointers); i++ {
		priorityQueue = append(priorityQueue, []int{arr[i], arr[pointers[i]], i})
		upCheck(priorityQueue, i)
	}
	cnt := 0
	for cnt < k {
		cur := priorityQueue[0]
		cnt++
		if cnt == k {
			return cur[:2]
		}
		pointer_idx := cur[2]
		pointers[pointer_idx]--

		if pointers[pointer_idx] <= pointer_idx {
			if len(priorityQueue) <= 1 {
				break
			}
			priorityQueue[0] = priorityQueue[len(priorityQueue)-1]
			priorityQueue = priorityQueue[:len(priorityQueue)-1]
			downCheck(priorityQueue, 0, len(priorityQueue))
			continue
		} else {
			priorityQueue[0] = []int{arr[pointer_idx], arr[pointers[pointer_idx]], pointer_idx}
			downCheck(priorityQueue, 0, len(priorityQueue))
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Print(kthSmallestPrimeFraction([]int{1, 13, 17, 59}, 6))
}
