package main

import "fmt"

func xorAllNums(nums1 []int, nums2 []int) int {
	res1 := 0
	res2 := 0
	if len(nums1)%2 == 1 {
		for i := 0; i < len(nums2); i++ {
			res2 ^= nums2[i]
		}
	}
	if len(nums2)%2 == 1 {
		for i := 0; i < len(nums1); i++ {
			res1 ^= nums1[i]
		}
	}
	return res1 ^ res2
}

func main() {
	fmt.Println(xorAllNums([]int{2, 1, 3}, []int{10, 2, 5, 0}))
}
