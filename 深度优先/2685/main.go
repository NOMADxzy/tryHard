package main

import "fmt"

func countCompleteComponents1(n int, edges [][]int) int {
	nexts := make([][]int, n)
	for _, edge := range edges {
		f, t := edge[0], edge[1]
		nexts[f] = append(nexts[f], t)
		nexts[t] = append(nexts[t], f)
	}

	ans := 0
	var dfs func(prePos, pos int) (int, int)
	visited := map[int]bool{}

	dfs = func(prePos, pos int) (int, int) {
		edgeCnt := 0
		pointCnt := 1
		visited[pos] = true
		for _, np := range nexts[pos] {
			if np != prePos {
				edgeCnt++
			}
			if !visited[np] {
				newEdgeCnt, newPointCnt := dfs(pos, np)
				edgeCnt += newEdgeCnt
				pointCnt += newPointCnt
			}
		}
		return edgeCnt, pointCnt
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			cnt1, cnt2 := dfs(-1, i)
			targetCnt1 := 0
			for j := cnt2 - 1; j >= 0; j-- {
				targetCnt1 += j
			}
			if cnt1 == targetCnt1 {
				ans++
			}
		}
	}
	return ans
}

func countCompleteComponents(n int, edges [][]int) (ans int) {

	g := make([][]int, n)

	for _, e := range edges {

		x, y := e[0], e[1]

		g[x] = append(g[x], y)

		g[y] = append(g[y], x)

	}

	vis := make([]bool, n)

	var v, e int

	var dfs func(int)

	dfs = func(x int) {

		vis[x] = true

		v++

		e += len(g[x])

		for _, y := range g[x] {

			if !vis[y] {

				dfs(y)

			}

		}

	}

	for i, b := range vis {

		if !b {

			v, e = 0, 0

			dfs(i)

			if e == v*(v-1) {

				ans++

			}

		}

	}

	return

}

func main() {
	fmt.Println(countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}))
}
