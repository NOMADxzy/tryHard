package main

import "fmt"

func maxWidthRamp(nums []int) int {
	n := len(nums)
	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}
	best := 0
	for i := 0; i < n-1-best; i++ {
		if rightMax[n-1] >= nums[i] {
			best = max(best, n-1-i)
		} else {
			left, right := i, n-1
			for left < right {
				mid := (left + right) / 2
				if rightMax[mid] < nums[i] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			best = max(best, right-1-i)
		}
	}
	return best
}

func main() {
	fmt.Println(maxWidthRamp([]int{2, 2, 1}))
}
