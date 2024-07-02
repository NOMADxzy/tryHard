package main

import "fmt"

func longestNiceSubarray(nums []int) int {
	ans := 0
	cnts := make([]int, 32)
	invalid_cnt := 0
	for l, r := 0, 0; r < len(nums); r++ {
		cur := nums[r]
		mask := 1
		for i := 0; mask <= cur; i++ {
			if cur&mask > 0 {
				cnts[i]++
				if cnts[i] == 2 {
					invalid_cnt++
				}
			}
			mask *= 2
		}
		for invalid_cnt > 0 {
			to_move := nums[l]
			mask = 1
			for i := 0; mask <= to_move; i++ {
				if to_move&mask > 0 {
					cnts[i]--
					if cnts[i] == 1 {
						invalid_cnt--
					}
				}
				mask *= 2
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

func main() {
	fmt.Println(longestNiceSubarray([]int{1, 3, 8, 48, 10}))
}
