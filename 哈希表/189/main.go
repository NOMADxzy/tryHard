package main

import "fmt"

// TODO 空间消耗大
func rotate1(nums []int, k int) {
	numsMap := map[int]int{}

	for i := 0; i < len(nums); i++ {
		numsMap[i] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		nums[(i+k)%len(nums)] = numsMap[i]
	}
}

// 最优法，翻转数组
func rotate(nums []int, k int) {
	size := len(nums)
	for i := 0; i < size/2; i++ {
		tmp := nums[i]
		nums[i] = nums[size-i-1]
		nums[size-i-1] = tmp
	}
	k1 := k % size
	for i := 0; i < k1/2; i++ {
		tmp := nums[i]
		nums[i] = nums[k1-i-1]
		nums[k1-i-1] = tmp
	}
	for i := k1; i < k1+(size-k1)/2; i++ {
		tmp := nums[i]
		nums[i] = nums[k1+size-i-1]
		nums[k1+size-i-1] = tmp
	}
}

func main() {
	nums := []int{-1, -100, 3, 99}
	rotate(nums, 2)
	fmt.Println(nums)
}
