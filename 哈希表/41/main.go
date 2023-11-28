package main

import "fmt"

func firstMissingPositive(nums []int) int {
	m := len(nums)
	for i := 0; i < m; i++ {
		targetPlace := nums[i] - 1
		if targetPlace >= 0 && targetPlace < m && nums[targetPlace] != targetPlace+1 {
			//nums[targetPlace] = targetPlace + 1
			for nums[targetPlace] > 0 && nums[targetPlace] <= m && nums[nums[targetPlace]-1] != nums[targetPlace] {
				tmp := nums[targetPlace] - 1
				nums[targetPlace] = targetPlace + 1
				targetPlace = tmp
			}
			nums[targetPlace] = targetPlace + 1
		}
	}
	ans := m + 1
	for i := 0; i < m; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}
