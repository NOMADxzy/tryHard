package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	NGidx := make([]int, len(temperatures))
	var stack [][]int

	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 {
			stack = append(stack, []int{temperatures[i], i})
		} else {
			for len(stack) > 0 && stack[len(stack)-1][0] < temperatures[i] {
				idx := stack[len(stack)-1][1]
				stack = stack[:len(stack)-1]
				NGidx[idx] = i
			}
			stack = append(stack, []int{temperatures[i], i})
		}
	}

	for i := 0; i < len(NGidx); i++ {
		if NGidx[i] > 0 {
			NGidx[i] -= i
		}
	}
	return NGidx
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
