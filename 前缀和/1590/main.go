package main

import "fmt"

func minSubarray(nums []int, p int) int {
	m := len(nums)
	sums := make([]int, m)
	sums[0] = nums[0]
	total := nums[0]
	for i := 1; i < m; i++ {
		sums[i] = sums[i-1] + nums[i]
		total += nums[i]
	}

	target := total % p
	if target == 0 {
		return 0
	}
	ans := -1

	tMap := map[int]int{}
	bMap := map[int]bool{}
	bMap[target] = true
	tMap[target] = -1
	for i := 0; i < m; i++ {
		r := sums[i] % p
		if bMap[r] {
			if ans == -1 || i-tMap[r] < ans {
				ans = i - tMap[r]
			}

		}
		right := (r + target) % p
		bMap[right] = true
		tMap[right] = i
	}
	if ans == m {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minSubarray([]int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2}, 148))
}
