package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	var stack []int

	for i := 0; i < len(asteroids); i++ {
		if len(stack) == 0 {
			stack = append(stack, asteroids[i])
		} else {
			if asteroids[i] > 0 {
				stack = append(stack, asteroids[i])
			} else {
				for len(stack) > 0 && stack[len(stack)-1] <= -asteroids[i] {
					if stack[len(stack)-1] == -asteroids[i] {
						stack = stack[:len(stack)-1]
						break
					} else if stack[len(stack)-1] < 0 {
						stack = append(stack, asteroids[i])
						break
					} else {
						stack = stack[:len(stack)-1]
						if len(stack) == 0 {
							stack = append(stack, asteroids[i])
							break
						}
					}

				}
			}
		}
	}
	return stack
}

func main() {
	fmt.Println(asteroidCollision([]int{-8, 8}))
}
