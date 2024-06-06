package main

import "fmt"

func findMaxLength(nums []int) int {
	hist := map[int]int{0: -1}
	cur := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cur--
		} else {
			cur++
		}
		if v, ok := hist[cur]; ok {
			ans = max(ans, i-v)
		} else {
			hist[cur] = i
		}
	}
	return ans
}

func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
}
