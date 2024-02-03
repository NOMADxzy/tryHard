package main

import "fmt"

func numTriplets(nums1 []int, nums2 []int) int {
	map1s := map[int]int{}
	map1d := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		map1s[nums1[i]*nums1[i]]++
		for j := i + 1; j < len(nums1); j++ {
			map1d[nums1[i]*nums1[j]]++
		}
	}
	ans := 0
	for i := 0; i < len(nums2); i++ {
		v := nums2[i] * nums2[i]
		ans += map1d[v]
		for j := i + 1; j < len(nums2); j++ {
			ans += map1s[nums2[i]*nums2[j]]
		}
	}
	return ans
}

func main() {
	fmt.Println(numTriplets([]int{1, 1}, []int{1, 1, 1}))
}
