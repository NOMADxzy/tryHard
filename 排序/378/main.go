package main

import "fmt"

func buildHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		check(arr, i, len(arr))
	}
}

func check(arr []int, i int, l int) {
	left := 2*i + 1
	right := 2*i + 2
	largetest := i
	if left < l && arr[left] > arr[largetest] {
		largetest = left
	}
	if right < l && arr[right] > arr[largetest] {
		largetest = right
	}
	if largetest != i {
		arr[i], arr[largetest] = arr[largetest], arr[i]
		check(arr, largetest, l)
	}
}

func kthSmallest(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	b := 1

	for b*m < k {
		b++
	}
	var arr []int
	for i := 0; i < m; i++ {
		for j := 0; j < b; j++ {
			if len(arr) < k {
				arr = append(arr, matrix[i][j])
				if len(arr) == k {
					buildHeap(arr)
				}
			} else if matrix[i][j] < arr[0] {
				arr[0] = matrix[i][j]
				check(arr, 0, k)
			}
		}
	}

	i := 0
	j := n - 1
	for j >= b {
		if matrix[i][j] > arr[0] {
			j--
			continue
		} else {
			for t := b; t <= j; t++ {
				if matrix[i][t] < arr[0] {
					arr[0] = matrix[i][t]
					check(arr, 0, k)
				}
			}
			i++
		}
	}
	return arr[0]
}

func main() {
	fmt.Println(kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 5))
}
