package main

import "fmt"

func longestConsecutive(nums []int) int {
	exists := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		exists[nums[i]] = true
	}
	var dfs func(x int) int
	dfs = func(x int) int {
		if !exists[x] {
			return 0
		}
		exists[x] = false
		cnt := 1 + dfs(x-1) + dfs(x+1)
		return cnt
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = max(ans, dfs(nums[i]))
	}
	return ans
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
