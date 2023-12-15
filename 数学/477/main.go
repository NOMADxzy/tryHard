package main

import "fmt"

func totalHammingDistance(nums []int) int {
	m := len(nums)
	hammingCntOfDim := make([]int, 32)
	for i := 0; i < m; i++ {
		num := nums[i]
		mask := 1
		for j := 0; mask <= num; j++ {
			if num&mask > 0 {
				hammingCntOfDim[j]++
			}
			mask *= 2
		}
	}
	ans := 0
	for i := 0; i < 32; i++ {
		cnt1 := hammingCntOfDim[i]
		if cnt1 > 0 && cnt1 < m {
			ans += cnt1 * (m - cnt1)
		}
	}
	return ans
}

func main() {
	fmt.Println(totalHammingDistance([]int{4}))
}
