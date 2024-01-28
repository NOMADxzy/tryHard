package main

import "fmt"

func repeatLimitedString(s string, repeatLimit int) string {
	remain := map[byte]int{}
	for i := 0; i < len(s); i++ {
		remain[s[i]]++
	}
	var ans string
	var dfs func(acc []byte, sameCnt int)
	dfs = func(acc []byte, sameCnt int) {
		if len(ans) > 0 {
			return
		}
		noMore := true
		for i := 25; i >= 0; i-- {
			c := byte('a' + i)
			if remain[c] > 0 && (sameCnt < repeatLimit || sameCnt == repeatLimit && c != acc[len(acc)-1]) {
				noMore = false
				remain[c]--
				newSameCnt := sameCnt + 1
				if c != acc[len(acc)-1] {
					newSameCnt = 1
				}
				dfs(append(acc, c), newSameCnt)
			}
		}
		if noMore {
			ans = string(acc[1:])
		}
	}
	dfs([]byte{'_'}, 1)
	return ans
}

func main() {
	fmt.Println(repeatLimitedString("abccccc", 1))
}
