package main

import "fmt"

func majorityElement(nums []int) []int {
	times := map[int]int{}
	thres := len(nums)/3 + 1

	var res []int

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		times[val] += 1
		if times[val] == thres {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	fmt.Println(majorityElement([]int{1, 2}))
}
