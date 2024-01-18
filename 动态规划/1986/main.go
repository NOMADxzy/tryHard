package main

import (
	"fmt"
	"sort"
)

func minSessions(tasks []int, sessionTime int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] < tasks[j]
	})
	n := len(tasks)
	hist := map[int]int{}
	var findSubset func(marks int, limit int, acc, mask int, pos int, ans *[]int, set map[int]bool, preMin int)
	findSubset = func(marks int, limit int, acc, mask int, pos int, ans *[]int, set map[int]bool, preMin int) {
		if pos == n {
			if !set[acc] && preMin > limit {
				set[acc] = true
				*ans = append(*ans, acc)
			}
			return
		}
		if marks&mask == 0 {
			if limit < tasks[pos] {
				if !set[acc] && preMin > limit {
					set[acc] = true
					*ans = append(*ans, acc)
				}
				return
			} else {
				findSubset(marks+mask, limit-tasks[pos], acc+mask, mask*2, pos+1, ans, set, preMin)
				findSubset(marks, limit, acc, mask*2, pos+1, ans, set, min(preMin, tasks[pos]))
			}
		} else {
			findSubset(marks, limit, acc, mask*2, pos+1, ans, set, preMin)
		}
	}
	full := 1<<n - 1
	var dfs func(marks, step int) int
	dfs = func(marks int, step int) int {
		if v, ok := hist[marks]; ok {
			return v
		}
		if marks == full {
			return step
		}
		var subsets []int
		findSubset(marks, sessionTime, 0, 1, 0, &subsets, map[int]bool{}, 1<<31)
		sort.Slice(subsets, func(i, j int) bool {
			return subsets[i] > subsets[j]
		})
		best := n
		for i := 0; i < len(subsets); i++ {
			step1 := dfs(marks+subsets[i], step+1)
			best = min(best, step1)
		}
		hist[marks] = best
		return best
	}
	return dfs(0, 0)
}

// 3 3 2 2 2 2 1 1 1
func main() {
	fmt.Println(minSessions([]int{1, 3, 3, 1, 2, 2, 2, 2, 1}, 3))
}
