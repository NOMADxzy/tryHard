package main

import "fmt"

func solve(weights []int, nexts [][]int, limit int) int {
	ans := 1 << 31
	var dfs func(acc int, pos int, maxVal int)

	dfs = func(acc int, pos int, maxVal int) {
		newAcc := acc + weights[pos]
		if weights[pos] > maxVal {
			maxVal = weights[pos]
		}
		if newAcc >= limit {
			if maxVal < ans {
				ans = maxVal
			}
			return
		}

		for _, nextPos := range nexts[pos] {
			dfs(newAcc, nextPos, maxVal)
		}
	}
	dfs(0, 0, weights[0])
	if ans == 1<<31 {
		return -1
	}
	return ans
}

func main() {
	var n, limit, tmp int
	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&limit)
	arr := make([]int, n)
	nexts := make([][]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&tmp)
		arr[i] = tmp
	}
	for i := 0; i < n; i++ {
		l := 0
		_, _ = fmt.Scan(&l)
		if l == 0 {
			nexts[i] = []int{}
			continue
		}
		childs := make([]int, l)
		for j := 0; j < l; j++ {
			_, _ = fmt.Scan(&tmp)
			childs[j] = tmp - 1
		}
		nexts[i] = childs
	}
	fmt.Println(solve(arr, nexts, limit))
}
