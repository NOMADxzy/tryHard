package main

import (
	"fmt"
	"sort"
)

func minimumAddedInteger(nums1 []int, nums2 []int) int {
	nums1, nums2 = nums2, nums1
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	for firstPlace := 2; firstPlace >= 0; firstPlace-- {
		delta := nums2[firstPlace] - nums1[0]
		accWrong := 0
		var i int
		for i = 1; i < len(nums1) && accWrong+firstPlace < 3; i++ {
			if nums1[i]+delta != nums2[i+firstPlace+accWrong] {
				accWrong++
				i--
			}
		}
		if i == len(nums1) {
			return -delta
		}
	}
	return -1
}

func main() {
	fmt.Println(minimumAddedInteger([]int{4, 20, 16, 12, 8}, []int{14, 18, 10}))
}
