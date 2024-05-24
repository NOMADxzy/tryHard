package main

import "fmt"

func singleNumber(nums []int) int {
	cnts := make([]int, 32)
	cnts_ := make([]int, 32)
	for _, num := range nums {
		mask := 1
		if num < 0 {
			for i := 0; mask <= -num; i++ {
				if (-num)&mask > 0 {
					cnts_[i]++
				}
				mask *= 2
			}
		} else {
			for i := 0; mask <= num; i++ {
				if num&mask > 0 {
					cnts[i]++
				}
				mask *= 2
			}
		}
	}
	ans := 0
	mask := 1
	for i := 0; i < 32; i++ {
		ans += (cnts[i] % 3) * mask
		mask *= 2
	}
	if ans == 0 {
		mask = 1
		for i := 0; i < 32; i++ {
			ans += (cnts_[i] % 3) * mask
			mask *= 2
		}
		return -ans
	}
	return ans
}

func main() {
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}
