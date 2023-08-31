package main

import "fmt"

func HeapSort(nums []int) []int {
	BuildHeap(nums)
	size := len(nums)
	for i := size - 1; i > 0; i-- {
		swap(nums, 0, i)
		check(nums, 0, i)
	}
	return nums
}

func BuildHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		check(nums, i, len(nums))
	}
}

func check(nums []int, i int, length int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < length && nums[left] > nums[largest] {
		largest = left
	}
	if right < length && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		swap(nums, largest, i)
		check(nums, largest, length)
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	arr := []int{33, 25, 46, 13, 58, 95, 18, 63}

	fmt.Println(HeapSort(arr))
}
