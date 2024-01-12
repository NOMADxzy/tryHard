package main

import "fmt"

// 1 2 4 8 16
func minPatches(nums []int, n int) int {
	preMax := 0
	ans := 0
	for i := 0; i < len(nums) && preMax < n; i++ {
		cur := nums[i]
		if cur > preMax+1 { //出现了中断
			for preMax < n && cur > preMax+1 {
				preMax = (preMax + 1) + preMax //补充了preMax + 1 这个数
				ans++
			}
		}
		preMax = preMax + cur
	}
	for preMax < n {
		preMax = (preMax + 1) + preMax //补充了preMax + 1 这个数
		ans++
	}
	return ans
}

func main() {
	fmt.Println(minPatches([]int{1, 2, 31, 33}, 2147483647))
}
