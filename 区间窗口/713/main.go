package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	l, r := 0, 0
	tmp := 1
	var idxArr [][]int
	cnt := 0
	for r < len(nums) {
		tmp *= nums[r]
		for tmp >= k && l <= r {
			tmp /= nums[l]
			l++
		}
		idxArr = append(idxArr, []int{l, r})
		if r >= l {
			cnt += r - l + 1
		}
		r++
	}

	return cnt
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 0))
}
