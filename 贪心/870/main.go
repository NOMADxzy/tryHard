package main

import (
	"fmt"
	"sort"
)

func buildHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		downCheck(arr, i, len(arr))
	}
}

func downCheck(arr []int, i int, size int) {
	l, r := 2*i+1, 2*i+2
	smallest := i
	if l < size && arr[l] < arr[smallest] {
		smallest = l
	}
	if r < size && arr[r] < arr[smallest] {
		smallest = r
	}
	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		downCheck(arr, smallest, size)
	}
}

func advantageCount(nums1 []int, nums2 []int) []int {
	buildHeap(nums1)
	num2Pair := make([][]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		num2Pair[i] = []int{nums2[i], i}
	}
	sort.Slice(num2Pair, func(i, j int) bool {
		return num2Pair[i][0] < num2Pair[j][0]
	})

	ans := make([]int, len(nums2))
	end := len(num2Pair) - 1
	length := len(nums1)

	for i := 0; length > 0; i++ {
		for length > 0 && nums1[0] <= num2Pair[i][0] {
			ans[num2Pair[end][1]] = nums1[0] //
			end--
			nums1[0], nums1[length-1] = nums1[length-1], nums1[0]
			length--
			downCheck(nums1, 0, length)
		}
		if length == 0 {
			break
		}
		ans[num2Pair[i][1]] = nums1[0]
		nums1[0], nums1[length-1] = nums1[length-1], nums1[0]
		length--
		downCheck(nums1, 0, length)
	}
	return ans
}

func main() {
	fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}
