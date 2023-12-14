package main

import "fmt"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets == 1 {
		return 0
	}
	k := minutesToTest / minutesToDie
	dim := k + 1
	wid := 1
	sum := dim
	for sum < buckets {
		wid++
		sum = dim * sum
	}
	return wid
}

func main() {
	fmt.Println(poorPigs(1000, 15, 60))
}
