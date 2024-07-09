package main

import "fmt"

func specialPerm(nums []int) int {
	m := len(nums)
	hist := map[int]int{}
	MOD := 1000000007
	var dfs func(pos int, state int, lastChooseIdx int) int
	dfs = func(pos int, state int, lastChooseIdx int) int {
		key := lastChooseIdx<<m + state
		if v, ok := hist[key]; ok {
			return v
		}
		if pos == m {
			return 1
		}
		cnt := 0
		mask := 1
		for i := 0; i < m; i++ {
			if state&mask == 0 {
				if lastChooseIdx == -1 || nums[lastChooseIdx]%nums[i] == 0 || nums[i]%nums[lastChooseIdx] == 0 {
					cnt = (cnt + dfs(pos+1, state|mask, i)) % MOD
				}
			}
			mask *= 2
		}
		hist[key] = cnt
		return cnt
	}
	return dfs(0, 0, -1)
}

func main() {
	fmt.Println(specialPerm([]int{2, 3, 6}))
}
