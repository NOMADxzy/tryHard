package main

import (
	"fmt"
	"sort"
)

// 动态规划
func largestNumber(cost []int, target int) string {
	dup := map[int]bool{}
	var nexts [][]int
	for i := 9; i > 0; i-- {
		if !dup[cost[i-1]] {
			nexts = append(nexts, []int{i, cost[i-1]})
			dup[cost[i-1]] = true
		}
	}
	sort.Slice(nexts, func(i, j int) bool {
		return nexts[i][1] < nexts[j][1]
	})

	insert := func(s string, v int) string {
		c := byte('0' + v)
		var i int
		for i = 0; i < len(s); i++ {
			if s[i] < c {
				break
			}
		}
		return s[:i] + string(c) + s[i:]
	}
	dp := make([]string, target+1)
	for i := nexts[0][1]; i <= target; i++ {
		best := dp[i]
		for _, next := range nexts {
			v, cos := next[0], next[1]
			if cos < i && len(dp[i-cos]) > 0 || cos == i {
				newS := insert(dp[i-cos], v)
				if len(newS) > len(best) || len(newS) == len(best) && newS > best {
					best = newS
				}
			}
		}
		dp[i] = best
	}
	if len(dp[target]) == 0 {
		return "0"
	}
	return dp[target]
}

func main() {
	fmt.Println(largestNumber([]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 9))
}

// 深度搜索，傻逼

func largestNumber1(cost []int, target int) string {
	dup := map[int]bool{}
	var nexts [][]int
	for i := 9; i > 0; i-- {
		if !dup[cost[i-1]] {
			nexts = append(nexts, []int{i, cost[i-1]})
			dup[cost[i-1]] = true
		}
	}
	sort.Slice(nexts, func(i, j int) bool {
		return nexts[i][1] < nexts[j][1]
	})
	var ans string
	ans_len := 0
	var dfs func(acc []byte, target int)
	done := false
	dfs = func(acc []byte, target int) {
		if done {
			return
		}
		if ans_len > 0 && target/nexts[0][1]+len(acc)+1 < ans_len {
			done = true
			return
		}
		if target == 0 {
			tmp := make([]byte, len(acc))
			copy(tmp, acc)
			sort.Slice(tmp, func(i, j int) bool {
				return tmp[i] > tmp[j]
			})
			if ans_len == 0 || len(tmp) == ans_len && string(tmp) > ans {
				ans = string(tmp)
				ans_len = len(ans)
			}
			return
		}
		for i := 0; i < len(nexts); i++ {
			v, c := nexts[i][0], nexts[i][1]
			if c <= target {
				dfs(append(acc, byte('0'+v)), target-c)
			} else {
				break
			}
		}
	}
	for len(nexts) > 0 {
		dfs([]byte{}, target)
		done = false
		nexts = nexts[1:]
	}
	if len(ans) == 0 {
		return "0"
	}
	return ans
}
