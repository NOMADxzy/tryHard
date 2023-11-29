package main

import "fmt"

func largestRectangleArea(heights []int) int {
	m := len(heights)
	var incStack []int
	nextSmaller := make([]int, m)
	for i := 0; i < m; i++ {
		for len(incStack) > 0 && heights[incStack[len(incStack)-1]] > heights[i] {
			idx := incStack[len(incStack)-1]
			incStack = incStack[:len(incStack)-1]
			nextSmaller[idx] = i
		}
		incStack = append(incStack, i)
		nextSmaller[i] = -1
	}
	preSmaller := make([]int, m)
	for i := m - 1; i >= 0; i-- {
		for len(incStack) > 0 && heights[incStack[len(incStack)-1]] > heights[i] {
			idx := incStack[len(incStack)-1]
			incStack = incStack[:len(incStack)-1]
			preSmaller[idx] = i
		}
		incStack = append(incStack, i)
		preSmaller[i] = -1
	}

	ans := 0
	for i := 0; i < m; i++ {
		lb, rb := 0, m-1
		if nextSmaller[i] >= 0 {
			rb = nextSmaller[i] - 1
		}
		if preSmaller[i] >= 0 {
			lb = preSmaller[i] + 1
		}
		if val := (rb - lb + 1) * heights[i]; val > ans {
			ans = val
		}
	}
	return ans
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 4}))
}
