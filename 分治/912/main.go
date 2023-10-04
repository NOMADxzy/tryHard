package main

import "fmt"

func merge(arr1, arr2 []int) []int {
	i, j := 0, 0
	newArr := make([]int, len(arr1)+len(arr2))

	pos := 0
	for ; i < len(arr1) && j < len(arr2); pos++ {
		if arr1[i] < arr2[j] {
			newArr[pos] = arr1[i]
			i++
		} else {
			newArr[pos] = arr2[j]
			j++
		}
	}
	for ; i < len(arr1); i++ {
		newArr[pos] = arr1[i]
		pos++
	}
	for ; j < len(arr2); j++ {
		newArr[pos] = arr2[j]
		pos++
	}
	return newArr
}

func partitionSort(arr []int) []int {
	if len(arr) > 1 {
		mid := len(arr) / 2
		arr1 := partitionSort(arr[:mid])
		arr2 := partitionSort(arr[mid:])
		return merge(arr1, arr2)
	} else {
		return arr
	}
}

func sortArray(nums []int) []int {
	return partitionSort(nums)
}

func main() {
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}
