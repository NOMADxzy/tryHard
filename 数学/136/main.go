package main

import "fmt"

func singleNumber(nums []int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		k = k ^ nums[i]
	}
	return k
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}
