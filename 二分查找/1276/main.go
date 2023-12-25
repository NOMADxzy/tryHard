package main

import "fmt"

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	Small, Big := 2, 4
	left := 0
	right := cheeseSlices
	if right*Small > tomatoSlices || right*Big < tomatoSlices {
		return []int{}
	}
	for left < right {
		mid := (left + right) / 2
		if mid*Small+(cheeseSlices-mid)*Big == tomatoSlices {
			return []int{cheeseSlices - mid, mid}
		} else if mid*Small+(cheeseSlices-mid)*Big < tomatoSlices {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left*Small+(cheeseSlices-left)*Big == tomatoSlices {
		return []int{cheeseSlices - left, left}
	}
	return []int{}
}

func main() {
	fmt.Println(numOfBurgers(2537427, 860448))
}
