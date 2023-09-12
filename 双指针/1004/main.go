package main

import "fmt"

func longestOnes(nums []int, k int) int {
	m := len(nums)
	l, r := 0, 0
	count0 := 0
	max := 0
	for ; r < m; r++ {
		if nums[r] != 1 {
			count0++
			if count0 > k {
				for l < r && nums[l] == 1 {
					l++
				}
				l++
				count0--
			}
		}
		if r-l+1 > max {
			max = r - l + 1
		}
	}
	return max
}

func main() {
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}
