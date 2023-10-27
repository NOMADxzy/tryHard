package main

import (
	"fmt"
	"strconv"
)

var check = [10]int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}

// 枚举
func rotatedDigits1(n int) (ans int) {
	for i := 1; i <= n; i++ {
		valid, diff := true, false
		for _, c := range strconv.Itoa(i) {
			if check[c-'0'] == -1 {
				valid = false
			} else if check[c-'0'] == 1 {
				diff = true
			}
		}
		if valid && diff {
			ans++
		}
	}
	return
}

//数位dp
func dfs(preValid bool, preLimit bool, target []int, pos int, ntype []int) int {
	if pos == len(target) {
		if preValid {
			return 1
		} else {
			return 0
		}
	}
	ans := 0
	top := 9
	if preLimit {
		top = target[pos]
	}
	for i := 0; i <= top; i++ {
		nextLimit := false
		if i == top && preLimit {
			nextLimit = true
		}
		if ntype[i] == 0 {
			ans += dfs(preValid, nextLimit, target, pos+1, ntype)
		} else if ntype[i] == 1 {
			ans += dfs(true, nextLimit, target, pos+1, ntype)
		}
	}
	return ans
}

func rotatedDigits(n int) (ans int) {
	var w int
	for i := n; i > 0; i /= 10 {
		w++
	}
	target := make([]int, w)

	for i := 0; i < w; i++ {
		target[w-1-i] = n % 10
		n /= 10
	}

	return dfs(false, true, target, 0, []int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1})
}

func main() {
	fmt.Println(rotatedDigits(857))
}
