package main

import "fmt"

func grayCode(n int) []int {
	var ans []int
	ans = append(ans, 0)
	mask := 1
	for i := 0; i < n; i++ {
		ans_ := make([]int, len(ans))
		for j := 0; j < len(ans_); j++ {
			ans_[j] = mask + ans[j]
		}
		for j := len(ans_) - 1; j >= 0; j-- {
			ans = append(ans, ans_[j])
		}
		mask <<= 1
	}
	return ans
}

func main() {
	fmt.Println(grayCode(16))
}
