package main

import "fmt"

func maxLength(arr []string) int {
	masks := make([]int, len(arr))
	valid := make([]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		mask := 0
		valid[i] = true
		for j := 0; j < len(arr[i]); j++ {
			if mask&(1<<(arr[i][j]-'a')) > 0 {
				valid[i] = false
				break
			}
			mask |= 1 << (arr[i][j] - 'a')
		}
		masks[i] = mask

	}
	hist := make([]map[int]int, len(arr))
	for i := 0; i < len(hist); i++ {
		hist[i] = map[int]int{}
	}
	var dfs func(pos int, state int) int
	dfs = func(pos int, state int) int {
		if pos == len(arr) {
			return 0
		}
		if v, ok := hist[pos][state]; ok {
			return v
		}
		ret := dfs(pos+1, state)

		if valid[pos] && state&masks[pos] == 0 {
			ret1 := dfs(pos+1, state|masks[pos]) + len(arr[pos])
			if ret1 > ret {
				ret = ret1
			}
		}
		hist[pos][state] = ret
		return ret
	}
	return dfs(0, 0)
}

func main() {
	fmt.Println(maxLength([]string{"cha", "r", "act", "ers"}))
}
