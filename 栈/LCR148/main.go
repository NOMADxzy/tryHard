package main

import "fmt"

func validateBookSequences(putIn []int, takeOut []int) bool {
	if len(putIn) != len(takeOut) {
		return false
	}
	if len(putIn) == 0 {
		return true
	}
	cur := 0
	var stack []int
	stack = append(stack, putIn[cur])
	for i := 0; i < len(takeOut); i++ {
		need := takeOut[i]
		if len(stack) > 0 && stack[len(stack)-1] == need {
			stack = stack[:len(stack)-1]
			continue
		} else {
			for cur+1 < len(putIn) && putIn[cur+1] != need {
				cur++
				stack = append(stack, putIn[cur])
			}
			if cur+1 == len(putIn) {
				return false
			} else {
				cur++
			}
		}
	}
	return true
}

func main() {
	fmt.Println(validateBookSequences([]int{6, 7, 8, 9, 10, 11}, []int{9, 11, 10, 8, 7, 6}))
}
