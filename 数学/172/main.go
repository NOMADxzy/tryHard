package main

import "fmt"

func trailingZeroes(n int) int {

	count2 := 0
	count5 := 0

	for i := 2; i <= n; i++ {
		num := i
		for num%5 == 0 {
			count5++
			num /= 5
		}
		for num%2 == 0 {
			count2++
			num /= 2
		}

	}
	if count5 < count2 {
		return count5
	}
	return count2
}

func main() {
	fmt.Println(trailingZeroes(20))
}
