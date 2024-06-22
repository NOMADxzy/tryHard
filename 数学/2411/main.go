package main

import "fmt"

func smallestSubarrays(nums []int) []int {
	m := len(nums)
	var cnts [32][]int
	for i := 0; i < 32; i++ {
		cnts[i] = make([]int, m)
	}
	mask := 1
	for i := 0; i < 32; i++ {
		acc := 0
		for j := 0; j < m; j++ {
			cur := nums[j] & mask
			if cur > 0 {
				acc++
			}
			cnts[i][j] = acc
		}
		mask *= 2
	}
	ans := make([]int, m)
	for j := 0; j < m; j++ {
		maxLen := 1
		mask = 1
		for i := 0; i < 32; i++ {
			cur := mask & nums[j]
			if cur == 0 && cnts[i][m-1] > cnts[i][j] { // 这一位可以补
				left, right := j, m-1
				for left < right {
					mid := (left + right) / 2
					if cnts[i][mid] > cnts[i][j] {
						right = mid
					} else {
						left = mid + 1
					}
				}
				curLen := right - j + 1
				if curLen > maxLen {
					maxLen = curLen
				}
			}
			mask *= 2
		}
		ans[j] = maxLen
	}
	return ans
}

func main() {
	fmt.Println(smallestSubarrays([]int{1, 2, 4, 8}))
}
