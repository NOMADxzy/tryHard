package main

import (
	"fmt"
	"strconv"
)

func handle(a int, b int) string {
	if a == b {
		return strconv.Itoa(a)
	} else {
		return strconv.Itoa(a) + "->" + strconv.Itoa(b)
	}
}

func summaryRanges(nums []int) []string {
	var ss []string
	left, right := 0, 0
	m := len(nums)

	for right < len(nums) {
		for right < m-1 && nums[right]+1 == nums[right+1] {
			right++
		}
		ss = append(ss, handle(nums[left], nums[right]))
		right++
		left = right
	}
	return ss
}

func main() {
	fmt.Println(summaryRanges([]int{}))
}
