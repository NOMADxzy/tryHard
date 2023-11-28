package main

import "fmt"

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	m := len(height)
	rightMax := make([]int, m)
	rightMax[m-1] = height[m-1]
	for i := m - 2; i >= 0; i-- {
		rightMax[i] = rightMax[i+1]
		if height[i] > rightMax[i] {
			rightMax[i] = height[i]
		}
	}
	leftMax := height[0]
	ans := 0
	for i := 1; i < m-1; i++ {
		if val := min(leftMax, rightMax[i+1]) - height[i]; val > 0 {
			ans += val
		}
		if height[i] > leftMax {
			leftMax = height[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
