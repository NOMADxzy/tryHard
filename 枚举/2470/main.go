package main

import "fmt"

func subarrayLCM(nums []int, k int) int {
	ans := 0

	getCommonMulVal := func(a, b int) int {
		tmp := a * b
		if a < b {
			a, b = b, a
		}
		for b != 0 {
			a, b = b, a%b
		}
		return tmp / a
	}

	for i := 0; i < len(nums); i++ {
		pre := 1
		for j := i; j >= 0 && pre <= k; j-- {
			pre = getCommonMulVal(nums[j], pre)
			if pre == k {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(subarrayLCM([]int{3, 6, 2, 7, 1}, 6))
}
