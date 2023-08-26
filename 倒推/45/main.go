package main

import "fmt"

func findPre(nums []int, pos int, step int) int {
	if pos == 0 {
		return step
	}
	for i := 0; i <= pos-1; i++ {
		if nums[i]+i >= pos {
			s := findPre(nums, i, step+1)
			if s > 0 {
				return s
			}
		}
	}
	return 0
}

func jump(nums []int) int {
	return findPre(nums, len(nums)-1, 0)
}

func main() {
	a := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(a))
}
