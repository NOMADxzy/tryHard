package main

import "fmt"

func partQuickSort(arr []int, start, end int, k int) {
	if end == start {
		return
	}
	guard := arr[start]
	l, r := start, end
	for l < r {
		for l < r && arr[r] >= guard { //TODO 易错，注意要有“=”
			r--
		}
		arr[l], arr[r] = arr[r], arr[l]
		for l < r && arr[l] <= guard {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	if l == k {
		return
	} else if l < k {
		partQuickSort(arr, l+1, end, k)
	} else {
		partQuickSort(arr, start, l-1, k)
	}
}

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	valArr := make([]int, m*n)
	valArr[0] = matrix[0][0]
	idx := 1
	for i := 1; i < m; i++ {
		matrix[i][0] = matrix[i-1][0] ^ matrix[i][0]
		valArr[idx] = matrix[i][0]
		idx++
	}
	for j := 1; j < n; j++ {
		matrix[0][j] = matrix[0][j-1] ^ matrix[0][j]
		valArr[idx] = matrix[0][j]
		idx++
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i][j] ^ matrix[i-1][j] ^ matrix[i][j-1] ^ matrix[i-1][j-1]
			valArr[idx] = matrix[i][j]
			idx++
		}
	}
	partQuickSort(valArr, 0, len(valArr)-1, len(valArr)-k)
	return valArr[len(valArr)-k]
}

func main() {
	matrix := [][]int{{8, 10, 5, 8, 5, 7, 6, 0, 1, 4, 10, 6, 4, 3, 6, 8, 7, 9, 4, 2}}
	fmt.Println(kthLargestValue(matrix, 2))
}
