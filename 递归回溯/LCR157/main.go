package main

import "fmt"

func goodsOrder(goods string) []string {
	cnts := make([]int, 26)
	for i := 0; i < len(goods); i++ {
		cnts[goods[i]-'a']++
	}
	var ans []string
	var dfs func(acc []byte, pos int)
	dfs = func(acc []byte, pos int) {
		if pos == len(goods) {
			ans = append(ans, string(acc))
			return
		}
		for i := 0; i < 26; i++ {
			if cnts[i] > 0 {
				cnts[i]--
				acc[pos] = byte('a' + i)
				dfs(acc, pos+1)
				cnts[i]++
			}
		}
	}
	acc := make([]byte, len(goods))
	dfs(acc, 0)
	return ans
}

func main() {
	fmt.Println(goodsOrder("aac"))
}
