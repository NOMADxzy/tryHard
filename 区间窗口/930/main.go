package main

import "fmt"

func numSubarraysWithSum(nums []int, goal int) int {
	l, r := 0, 0
	cnt := 0
	sum := 0

	for r < len(nums) {
		sum += nums[r]
		for sum > goal && l < len(nums) {
			sum -= nums[l]
			l++
		}
		if l == len(nums) {
			break
		}
		if sum == goal && r >= l { //goal==0时，不加r>=l会多算空数组
			cnt++
			for j := l - 1; j >= 0 && nums[j] == 0; j-- {
				cnt++
			}
			for l < r && nums[l] == 0 {
				cnt++
				l++

			}
		}
		r++
	}
	return cnt
}

func main() {
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}
