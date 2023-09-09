package main

import "fmt"

func climbStairs(n int) int {
	a := 1
	b := 1
	var tmp int

	for i := 2; i <= n; i++ {
		tmp = b
		b = a + b
		a = tmp
	}
	return b
}

func main() {
	fmt.Println(climbStairs(30))
}
