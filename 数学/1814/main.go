package main

import "fmt"

func countNicePairs(nums []int) int {
	rev := func(n int) int {
		sum := 0
		for n > 0 {
			v := n % 10
			sum = sum*10 + v
			n /= 10
		}
		return sum
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] - rev(nums[i])
	}
	kMap := map[int]int{}
	MOD := 1000000007
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = (ans + kMap[nums[i]]) % MOD
		kMap[nums[i]]++
	}
	return ans
}

func main() {
	fmt.Println(countNicePairs([]int{42, 11, 1, 97}))
}
