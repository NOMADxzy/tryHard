package main

import "fmt"

func minimumSubarrayLength(nums []int, k int) int {
	cnts := make([]int, 32)
	state := 0
	ans := len(nums) + 1
	for l, r := 0, 0; r < len(nums); r++ {
		cur := nums[r]
		mask := 1
		for i := 0; mask <= cur; i++ {
			if mask&cur > 0 {
				cnts[i]++
				if cnts[i] == 1 {
					state += mask
				}
			}
			mask *= 2
		}
		for state >= k && l <= r {
			ans = min(r-l+1, ans)
			remove := nums[l]
			mask = 1
			for i := 0; mask <= remove; i++ {
				if mask&remove > 0 {
					cnts[i]--
					if cnts[i] == 0 {
						state -= mask
					}
				}
				mask *= 2
			}
			l++
		}
	}
	if ans == len(nums)+1 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minimumSubarrayLength([]int{1, 2}, 0))
}
