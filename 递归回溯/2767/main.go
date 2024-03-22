package main

import "fmt"

func minimumBeautifulSubstrings(s string) int {
	addStr := func(s1, s2 string) string {
		if len(s1) < len(s2) {
			s1, s2 = s2, s1
		}
		c := 0
		res := make([]byte, len(s1))
		for i := 0; i < len(s1); i++ {
			a := int(s1[len(s1)-1-i] - '0')

			b := 0
			if len(s2)-1-i >= 0 {
				b = int(s2[len(s2)-1-i] - '0')
			}

			val := a + b + c
			if val >= 2 {
				val -= 2
				c = 1
			} else {
				c = 0
			}
			if val == 0 {
				res[len(res)-1-i] = '0' // TODO 易错 注意索引是从右往左 进行的
			} else {
				res[len(res)-1-i] = '1'
			}
		}
		resStr := string(res)
		if c > 0 {
			resStr = "1" + resStr
		}
		return resStr
	}
	var valids []string
	cur := "1"
	for len(cur) <= len(s) {
		valids = append(valids, cur)
		cur = addStr(cur+"00", cur)
	}
	//ans := len(s)
	hist := map[int]int{}
	var dfs func(pos int) int
	dfs = func(pos int) int {
		if pos == len(s) {
			return 0
		}
		if v, ok := hist[pos]; ok {
			return v
		}
		best := -1
		for i := 0; i < len(valids); i++ {
			str := valids[i]
			if pos+len(str) <= len(s) && s[pos:pos+len(str)] == str {
				if nextAns := dfs(pos + len(str)); nextAns != -1 {
					if best == -1 {
						best = nextAns + 1
					} else {
						best = min(nextAns+1, best)
					}
				}
			}
		}
		hist[pos] = best
		return best
	}
	return dfs(0)
}

func main() {
	fmt.Println(minimumBeautifulSubstrings("100111000110111"))
}
