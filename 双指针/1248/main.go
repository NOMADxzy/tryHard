package main

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	l, r := 0, 0
	cnt := 0
	ans := 0
	for r < len(nums) && cnt < k {
		if nums[r]%2 == 1 {
			cnt++
		}
		r++
	}
	if cnt < k {
		return 0
	}
	r--

	for r < len(nums) {
		nextL, nextR := l, r+1
		for nextR < len(nums) && nums[nextR]%2 != 1 {
			nextR++
		}
		for nums[nextL]%2 != 1 {
			nextL++
		}
		ans += (nextR - r) * (nextL - l + 1)

		l = nextL + 1
		r = nextR
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}
