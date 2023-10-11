package main

import "fmt"

func minOperations(nums []int, x int) int {
	m := len(nums)
	sums := 0
	for i := 0; i < m; i++ {
		sums += nums[i]

	}
	target := sums - x
	if target == 0 { //TODO 容易漏掉空序列
		return m
	}

	sum := 0
	ans := -1
	tMap := map[int]int{target: -1}
	for i := 0; i < m; i++ {
		sum += nums[i]
		if v, ok := tMap[sum]; ok {
			if i-v > ans {
				ans = i - v
			}
		}
		tMap[sum+target] = i
	}
	if ans == -1 {
		return -1
	}
	return m - ans
}

func main() {
	fmt.Println(minOperations([]int{8828, 9581, 49, 9818, 9974, 9869, 9991, 10000, 10000, 10000, 9999, 9993, 9904, 8819, 1231, 6309}, 134365))
}
