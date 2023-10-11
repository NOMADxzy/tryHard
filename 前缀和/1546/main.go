package main

import "fmt"

//贪心
func maxNonOverlapping(nums []int, target int) int {
	m := len(nums)
	sums := make([]int, m)
	sums[0] = nums[0]
	for i := 1; i < m; i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	ans := 0

	for i := 0; i < len(nums); {
		var newMap map[int]bool
		if i == 0 {
			newMap = map[int]bool{target: true}
		} else {
			newMap = map[int]bool{sums[i-1] + target: true}
		}
		var j int
		for j = i; j < len(sums); j++ {
			if newMap[sums[j]] {
				ans++
				break
			} else {
				newMap[sums[j]+target] = true
			}
		}
		i = j + 1
	}
	return ans
}

// 超时
func maxNonOverlapping1(nums []int, target int) int {
	m := len(nums)
	sums := make([]int, m)
	sums[0] = nums[0]
	for i := 1; i < m; i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	ans := 0
	lastIdx := -1

	for i := 0; i < m; i++ {
		for j := lastIdx + 1; j <= i; j++ {
			if j == 0 && sums[i] == target || j > 0 && sums[i]-sums[j-1] == target {
				lastIdx = i
				ans++
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxNonOverlapping([]int{1, 1, 1, 1, 1}, 2))
}
