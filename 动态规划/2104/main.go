package main

import "fmt"

func subArrayRanges(nums []int) int64 {
	premins := []int{nums[0]}
	premaxs := []int{nums[0]}
	var ans int64
	for i := 1; i < len(nums); i++ {
		var premins1, premaxs1 []int
		premins1 = make([]int, i+1)
		premaxs1 = make([]int, i+1)
		premins1[i] = nums[i]
		premaxs1[i] = nums[i]
		for j := 0; j < len(premins); j++ {
			if nums[i] < premins[j] {
				premins1[j] = nums[i]
			} else {
				premins1[j] = premins[j]
			}
			if nums[i] > premaxs[j] {
				premaxs1[j] = nums[i]
			} else {
				premaxs1[j] = premaxs[j]
			}
			ans += int64(premaxs1[j] - premins1[j])
		}
		premins = premins1
		premaxs = premaxs1
	}
	return ans
}

// 优等解：单调栈，求每个元素为极值的子数组数量，ans = sum(最大值) - sum(最小值)

func main() {
	fmt.Println(subArrayRanges([]int{1, 2, 3}))
}
