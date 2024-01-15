package main

import "fmt"

func maxDistance(nums1 []int, nums2 []int) int {
	p1, p2 := 0, 0
	ans := 0
	for p2 < len(nums2) {
		for p2 < len(nums2) && nums1[p1] <= nums2[p2] {
			p2++
		}
		ans = max(0, p2-p1-1)
		for p2 < len(nums2) && p1 < len(nums1) && nums1[p1] > nums2[p2] {
			p1++
			p2++
		}
	}
	return ans
}

func main() {
	fmt.Println(maxDistance([]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}))
}
