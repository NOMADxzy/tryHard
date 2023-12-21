package main

import "fmt"

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	sum := 0
	for i := 1; ; i++ {
		sum += i
		if (sum-target) >= 0 && (sum-target)%2 == 0 {
			return i
		}
	}
}

func main() {
	fmt.Println(reachNumber(2))
}
