package main

import "fmt"

func restoreIpAddresses(s string) []string {
	for i := 0; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return []string{}
		}
	}

	getIntVal := func(s string) int {
		if len(s) == 0 || len(s) > 1 && s[0] == '0' || len(s) > 3 {
			return -1
		}
		val := 0
		for i := 0; i < len(s); i++ {
			val = val*10 + int(s[i]-'0')
		}
		if val > 255 {
			return -1
		}
		return val
	}
	resultsCache := map[int][]string{}

	var dfs func(pos int, i int) []string
	dfs = func(pos int, i int) []string {
		key := i*4 + pos
		if v, ok := resultsCache[key]; ok {
			return v
		}
		if pos == 3 {
			lastVal := getIntVal(s[i:])
			if lastVal >= 0 {
				return []string{s[i:]}
			} else {
				return []string{}
			}
		}
		var res []string
		for j := i + 1; j <= len(s) && j-i <= 3; j++ {
			curVal := getIntVal(s[i:j])
			if curVal >= 0 {
				nextStrs := dfs(pos+1, j)
				for k := 0; k < len(nextStrs); k++ {
					res = append(res, s[i:j]+"."+nextStrs[k])
				}
			}
		}
		resultsCache[key] = res
		return res
	}
	return dfs(0, 0)
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
