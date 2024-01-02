package main

import "fmt"

func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, 1
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > right {
			right = nums[i]
		}
	}
	getSum := func(div int) int {
		sum := 0
		for i := 0; i < n; i++ {
			sum += nums[i] / div
			if nums[i]%div != 0 {
				sum++
			}
		}
		return sum
	}
	for left < right {
		mid := (left + right) / 2
		if getSum(mid) <= threshold {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	fmt.Println(smallestDivisor([]int{1, 2, 5, 9}, 6))
}
