package main

import "fmt"

func subsets(nums []int) [][]int {
	m := len(nums)
	var ans [][]int
	for i := 0; i < 1<<m; i++ {
		mask := 1
		var tmp []int
		for j := 0; mask <= i; j++ {
			if mask&i > 0 {
				tmp = append(tmp, nums[j])
			}
			mask *= 2
		}
		ans = append(ans, tmp)
	}
	return ans
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
