package main

import "fmt"

func partSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	val := nums[l]

	for i < j {
		for ; i < j && nums[j] >= val; j-- {
		}
		nums[i], nums[j] = nums[j], nums[i]

		for ; i < j && nums[i] <= val; i++ {
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	partSort(nums, l, i-1)
	partSort(nums, i+1, r)
}

func quickSort(nums []int) {
	partSort(nums, 0, len(nums)-1)
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	quickSort(nums)
	fmt.Println(nums)
}
