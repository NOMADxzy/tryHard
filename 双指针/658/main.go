package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	var leftArr []int
	var rightArr []int

	l, r := 0, len(arr)-1
	if arr[l] >= x {
		return arr[:k]
	}
	if arr[r] <= x {
		return arr[len(arr)-k:]
	}
	for l < r {
		mid := (l + r) / 2
		if arr[mid] <= x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	l--
	cnt := 0
	for cnt < k && l >= 0 && r < len(arr) {
		if x-arr[l] <= arr[r]-x {
			leftArr = append(leftArr, arr[l])
			l--
		} else {
			rightArr = append(rightArr, arr[r])
			r++
		}
		cnt++
	}
	for cnt < k && l >= 0 {
		leftArr = append(leftArr, arr[l])
		cnt++
		l--
	}
	for cnt < k && r < len(arr) {
		rightArr = append(rightArr, arr[r])
		cnt++
		r++
	}
	res := make([]int, k)
	for i := 0; i < len(leftArr); i++ {
		res[i] = leftArr[len(leftArr)-1-i]
	}
	for i := 0; i < len(rightArr); i++ {
		res[i+len(leftArr)] = rightArr[i]
	}
	return res
}

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 2))
}
