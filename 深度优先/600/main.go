package main

import (
	"fmt"
)

// dfs超时
func findIntegers1(n int) int {
	var dfs func(acc int, preLimit bool)

	ans := 0
	dfs = func(acc int, preLimit bool) {
		if acc > n {
			return
		}
		ans++
		if acc == 0 {
			return
		}
		dfs(acc*2, false)
		if !preLimit {
			dfs(acc*2+1, true)
		}
	}
	dfs(0, false)
	dfs(1, true)
	return ans
}

// 数位dp
func findIntegers(n int) int {
	var num []int
	for i := n; i > 0; i /= 2 {
		num = append(num, i%2)
	}
	for i := 0; i < len(num)/2; i++ {
		num[i], num[len(num)-1-i] = num[len(num)-1-i], num[i]
	}
	hist := map[int]int{}
	var dfs func(pre1 bool, isLimit bool, i int) int
	dfs = func(pre1 bool, isLimit bool, i int) int {
		key := i
		if isLimit {
			key += len(num)
		}
		if pre1 {
			key = -key
		}
		if i == len(num) {
			return 1
		} else if hist[key] > 0 {
			return hist[key]
		}
		res := 0
		res += dfs(false, isLimit && num[i] == 0, i+1)
		if !pre1 && (!isLimit || num[i] == 1) {
			res += dfs(true, isLimit && num[i] == 1, i+1)
		}
		hist[key] = res
		return res
	}
	return dfs(false, true, 0)
}

func main() {
	fmt.Println(findIntegers(5))
}
