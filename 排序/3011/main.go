package main

import "fmt"

func canSortArray(nums []int) bool {
	get1cnt := func(x int) int {
		cnt := 0
		for x > 0 {
			if x%2 == 1 {
				cnt++
			}
			x /= 2
		}
		return cnt
	}
	preMax := -1 << 31
	for i := 0; i < len(nums); {
		curMax := nums[i]
		curMin := nums[i]
		cnt := get1cnt(nums[i])
		j := i
		for j+1 < len(nums) && get1cnt(nums[j+1]) == cnt {
			j++
			curMin = min(curMin, nums[j])
			curMax = max(curMax, nums[j])
		}
		if curMin < preMax {
			return false
		}
		preMax = curMax
		i = j + 1
	}
	return true
}

func main() {
	fmt.Println(canSortArray([]int{8, 4, 2, 30, 15}))
}
