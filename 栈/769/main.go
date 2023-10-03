package main

import "fmt"

//2,1,0,3,4
//分析法
func maxChunksToSorted(arr []int) int {
	maxX := 0
	ans := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxX {
			maxX = arr[i]
		}
		if i == maxX {
			ans++
		}
	}
	return ans
}

//单调栈法
func maxChunksToSorted1(arr []int) int {
	var stack []int

	for i := 0; i < len(arr); i++ {
		if len(stack) == 0 {
			stack = append(stack, arr[i])
		} else {
			maxVal := arr[i]
			if stack[len(stack)-1] > arr[i] {
				maxVal = stack[len(stack)-1]
			}
			for len(stack) > 0 && stack[len(stack)-1] > arr[i] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, maxVal)
		}
	}
	return len(stack)
}

func main() {
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
}
