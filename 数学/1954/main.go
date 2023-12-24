package main

import "fmt"

func minimumPerimeter(neededApples int64) int64 {
	acc := func(i int64) int64 {
		return (i * (i + 1) * (2*i + 1) / 6) * 3
	}
	left, right := int64(0), int64(100000)
	for left < right {
		mid := (left + right) / 2
		if acc(mid) < neededApples {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return 4 * 2 * left
}

func main() {
	fmt.Println(minimumPerimeter(100000000000000))
}
