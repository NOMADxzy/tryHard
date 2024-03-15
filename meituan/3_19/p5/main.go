package main

import (
	"fmt"
)

// 逆向并查集
func solve(n int, relations, querys [][]int) []string {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		if parents[i] != i {
			parents[i] = find(parents[i])
		}
		return parents[i]
	}
	union := func(i1, i2 int) {
		parents[find(i1)] = find(i2)
	}
	getHash := func(x, y int) int {
		return x*(n+1) + y
	}
	forgetsMap := map[int]bool{}
	for i := 0; i < len(querys); i++ {
		if querys[i][0] == 1 { // 忘记
			forgetsMap[getHash(querys[i][0], querys[i][1])] = true
		}
	}
	for i := 0; i < len(relations); i++ {
		x, y := relations[i][0], relations[i][1]
		if v := getHash(x, y); forgetsMap[v] {
			continue
		}
		union(x, y)
	}
	var ans []string
	for i := 0; i < len(querys); i++ {
		tp, x, y := querys[i][0], querys[i][1], querys[i][2]
		if tp == 1 {
			union(x, y)
		} else { // 查询
			if find(x) == find(y) {
				ans = append(ans, "Yes")
			} else {
				ans = append(ans, "No")
			}
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}

func main() {
	var n, m, q int
	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&m)
	_, _ = fmt.Scan(&q)
	var relations, querys [][]int
	for i := 0; i < m; i++ {
		a, b := 0, 0
		_, _ = fmt.Scan(&a, &b)
		relations = append(relations, []int{a, b})
	}
	for i := 0; i < q; i++ {
		tp, a, b := 0, 0, 0
		_, _ = fmt.Scan(&tp)
		_, _ = fmt.Scan(&a)
		_, _ = fmt.Scan(&b)
		querys = append(querys, []int{tp, a, b})
	}
	ans := solve(n, relations, querys)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}
