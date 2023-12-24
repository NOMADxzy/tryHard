package main

import "fmt"

func flipgame(fronts []int, backs []int) int {
	n := len(fronts)
	mark := map[int]bool{}
	for i := 0; i < n; i++ {
		if fronts[i] == backs[i] {
			mark[fronts[i]] = true
		}
	}
	ans := 1 << 31
	for i := 0; i < n; i++ {
		v := fronts[i]
		if !mark[v] && v < ans {
			ans = v
		}
		v = backs[i]
		if !mark[v] && v < ans {
			ans = v
		}
	}
	if ans == 1<<31 {
		return 0
	} else {
		return ans
	}
}

func main() {
	fmt.Println(flipgame([]int{1, 2, 4, 4, 7}, []int{1, 3, 4, 1, 3}))
}
