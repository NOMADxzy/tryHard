package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	var stack []int
	var checkNext func(i, j int) bool
	checkNext = func(i, j int) bool {
		if j == n {
			return true
		}
		for i < n && (len(stack) == 0 || stack[len(stack)-1] != popped[j]) {
			stack = append(stack, pushed[i])
			i++
		}
		if i < n {
			tmp := make([]int, len(stack))
			copy(tmp, stack)
			stack = append(stack, pushed[i])
			if checkNext(i+1, j) {
				return true
			}
			stack = tmp
		}
		if stack[len(stack)-1] == popped[j] {
			tmp := make([]int, len(stack))
			copy(tmp, stack)
			stack = stack[:len(stack)-1]
			if checkNext(i, j+1) {
				return true
			}
			stack = tmp
		}
		return false
	}
	return checkNext(0, 0)
}
func main() {
	fmt.Println(validateStackSequences([]int{4, 5, 6, 5, 5, 7}, []int{4, 6, 7, 5, 5, 5}))
}
