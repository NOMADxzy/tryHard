package main

import "fmt"

func takeCharacters(s string, k int) int {
	getIdx := func(b byte) int {
		if b == 'a' {
			return 0
		} else if b == 'b' {
			return 1
		} else {
			return 2
		}
	}
	cnt := [3]int{}
	for i := 0; i < len(s); i++ {
		cnt[getIdx(s[i])]++
	}
	if cnt[0] < k || cnt[1] < k || cnt[2] < k {
		return -1
	}
	cnt[0] -= k
	cnt[1] -= k
	cnt[2] -= k

	l, r := 0, 0
	tmp := [3]int{}
	ans := 0
	for r < len(s) {
		cur := getIdx(s[r])
		tmp[cur]++
		for tmp[cur] > cnt[cur] {
			tmp[getIdx(s[l])]--
			l++
		}
		ans = max(ans, r-l+1)
		r++
	}
	return len(s) - ans
}

func main() {
	fmt.Println(takeCharacters("aabaaaacaabc", 2))
}
