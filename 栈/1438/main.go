package main

import "fmt"

func longestSubarray(nums []int, limit int) int {
	m := len(nums)

	var upStack, downStack []int

	ans := 0
	l, r := 0, 0
	for ; r < m; r++ {
		for len(upStack) > 0 && nums[upStack[len(upStack)-1]] >= nums[r] {
			upStack = upStack[:len(upStack)-1]
		}
		upStack = append(upStack, r)
		for len(downStack) > 0 && nums[downStack[len(downStack)-1]] <= nums[r] {
			downStack = downStack[:len(downStack)-1]
		}
		downStack = append(downStack, r)

		//maxN, minN := findMaxMin()
		for nums[downStack[0]]-nums[upStack[0]] > limit && r-l+1 > ans {
			if downStack[0] == l {
				downStack = downStack[1:]
			}
			if upStack[0] == l {
				upStack = upStack[1:]
			}
			l++
		}
		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(longestSubarray([]int{8, 2, 4, 7}, 4))
}
