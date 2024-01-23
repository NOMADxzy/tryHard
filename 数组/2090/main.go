package main

import "fmt"

func getAverages(nums []int, k int) []int {
	cnt := 2*k + 1
	m := len(nums)
	ans := make([]int, m)
	tmpSum := 0
	for i := 0; i < m; i++ {
		if i-k < 0 || i+k >= m {
			ans[i] = -1
		} else if i-k == 0 {
			for j := i - k; j <= i+k; j++ {
				tmpSum += nums[j]
			}
			ans[i] = tmpSum / cnt
		} else {
			tmpSum = tmpSum + nums[i+k] - nums[i-k-1]
			ans[i] = tmpSum / cnt
		}
	}
	return ans
}

func main() {
	fmt.Println(getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
}
