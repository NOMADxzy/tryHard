package main

import "fmt"

func countEval(s string, result int) int {

	hist := [2]map[string]int{{}, {}}
	isSymb := func(b byte) bool {
		if b == '&' || b == '|' || b == '^' {
			return true
		}
		return false
	}
	var dfs func(s string) (int, int)
	dfs = func(s string) (int, int) {
		if v, ok := hist[0][s]; ok {
			return v, hist[1][s]
		}
		if len(s) == 1 {
			if s == "0" {
				return 1, 0
			} else {
				return 0, 1
			}
		}
		cnt0 := 0
		cnt1 := 0
		for i := 0; i < len(s); i++ {
			if isSymb(s[i]) {
				preCnt0, preCnt1 := dfs(s[:i])
				nextCnt0, nextCnt1 := dfs(s[i+1:])
				if s[i] == '&' {
					cnt0 += preCnt1*nextCnt0 + preCnt0*(nextCnt0+nextCnt1)
					cnt1 += preCnt1 * nextCnt1
				} else if s[i] == '|' {
					cnt0 += preCnt0 * nextCnt0
					cnt1 += preCnt0*nextCnt1 + preCnt1*(nextCnt0+nextCnt1)
				} else {
					cnt0 += preCnt0*nextCnt0 + preCnt1*nextCnt1
					cnt1 += preCnt0*nextCnt1 + preCnt1*nextCnt0
				}
			}
		}
		hist[0][s] = cnt0
		hist[1][s] = cnt1
		return cnt0, cnt1
	}
	cnt0, cnt1 := dfs(s)
	if result == 0 {
		return cnt0
	} else {
		return cnt1
	}
}

func main() {
	fmt.Println(countEval("0&0&0&1^1|0", 1))
}
