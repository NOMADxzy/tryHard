package main

import "fmt"

func tupleSameProduct(nums []int) int {
	vMap := map[int]int{}
	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cur := nums[i] * nums[j]
			ans += vMap[cur] * 8
			vMap[cur]++
		}
	}
	return ans
}

func main() {
	fmt.Println(tupleSameProduct([]int{1, 2, 4, 5, 10}))
}
