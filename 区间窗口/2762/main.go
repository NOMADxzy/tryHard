package main

import (
	"fmt"
	"math"
)

func continuousSubarrays(nums []int) int64 {
	ans := int64(0)
	left := 0
	cnt := make(map[int]int)

	maxVal := func(cnt map[int]int) int {
		max := 0
		for key := range cnt {
			if key > max {
				max = key
			}
		}
		return max
	}

	minVal := func(cnt map[int]int) int {
		min := math.MaxInt64 // 初始化为最大可能的整数值
		for key, _ := range cnt {
			if key < min {
				min = key
			}
		}
		return min
	}
	for right := 0; right < len(nums); right++ {
		cnt[nums[right]]++
		for maxVal(cnt)-minVal(cnt) > 2 {
			cnt[nums[left]]--
			if cnt[nums[left]] == 0 {
				delete(cnt, nums[left])
			}
			left++
		}
		ans += int64(right - left + 1)
	}

	return ans
}

func main() {
	fmt.Println(continuousSubarrays([]int{5, 4, 2, 4}))
}
