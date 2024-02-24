package main

import "fmt"

func distributeCookies(cookies []int, k int) int {
	hist := map[int]int{}
	n := len(cookies)
	FULL := 1<<n - 1
	var queryBest func(state int, space int) int
	var findAll func(state int, i, mask int, combs map[int]bool)
	findAll = func(state int, i, mask int, combs map[int]bool) {
		if i == n {
			return
		} else if state&mask == 0 {
			combs[state+mask] = true
			findAll(state+mask, i+1, mask*2, combs)
		}
		combs[state] = true
		findAll(state, i+1, mask*2, combs)
	}
	sum := map[int]int{}
	for i := 0; i < 1<<n; i++ {
		v := 0
		mask := 1
		for j := 0; j < n; j++ {
			if i&mask > 0 {
				v += cookies[j]
			}
			mask *= 2
		}
		sum[i] = v
	}
	queryBest = func(state int, space int) int {
		key := state*n + space
		if hist[key] > 0 {
			return hist[key]
		}
		if state == FULL {
			return 0
		}
		if space == 1 {
			ans := 0
			mask := 1
			for i := 0; i < n; i++ {
				if state&mask == 0 {
					ans += cookies[i]
				}
				mask *= 2
			}
			hist[key] = ans
			return ans
		} else {
			ans := 1 << 31
			combs := map[int]bool{}
			findAll(state, 0, 1, combs)
			for st, _ := range combs {
				if st == state {
					continue
				}
				cur := sum[st] - sum[state]
				v := max(cur, queryBest(st, space-1))
				ans = min(ans, v)
			}
			hist[key] = ans
			return ans
		}
	}
	return queryBest(0, k)
}

func main() {
	fmt.Println(distributeCookies([]int{2, 3, 2, 1, 1}, 2))
}

// 1 0 1 1 0 1 0
