package main

import "fmt"

func swapNumbers(numbers []int) []int {
	if numbers[0] == 0 {
		numbers[0] = numbers[1]
		numbers[1] = 0
	} else if numbers[1] == 0 {
		numbers[1] = numbers[0]
		numbers[0] = 0
	} else {
		numbers[1] *= numbers[0]
		numbers[0] = numbers[1] / numbers[0]
		numbers[1] = numbers[1] / numbers[0]
	}
	return numbers
}

func main() {
	fmt.Println(swapNumbers([]int{-1904, 2493}))
}
