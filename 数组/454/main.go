package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)
	vMap := map[int]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			vMap[nums1[i]+nums2[j]]++
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := nums3[i] + nums4[j]
			if vMap[-val] > 0 {
				ans += vMap[-val]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(fourSumCount([]int{1}, []int{-1}, []int{0}, []int{1}))
}
