package main

import (
	"fmt"
	"sort"
)

func upCheck(arr [][]int, i int) {
	f := (i - 1) / 2
	if f >= 0 && arr[i][0] > arr[f][0] {
		arr[f], arr[i] = arr[i], arr[f]
		upCheck(arr, f)
	}
}

func downCheck(arr [][]int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < len(arr) && arr[left][0] > arr[largest][0] {
		largest = left
	}
	if right < len(arr) && arr[right][0] > arr[largest][0] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		downCheck(arr, largest)
	}
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	m := len(courses)
	var pqueue [][]int
	var sumTime int
	for i := 0; i < m; i++ {
		curCourse := courses[i]
		if curCourse[0]+sumTime <= curCourse[1] {
			pqueue = append(pqueue, curCourse)
			sumTime += curCourse[0]
			upCheck(pqueue, len(pqueue)-1)
		} else if len(pqueue) > 0 && pqueue[0][0] > curCourse[0] {
			sumTime += curCourse[0] - pqueue[0][0]
			pqueue[0] = curCourse
			downCheck(pqueue, 0)
		}
	}
	return len(pqueue)
}

func main() {
	fmt.Println(scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}))
}
