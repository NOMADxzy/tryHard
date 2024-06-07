package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	var stack []int
	stack = append(stack, asteroids[0])
	for i := 1; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			stack = append(stack, asteroids[i])
		} else {
			acc := asteroids[i]
			for len(stack) > 0 && stack[len(stack)-1] > 0 {
				if acc+stack[len(stack)-1] == 0 {
					stack = stack[:len(stack)-1]
					acc = 0
					break
				} else if acc+stack[len(stack)-1] < 0 {
					stack = stack[:len(stack)-1]
				} else {
					acc = 0
					break
				}
			}
			if acc != 0 {
				stack = append(stack, acc)
			}
		}
	}
	return stack
}

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
}
