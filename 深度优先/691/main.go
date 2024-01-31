package main

import (
	"fmt"
	"strconv"
)

func minStickers(stickers []string, target string) int {
	state := make([]int, 26)
	maxCnt := 0
	for i := 0; i < len(target); i++ {
		state[int(target[i]-'a')]++
		maxCnt = max(maxCnt, state[int(target[i]-'a')])
	}

	hist := map[string]int{}
	var dfs func(state []int, acc int) int
	dfs = func(state []int, acc int) int {
		if acc >= len(target) {
			return 0
		}
		key := ""
		for i := 0; i < 26; i++ {
			key += "_" + strconv.Itoa(state[i])
		}
		if hist[key] > 0 {
			return hist[key]
		}
		best := 1 << 31
		for i := 0; i < len(stickers); i++ {
			ok := 0
			word := stickers[i]
			preState := make([]int, len(state))
			copy(preState, state)
			for j := 0; j < len(word); j++ {
				p := int(word[j] - 'a')
				if state[p] > 0 {
					state[p]--
					ok++
				}
			}
			if ok > 0 {
				v := 1 + dfs(state, acc+ok)
				best = min(best, v)
			}
			copy(state, preState)
		}
		hist[key] = best
		return best
	}
	ans := dfs(state, 0)
	if ans == 1<<31 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minStickers([]string{"ab", "cd", "cd"}, "abcdcdcdcd"))
}
