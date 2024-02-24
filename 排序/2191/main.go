package main

import (
	"fmt"
	"sort"
)

func sortJumbled(mapping []int, nums []int) []int {
	hist := map[int]int{}
	trueValue := func(x int) int {
		if v, ok := hist[x]; ok {
			return v
		}
		val := 0
		pow := 1
		if x == 0 {
			hist[x] = mapping[0]
			return hist[x]
		}
		for tmp := x; tmp > 0; tmp /= 10 {
			val += mapping[tmp%10] * pow
			pow *= 10
		}
		hist[x] = val
		return val
	}
	// 使用 sort.Slice ... trueValue(nums[i]) < trueValue(nums[j]) || trueValue(nums[i]) == trueValue(nums[j]) && i<j 不行
	sort.SliceStable(nums, func(i, j int) bool { // 使用稳定排序 TODO
		return trueValue(nums[i]) < trueValue(nums[j])
	})
	return nums
}

func main() {
	fmt.Println(sortJumbled([]int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38, 0}))
}
