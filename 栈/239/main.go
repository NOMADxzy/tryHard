package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	m := len(nums)
	var downStack []int
	ans := make([]int, m+1-k)
	idx := 0

	l, r := 0, 0
	for ; r < m; r++ {
		for len(downStack) > 0 && nums[downStack[len(downStack)-1]] <= nums[r] {
			downStack = downStack[:len(downStack)-1]
		}
		downStack = append(downStack, r)

		for r-l+1 > k {
			if downStack[0] <= l {
				downStack = downStack[1:]
			}
			l++
		}
		if r >= k-1 {
			ans[idx] = nums[downStack[0]]
			idx++
		}
	}
	return ans
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
