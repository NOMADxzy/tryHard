package main

import "fmt"

// 冒泡法(超时)
func findKthLargest1(nums []int, k int) int {
	for i := 0; i < k; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
	}
	return nums[len(nums)-k]
}

// 快排法
func findKthLargest(nums []int, k int) int {
	size := len(nums)
	return partSort(nums, 0, size-1, k)
}

func partSort(nums []int, l int, r int, k int) int {
	if l >= r {
		return nums[len(nums)-k]
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
	if len(nums)-i+1 > k {
		return partSort(nums, i+1, r, k)
	} else {
		return partSort(nums, l, i-1, k)
	}
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
