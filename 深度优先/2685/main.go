package main

import "fmt"

func countCompleteComponents(n int, edges [][]int) int {
	nexts := make([][]int, n)
	for _, edge := range edges {
		f, t := edge[0], edge[1]
		nexts[f] = append(nexts[f], t)
		nexts[t] = append(nexts[t], f)
	}

	ans := 0
	var dfs func(prePos, pos int) (int, int)
	visited := map[int]bool{}
	eVisted := map[int]bool{}

	dfs = func(prePos, pos int) (int, int) {
		edgeCnt := 0
		pointCnt := 1
		visited[pos] = true
		for _, np := range nexts[pos] {
			key := (n+1)*pos + np
			if !eVisted[key] {
				eVisted[key] = true
				eVisted[(n+1)*np+pos] = true
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

func main() {
	fmt.Println(countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}))
}
