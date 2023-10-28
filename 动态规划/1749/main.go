package main

import "fmt"

func maxAbsoluteSum(nums []int) int {
	m := len(nums)
	pre := nums[0]
	best := nums[0]
	if best < 0 {
		best = -best
	}

	for i := 1; i < m; i++ {
		now := nums[i]
		if pre > 0 {
			now += pre
		}
		if now > best {
			best = now
		}
		pre = now
	}
	pre = nums[0]
	for i := 1; i < m; i++ {
		now := nums[i]
		if pre < 0 {
			now += pre
		}
		if -now > best {
			best = -now
		}
		pre = now
	}
	return best
}

func main() {
	fmt.Println(maxAbsoluteSum([]int{2, -5, 1, -4, 3, -2}))
}
