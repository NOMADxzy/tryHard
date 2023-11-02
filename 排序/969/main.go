package main

import "fmt"

func swap(arr []int, k int) {
	k--
	for i := 0; i <= k/2; i++ {
		arr[i], arr[k-i] = arr[k-i], arr[i]
	}
}

func pancakeSort(arr []int) []int {
	n := len(arr)
	cur := 1
	var opts []int
	var i int
	for ; cur <= n; cur++ {
		for i = 0; i < n && arr[i] != cur; i++ {
		}
		if i == n-cur {
			continue
		}
		opts = append(opts, i+1)
		swap(arr, i+1)
		opts = append(opts, n+1-cur)
		swap(arr, n+1-cur)
	}
	opts = append(opts, n)
	return opts
}

func main() {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
}
