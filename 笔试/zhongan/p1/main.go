package main

import "fmt"

//小强在研究有关字典序的问题。同长度下的字典序比较顺序为从左往右，比如"ac"<"ad","bc">"ad"。
//
//他想知道，如果每种字母组合都能构成一个单词，给定长度为n的两个单词A和B，字典序小于B但大于A且长度等于n的单词有多少个。

func solve(a, b string) int {
	cache := map[int]int{}
	n := len(a)
	var dfs func(pos int, limitMin, limitMax bool) int
	dfs = func(pos int, limitMin, limitMax bool) int {
		limit := 0
		if limitMin && limitMax {
			limit = 3
		} else if limitMin {
			limit = 1
		} else if limitMax {
			limit = 2
		}
		key := limit*n + pos
		if val, ok := cache[key]; ok {
			return val
		}
		if pos == n {
			if limit != 0 {
				return 0
			}
			return 1
		}
		if limit == 0 {
			return 26 * dfs(pos+1, false, false)
		}
		cnt := 0
		for i := uint8('a'); i <= uint8('z'); i++ {
			if i < a[pos] && limitMin || i > b[pos] && limitMax {
				continue
			}
			newLimitMin, newLimitMax := i == a[pos] && limitMin, i == b[pos] && limitMax
			cnt += dfs(pos+1, newLimitMin, newLimitMax)
		}
		cache[key] = cnt
		return cnt
	}
	return dfs(0, true, true)
}

func main() {
	//fmt.Println(solve("aa", "zz"))
	var T, n int
	_, _ = fmt.Scan(&T)

	for i := 0; i < T; i++ {
		_, _ = fmt.Scan(&n)
		var a, b string
		_, _ = fmt.Scan(&a, &b)
		fmt.Println(solve(a, b))
	}
}
