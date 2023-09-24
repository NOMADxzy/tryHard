package main

import "fmt"

func singleNumber(nums []int) []int {
	x := 0
	for i := 0; i < len(nums); i++ {
		x = x ^ nums[i]
	}
	mask := x & (-x)

	class1, class2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&mask == 0 {
			class1 = class1 ^ nums[i]
		} else {
			class2 = class2 ^ nums[i]
		}
	}
	return []int{class2, class1}
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}
