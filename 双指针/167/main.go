package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		if numbers[l]+numbers[r] > target {
			r--
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			break
		}
	}
	return []int{l + 1, r + 1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
